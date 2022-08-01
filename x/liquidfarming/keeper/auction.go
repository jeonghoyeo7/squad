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
	auctionId := k.GetLastRewardsAuctionId(ctx, msg.PoolId)
	_, found := k.GetRewardsAuction(ctx, msg.PoolId, auctionId)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "auction %d not found", auctionId)
	}

	liquidFarm, found := k.GetLiquidFarm(ctx, msg.PoolId)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "liquid farm by pool %d not found", msg.PoolId)
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
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "refund bid to place new bid")
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

// ClaimFarmingRewards harvests farming rewards and send them to the winner.
func (k Keeper) ClaimFarmingRewards(ctx sdk.Context, poolId uint64, winner sdk.AccAddress, rewards sdk.Coins) error {
	liquidFarmReserveAddr := types.LiquidFarmReserveAddress(poolId)
	poolCoinDenom := liquiditytypes.PoolCoinDenom(poolId)

	if err := k.farmingKeeper.Harvest(ctx, liquidFarmReserveAddr, []string{poolCoinDenom}); err != nil {
		return err
	}

	if err := k.bankKeeper.SendCoins(ctx, liquidFarmReserveAddr, winner, rewards); err != nil {
		return err
	}

	return nil
}

// RefundAllBids refunds all bids at once as the rewards auction is finished and delete all bids.
func (k Keeper) RefundAllBids(ctx sdk.Context, poolId uint64, payingReserveAddr, winner sdk.AccAddress) error {
	inputs := []banktypes.Input{}
	outputs := []banktypes.Output{}
	for _, bid := range k.GetBidsByPoolId(ctx, poolId) {
		if bid.Bidder != winner.String() { // skip the winning bid
			inputs = append(inputs, banktypes.NewInput(payingReserveAddr, sdk.NewCoins(bid.Amount)))
			outputs = append(outputs, banktypes.NewOutput(bid.GetBidder(), sdk.NewCoins(bid.Amount)))

			k.DeleteBid(ctx, bid)
		}
	}
	if err := k.bankKeeper.InputOutputCoins(ctx, inputs, outputs); err != nil {
		panic(err)
	}
	return nil
}

// FinishAndCreateRewardsAuction finishes the ongoing rewards auction and creates a new one.
func (k Keeper) FinishAndCreateRewardsAuction(ctx sdk.Context, auction types.RewardsAuction) {
	liquidFarmReserveAddr := types.LiquidFarmReserveAddress(auction.PoolId)
	poolCoinDenom := liquiditytypes.PoolCoinDenom(auction.PoolId)
	rewards := k.farmingKeeper.Rewards(ctx, liquidFarmReserveAddr, poolCoinDenom)

	// Finish auction states
	winningBid, found := k.GetWinningBid(ctx, auction.PoolId, auction.Id)
	auction.SetWinner(winningBid.Bidder)
	auction.SetRewards(rewards)
	auction.SetStatus(types.AuctionStatusFinished)
	k.SetRewardsAuction(ctx, auction)

	if found {
		payingReserveAddr := auction.GetPayingReserveAddress()

		if err := k.ClaimFarmingRewards(ctx, auction.PoolId, winningBid.GetBidder(), rewards); err != nil {
			panic(err)
		}

		if err := k.RefundAllBids(ctx, auction.PoolId, payingReserveAddr, winningBid.GetBidder()); err != nil {
			panic(err)
		}

		if err := k.bankKeeper.SendCoins(ctx, payingReserveAddr, liquidFarmReserveAddr, sdk.NewCoins(winningBid.Amount)); err != nil {
			panic(err)
		}

		// Stake winning bid amount for auto compounding
		if err := k.farmingKeeper.Stake(ctx, liquidFarmReserveAddr, sdk.NewCoins(winningBid.Amount)); err != nil {
			panic(err)
		}
	}

	// Covers a case when there isn't a single bid placed for the auction
	k.CreateRewardsAuction(ctx, auction.PoolId)
}
