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
func (k Keeper) ValidateMsgPlaceBid(ctx sdk.Context, msg *types.MsgPlaceBid) error {
	liquidFarm, found := k.GetLiquidFarm(ctx, msg.PoolId)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "liquid farm by pool %d not found", msg.PoolId)
	}

	auctionId := k.GetLastRewardsAuctionId(ctx, msg.PoolId)
	_, found = k.GetRewardsAuction(ctx, msg.PoolId, auctionId)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "auction by pool %d not found", msg.PoolId)
	}

	balance := k.bankKeeper.SpendableCoins(ctx, msg.GetBidder()).AmountOf(msg.BiddingCoin.Denom)
	if balance.LT(msg.BiddingCoin.Amount) {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "%s is smaller than %s", balance, msg.BiddingCoin.Amount)
	}

	if msg.BiddingCoin.Amount.LT(liquidFarm.MinimumBidAmount) {
		return sdkerrors.Wrapf(types.ErrSmallerThanMinimumAmount, "%s is smaller than %s", msg.BiddingCoin.Amount, liquidFarm.MinimumBidAmount)
	}

	_, found = k.GetBid(ctx, auctionId, msg.GetBidder())
	if found {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "refund the previous bid to place new bid")
	}

	return nil
}

// PlaceBid handles types.MsgPlaceBid and stores bid object.
func (k Keeper) PlaceBid(ctx sdk.Context, msg *types.MsgPlaceBid) (types.Bid, error) {
	if err := k.ValidateMsgPlaceBid(ctx, msg); err != nil {
		return types.Bid{}, err
	}

	auctionId := k.GetLastRewardsAuctionId(ctx, msg.PoolId)
	auction, _ := k.GetRewardsAuction(ctx, msg.PoolId, auctionId)

	if err := k.bankKeeper.SendCoins(ctx, msg.GetBidder(), auction.GetPayingReserveAddress(), sdk.NewCoins(msg.BiddingCoin)); err != nil {
		return types.Bid{}, err
	}

	bid := types.NewBid(
		msg.PoolId,
		msg.Bidder,
		msg.BiddingCoin,
	)
	k.SetBid(ctx, bid)
	k.SetWinningBid(ctx, bid, auction.Id)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypePlaceBid,
			sdk.NewAttribute(types.AttributeKeyPoolId, strconv.FormatUint(msg.PoolId, 10)),
			sdk.NewAttribute(types.AttributeKeyAuctionId, strconv.FormatUint(auction.Id, 10)),
			sdk.NewAttribute(types.AttributeKeyBidder, msg.Bidder),
			sdk.NewAttribute(types.AttributeKeyBiddingCoin, msg.BiddingCoin.String()),
		),
	})

	return bid, nil
}

// RefundBid handles types.MsgRefundBid and refunds bid amount to the bidder and
// delete the bid object.
func (k Keeper) RefundBid(ctx sdk.Context, msg *types.MsgRefundBid) error {
	auctionId := k.GetLastRewardsAuctionId(ctx, msg.PoolId)
	auction, found := k.GetRewardsAuction(ctx, msg.PoolId, auctionId)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "auction by pool %d not found", msg.PoolId)
	}

	winningBid, found := k.GetWinningBid(ctx, msg.PoolId, auctionId)
	if found && winningBid.Bidder == msg.Bidder {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "winning bid can't be refunded")
	}

	bid, found := k.GetBid(ctx, msg.PoolId, msg.GetBidder())
	if !found {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "bid not found")
	}

	if err := k.bankKeeper.SendCoins(ctx, auction.GetPayingReserveAddress(), msg.GetBidder(), sdk.NewCoins(bid.Amount)); err != nil {
		return err
	}

	k.DeleteBid(ctx, bid)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeRefundBid,
			sdk.NewAttribute(types.AttributeKeyPoolId, strconv.FormatUint(msg.PoolId, 10)),
			sdk.NewAttribute(types.AttributeKeyBidder, msg.Bidder),
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
func (k Keeper) FinishRewardsAuction(ctx sdk.Context, auction types.RewardsAuction) {
	liquidFarmReserveAddr := types.LiquidFarmReserveAddress(auction.PoolId)
	payingReserveAddr := auction.GetPayingReserveAddress()
	poolCoinDenom := liquiditytypes.PoolCoinDenom(auction.PoolId)
	rewards := k.farmingKeeper.Rewards(ctx, liquidFarmReserveAddr, poolCoinDenom)
	compoundingRewards := types.CompoundingRewards{Amount: sdk.ZeroInt()}

	// Finishing a rewards auction can have two different scenarios depending on winning bid existence
	// When there is winning bid, harvest farming rewards first and send them to the winner and
	// stake the winning bid amount in the farming module for farmers so that it acts as auto compounding functionality.
	winningBid, found := k.GetWinningBid(ctx, auction.PoolId, auction.Id)
	if found {
		// TODO: do we need to all these panics?
		if err := k.farmingKeeper.Harvest(ctx, liquidFarmReserveAddr, []string{poolCoinDenom}); err != nil {
			panic(err)
		}

		if err := k.bankKeeper.SendCoins(ctx, liquidFarmReserveAddr, winningBid.GetBidder(), rewards); err != nil {
			panic(err)
		}

		if err := k.RefundAllBids(ctx, auction, winningBid); err != nil {
			panic(err)
		}

		if err := k.bankKeeper.SendCoins(ctx, payingReserveAddr, liquidFarmReserveAddr, sdk.NewCoins(winningBid.Amount)); err != nil {
			panic(err)
		}

		if err := k.farmingKeeper.Stake(ctx, liquidFarmReserveAddr, sdk.NewCoins(winningBid.Amount)); err != nil {
			panic(err)
		}

		compoundingRewards.Amount = winningBid.Amount.Amount
	}

	auction.SetWinner(winningBid.Bidder)
	auction.SetRewards(rewards)
	auction.SetStatus(types.AuctionStatusFinished)
	k.SetRewardsAuction(ctx, auction)
	k.SetCompoundingRewards(ctx, auction.PoolId, compoundingRewards)
}
