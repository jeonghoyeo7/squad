package keeper

import (
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	farmingtypes "github.com/cosmosquad-labs/squad/v2/x/farming/types"
	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"
	liquiditytypes "github.com/cosmosquad-labs/squad/v2/x/liquidity/types"
)

// ValidateMsgPlaceBid validates types.MsgPlaceBid.
func (k Keeper) ValidateMsgPlaceBid(ctx sdk.Context, poolId uint64, bidder sdk.AccAddress, biddingCoin sdk.Coin) error {
	liquidFarm, found := k.GetLiquidFarm(ctx, poolId)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "liquid farm by pool %d not found", poolId)
	}

	auctionId := k.GetLastRewardsAuctionId(ctx, poolId)
	_, found = k.GetRewardsAuction(ctx, poolId, auctionId)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "auction by pool %d not found", poolId)
	}

	if biddingCoin.Amount.LT(liquidFarm.MinimumBidAmount) {
		return sdkerrors.Wrapf(types.ErrSmallerThanMinimumAmount, "%s is smaller than %s", biddingCoin.Amount, liquidFarm.MinimumBidAmount)
	}

	// REVIEW: how about making auto refund bid so that bidders cae place their bid?
	_, found = k.GetBid(ctx, auctionId, bidder)
	if found {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "refund the previous bid to place new bid")
	}

	return nil
}

// PlaceBid handles types.MsgPlaceBid and stores bid object.
func (k Keeper) PlaceBid(ctx sdk.Context, poolId uint64, bidder sdk.AccAddress, biddingCoin sdk.Coin) (types.Bid, error) {
	if err := k.ValidateMsgPlaceBid(ctx, poolId, bidder, biddingCoin); err != nil {
		return types.Bid{}, err
	}

	// REVIEW: validation check for the amount that exceeds winning bid amount? (verify with test case)

	auctionId := k.GetLastRewardsAuctionId(ctx, poolId)
	auction, _ := k.GetRewardsAuction(ctx, poolId, auctionId)

	if err := k.bankKeeper.SendCoins(ctx, bidder, auction.GetPayingReserveAddress(), sdk.NewCoins(biddingCoin)); err != nil {
		return types.Bid{}, err
	}

	bid := types.NewBid(
		poolId,
		bidder.String(),
		biddingCoin,
	)
	k.SetBid(ctx, bid)
	k.SetWinningBid(ctx, bid, auction.Id)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypePlaceBid,
			sdk.NewAttribute(types.AttributeKeyPoolId, strconv.FormatUint(poolId, 10)),
			sdk.NewAttribute(types.AttributeKeyAuctionId, strconv.FormatUint(auction.Id, 10)),
			sdk.NewAttribute(types.AttributeKeyBidder, bidder.String()),
			sdk.NewAttribute(types.AttributeKeyBiddingCoin, biddingCoin.String()),
		),
	})

	return bid, nil
}

// RefundBid handles types.MsgRefundBid and refunds bid amount to the bidder and
// delete the bid object.
func (k Keeper) RefundBid(ctx sdk.Context, poolId uint64, bidder sdk.AccAddress) error {
	auctionId := k.GetLastRewardsAuctionId(ctx, poolId)
	auction, found := k.GetRewardsAuction(ctx, poolId, auctionId)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "auction by pool %d not found", poolId)
	}

	winningBid, found := k.GetWinningBid(ctx, poolId, auctionId)
	if found && winningBid.Bidder == bidder.String() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "winning bid can't be refunded")
	}

	bid, found := k.GetBid(ctx, poolId, bidder)
	if !found {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "bid not found")
	}

	if err := k.bankKeeper.SendCoins(ctx, auction.GetPayingReserveAddress(), bidder, sdk.NewCoins(bid.Amount)); err != nil {
		return err
	}

	k.DeleteBid(ctx, bid)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeRefundBid,
			sdk.NewAttribute(types.AttributeKeyPoolId, strconv.FormatUint(poolId, 10)),
			sdk.NewAttribute(types.AttributeKeyBidder, bidder.String()),
			sdk.NewAttribute(types.AttributeKeyRefundCoin, bid.Amount.String()),
		),
	})

	return nil
}

// getNextAuctionIdWithUpdate increments rewards auction id by one and store it.
func (k Keeper) getNextAuctionIdWithUpdate(ctx sdk.Context, poolId uint64) uint64 {
	id := k.GetLastRewardsAuctionId(ctx, poolId) + 1
	k.SetRewardsAuctionId(ctx, poolId, id)
	return id
}

// CreateRewardsAuction creates new rewards auction and store it.
func (k Keeper) CreateRewardsAuction(ctx sdk.Context, poolId uint64) {
	nextAuctionId := k.getNextAuctionIdWithUpdate(ctx, poolId)
	poolCoinDenom := liquiditytypes.PoolCoinDenom(poolId)
	currentEpochDays := k.farmingKeeper.GetCurrentEpochDays(ctx)
	startTime := ctx.BlockTime()
	endTime := startTime.Add(time.Duration(currentEpochDays) * farmingtypes.Day)

	k.SetRewardsAuction(ctx, types.NewRewardsAuction(
		nextAuctionId,
		poolId,
		poolCoinDenom,
		startTime,
		endTime,
	))
}

// RefundAllBids refunds all bids at once as the rewards auction is finished and delete all bids.
func (k Keeper) RefundAllBids(ctx sdk.Context, auction types.RewardsAuction, winningBid types.Bid) error {
	inputs := []banktypes.Input{}
	outputs := []banktypes.Output{}
	for _, bid := range k.GetBidsByPoolId(ctx, auction.PoolId) {
		if bid.Bidder != winningBid.Bidder {
			inputs = append(inputs, banktypes.NewInput(auction.GetPayingReserveAddress(), sdk.NewCoins(bid.Amount)))
			outputs = append(outputs, banktypes.NewOutput(bid.GetBidder(), sdk.NewCoins(bid.Amount)))
		}
		k.DeleteBid(ctx, bid) // delete all
	}
	if err := k.bankKeeper.InputOutputCoins(ctx, inputs, outputs); err != nil {
		panic(err)
	}
	return nil
}

// FinishRewardsAuction finishes the ongoing rewards auction.
func (k Keeper) FinishRewardsAuction(ctx sdk.Context, auction types.RewardsAuction) error {
	liquidFarmReserveAddr := types.LiquidFarmReserveAddress(auction.PoolId)
	payingReserveAddr := auction.GetPayingReserveAddress()
	poolCoinDenom := liquiditytypes.PoolCoinDenom(auction.PoolId)
	rewards := k.farmingKeeper.Rewards(ctx, liquidFarmReserveAddr, poolCoinDenom)
	compoundingRewards := types.CompoundingRewards{Amount: sdk.ZeroInt()}
	status := types.AuctionStatusFinished

	// Finishing a rewards auction can have two different scenarios depending on winning bid existence
	// When there is winning bid, harvest farming rewards first and send them to the winner and
	// stake the winning bid amount in the farming module for farmers so that it acts as auto compounding functionality.
	winningBid, found := k.GetWinningBid(ctx, auction.PoolId, auction.Id)
	if found {
		if err := k.farmingKeeper.Harvest(ctx, liquidFarmReserveAddr, []string{poolCoinDenom}); err != nil {
			return err
		}

		if err := k.bankKeeper.SendCoins(ctx, liquidFarmReserveAddr, winningBid.GetBidder(), rewards); err != nil {
			return err
		}

		if err := k.RefundAllBids(ctx, auction, winningBid); err != nil {
			return err
		}

		if err := k.bankKeeper.SendCoins(ctx, payingReserveAddr, liquidFarmReserveAddr, sdk.NewCoins(winningBid.Amount)); err != nil {
			return err
		}

		if err := k.farmingKeeper.Stake(ctx, liquidFarmReserveAddr, sdk.NewCoins(winningBid.Amount)); err != nil {
			return err
		}

		compoundingRewards.Amount = winningBid.Amount.Amount
	} else {
		status = types.AuctionStatusSkipped
	}

	auction.SetWinner(winningBid.Bidder)
	auction.SetRewards(rewards)
	auction.SetStatus(status)
	k.SetRewardsAuction(ctx, auction)
	k.SetCompoundingRewards(ctx, auction.PoolId, compoundingRewards)

	return nil
}
