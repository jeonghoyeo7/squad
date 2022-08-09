package keeper

import (
	"fmt"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	farmingtypes "github.com/cosmosquad-labs/squad/v2/x/farming/types"
	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"
	liquiditytypes "github.com/cosmosquad-labs/squad/v2/x/liquidity/types"
)

// ValidateMsgFarm validates types.MsgFarm.
func (k Keeper) ValidateMsgFarm(ctx sdk.Context, msg *types.MsgFarm) error {
	liquidFarm, found := k.GetLiquidFarm(ctx, msg.PoolId)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "liquid farm by pool %d not found", msg.PoolId)
	}

	if msg.FarmingCoin.Amount.LT(liquidFarm.MinimumFarmAmount) {
		return sdkerrors.Wrapf(types.ErrSmallerThanMinimumAmount, "%s is smaller than %s", msg.FarmingCoin.Amount, liquidFarm.MinimumFarmAmount)
	}

	pool, found := k.liquidityKeeper.GetPool(ctx, liquidFarm.PoolId)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "pool %d not found", liquidFarm.PoolId)
	}

	if pool.PoolCoinDenom != msg.FarmingCoin.Denom {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "expected denom %s, but got %s", pool.PoolCoinDenom, msg.FarmingCoin.Denom)
	}

	poolCoinBalance := k.bankKeeper.SpendableCoins(ctx, msg.GetFarmer()).AmountOf(pool.PoolCoinDenom)
	if poolCoinBalance.LT(msg.FarmingCoin.Amount) {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "%s is smaller than %s", poolCoinBalance, msg.FarmingCoin.Amount)
	}

	return nil
}

// Farm handles types.MsgFarm to liquid farm.
func (k Keeper) Farm(ctx sdk.Context, msg *types.MsgFarm) error {
	if err := k.ValidateMsgFarm(ctx, msg); err != nil {
		return err
	}

	reserveAddr := types.LiquidFarmReserveAddress(msg.PoolId)
	if err := k.bankKeeper.SendCoins(ctx, msg.GetFarmer(), reserveAddr, sdk.NewCoins(msg.FarmingCoin)); err != nil {
		return err
	}

	// Impose more gas in relative to a number of queued farmings farmed by the farmer
	// This prevents from potential spamming attack
	numQueuedFarmings := 0
	for range k.GetQueuedFarmingsByFarmer(ctx, msg.GetFarmer()) {
		numQueuedFarmings++
	}
	if numQueuedFarmings > 0 {
		params := k.GetParams(ctx)
		ctx.GasMeter().ConsumeGas(sdk.Gas(numQueuedFarmings)*params.DelayedFarmGasFee, "DelayedFarmGasFee")
	}

	// Stake in the farming module with the reserve account
	if err := k.farmingKeeper.Stake(ctx, reserveAddr, sdk.NewCoins(msg.FarmingCoin)); err != nil {
		return err
	}

	currentEpochDays := k.farmingKeeper.GetCurrentEpochDays(ctx)
	endTime := ctx.BlockTime().Add(time.Duration(currentEpochDays) * farmingtypes.Day) // current time + epoch days

	k.SetQueuedFarming(ctx, endTime, liquiditytypes.PoolCoinDenom(msg.PoolId), msg.GetFarmer(), types.QueuedFarming{
		PoolId: msg.PoolId,
		Amount: msg.FarmingCoin.Amount,
	})

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeFarm,
			sdk.NewAttribute(types.AttributeKeyPoolId, strconv.FormatUint(msg.PoolId, 10)),
			sdk.NewAttribute(types.AttributeKeyFarmer, msg.Farmer),
			sdk.NewAttribute(types.AttributeKeyFarmingCoin, msg.FarmingCoin.String()),
			sdk.NewAttribute(types.AttributeKeyLiquidFarmReserveAddress, reserveAddr.String()),
		),
	})

	return nil
}

// UnfarmInfo holds information about unfarm.
type UnfarmInfo struct {
	Farmer     sdk.AccAddress
	UnfarmCoin sdk.Coin
}

// ValidateMsgUnfarm validates MsgUnfarm.
// It doesn't check if the liquid farm exists because farmers need to be able to unfarm their LFCoin
// even if the liquid farm is removed in params.
func (k Keeper) ValidateMsgUnfarm(ctx sdk.Context, poolId uint64, farmer sdk.AccAddress, unfarmingCoin sdk.Coin) error {
	lfCoinBalance := k.bankKeeper.SpendableCoins(ctx, farmer).AmountOf(types.LiquidFarmCoinDenom(poolId))
	if lfCoinBalance.LT(unfarmingCoin.Amount) {
		return sdkerrors.Wrapf(types.ErrInsufficientUnfarmingAmount, "%s is smaller than %s", lfCoinBalance, unfarmingCoin.Amount)
	}

	_, found := k.liquidityKeeper.GetPool(ctx, poolId)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "pool %d not found", poolId)
	}

	return nil
}

// Unfarm handles types.MsgUnfarm to unfarm LFCoin.
func (k Keeper) Unfarm(ctx sdk.Context, poolId uint64, farmer sdk.AccAddress, unfarmingCoin sdk.Coin) (UnfarmInfo, error) {
	if err := k.ValidateMsgUnfarm(ctx, poolId, farmer, unfarmingCoin); err != nil {
		return UnfarmInfo{}, err
	}

	reserveAddr := types.LiquidFarmReserveAddress(poolId)
	lfCoinDenom := types.LiquidFarmCoinDenom(poolId)
	poolCoinDenom := liquiditytypes.PoolCoinDenom(poolId)

	totalSupplyLFAmt := k.bankKeeper.GetSupply(ctx, lfCoinDenom).Amount
	totalStakedLPAmt := k.farmingKeeper.GetAllStakedCoinsByFarmer(ctx, reserveAddr).AmountOf(poolCoinDenom)
	_, found := k.GetLiquidFarm(ctx, poolId)
	if !found {
		// Handle a case when liquid farm is removed in params
		// The reserve account must have unstaked all coins from the farming module so
		// Add all pool coin balances and deduct all queued coins from queued farming objects
		spendable := k.bankKeeper.SpendableCoins(ctx, reserveAddr).AmountOf(poolCoinDenom)
		allQueuedAmt := k.GetQueuedFarmingCoinByPoolId(ctx, poolId).Amount
		totalStakedLPAmt = spendable.Sub(allQueuedAmt)
	}

	unfarmAmt := types.CalculateUnfarmAmount(totalStakedLPAmt, totalSupplyLFAmt, unfarmingCoin.Amount, k.GetUnfarmFeeRate(ctx))
	unfarmCoin := sdk.NewCoin(poolCoinDenom, unfarmAmt)

	// Burn the unfarming LFCoin by sending it to module account
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, farmer, types.ModuleName, sdk.NewCoins(unfarmingCoin)); err != nil {
		return UnfarmInfo{}, err
	}
	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(unfarmingCoin)); err != nil {
		return UnfarmInfo{}, err
	}

	// Unstake unfarm coin in the farming module and release it to the farmer
	if err := k.farmingKeeper.Unstake(ctx, reserveAddr, sdk.NewCoins(unfarmCoin)); err != nil {
		return UnfarmInfo{}, err
	}
	if err := k.bankKeeper.SendCoins(ctx, reserveAddr, farmer, sdk.NewCoins(unfarmCoin)); err != nil {
		return UnfarmInfo{}, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUnfarm,
			sdk.NewAttribute(types.AttributeKeyPoolId, strconv.FormatUint(poolId, 10)),
			sdk.NewAttribute(types.AttributeKeyFarmer, farmer.String()),
			sdk.NewAttribute(types.AttributeKeyUnfarmingCoin, unfarmingCoin.String()),
			sdk.NewAttribute(types.AttributeKeyUnfarmCoin, unfarmCoin.String()),
		),
	})

	return UnfarmInfo{Farmer: farmer, UnfarmCoin: unfarmCoin}, nil
}

// UnfarmAndWithdraw handles types.MsgUnfarmAndWithdraw to unfarm LFCoin and withdraw pool coin from the pool.
func (k Keeper) UnfarmAndWithdraw(ctx sdk.Context, msg *types.MsgUnfarmAndWithdraw) error {
	unfarmInfo, err := k.Unfarm(ctx, msg.PoolId, msg.GetFarmer(), msg.UnfarmingCoin)
	if err != nil {
		return sdkerrors.Wrapf(err, "unable to unfarm")
	}

	spendable := k.bankKeeper.SpendableCoins(ctx, msg.GetFarmer()).AmountOf(unfarmInfo.UnfarmCoin.Denom)
	if spendable.LT(unfarmInfo.UnfarmCoin.Amount) {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "%s is smaller than %s", spendable, unfarmInfo.UnfarmCoin.Amount)
	}

	if _, err = k.liquidityKeeper.Withdraw(ctx, &liquiditytypes.MsgWithdraw{
		PoolId:     msg.PoolId,
		Withdrawer: msg.Farmer,
		PoolCoin:   unfarmInfo.UnfarmCoin,
	}); err != nil {
		return sdkerrors.Wrapf(err, "unable to withdraw coin from the pool")
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUnfarmAndWithdraw,
			sdk.NewAttribute(types.AttributeKeyPoolId, strconv.FormatUint(msg.PoolId, 10)),
			sdk.NewAttribute(types.AttributeKeyFarmer, msg.Farmer),
			sdk.NewAttribute(types.AttributeKeyUnfarmCoin, unfarmInfo.UnfarmCoin.String()),
		),
	})

	return nil
}

// CancelQueuedFarming handles types.MsgCancelQueuedFarming to cancel queued farming.
func (k Keeper) CancelQueuedFarming(ctx sdk.Context, msg *types.MsgCancelQueuedFarming) error {
	queuedFarmings := k.GetQueuedFarmingsByFarmer(ctx, msg.GetFarmer())
	if len(queuedFarmings) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "queued farming by %s not found", msg.Farmer)
	}

	farmerAddr := msg.GetFarmer()
	farmingCoin := msg.UnfarmingCoin

	canceled := sdk.ZeroInt()
	k.IterateQueuedFarmingsByFarmerAndDenomReverse(ctx, farmerAddr, farmingCoin.Denom, func(endTime time.Time, queuedFarming types.QueuedFarming) (stop bool) {
		if endTime.After(ctx.BlockTime()) { // sanity check
			amtToCancel := sdk.MinInt(farmingCoin.Amount.Sub(canceled), queuedFarming.Amount)
			queuedFarming.Amount = queuedFarming.Amount.Sub(amtToCancel)
			if queuedFarming.Amount.IsZero() {
				k.DeleteQueuedFarming(ctx, endTime, farmingCoin.Denom, farmerAddr)
			} else {
				k.SetQueuedFarming(ctx, endTime, farmingCoin.Denom, farmerAddr, queuedFarming)
			}

			canceled = canceled.Add(amtToCancel)
			if canceled.Equal(farmingCoin.Amount) { // fully canceled from queued farmings, so stop
				return true
			}
		}
		return false
	})

	if farmingCoin.Amount.GT(canceled) {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "%s is smaller than %s", farmingCoin.Amount, canceled)
	}

	reserveAddr := types.LiquidFarmReserveAddress(msg.PoolId)
	canceledCoin := sdk.NewCoin(farmingCoin.Denom, canceled)

	// Unstake the canceled amount with the reserve account in the farming module
	if err := k.farmingKeeper.Unstake(ctx, reserveAddr, sdk.NewCoins(canceledCoin)); err != nil {
		return err
	}

	// Release the corresponding pool coin amount back to the farmer
	if err := k.bankKeeper.SendCoins(ctx, reserveAddr, msg.GetFarmer(), sdk.NewCoins(canceledCoin)); err != nil {
		return err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCancelQueuedFarming,
			sdk.NewAttribute(types.AttributeKeyPoolId, strconv.FormatUint(msg.PoolId, 10)),
			sdk.NewAttribute(types.AttributeKeyFarmer, msg.Farmer),
			sdk.NewAttribute(types.AttributeKeyCanceledCoin, canceledCoin.String()),
		),
	})

	return nil
}

// RemoveLiquidFarm unstakes all staked pool coins from the farming module and
// remove the liquid farm object in the store
func (k Keeper) RemoveLiquidFarm(ctx sdk.Context, liquidFarm types.LiquidFarm) error {
	reserveAddr := types.LiquidFarmReserveAddress(liquidFarm.PoolId)
	stakedCoins := k.farmingKeeper.GetAllStakedCoinsByFarmer(ctx, reserveAddr)

	// TODO: handle when stakedCoins is zero?

	if err := k.farmingKeeper.Unstake(ctx, reserveAddr, stakedCoins); err != nil {
		return err
	}

	auctionId := k.GetLastRewardsAuctionId(ctx, liquidFarm.PoolId)
	auction, found := k.GetRewardsAuction(ctx, liquidFarm.PoolId, auctionId)
	if !found {
		panic(fmt.Errorf("rewards auction %d must exist, but somehow it is not found", auctionId))
	}

	if err := k.RefundAllBids(ctx, auction, types.Bid{}); err != nil {
		return err
	}

	auction.SetStatus(types.AuctionStatusFinished)
	k.SetRewardsAuction(ctx, auction)
	k.DeleteWinningBid(ctx, liquidFarm.PoolId, auctionId)
	k.DeleteLiquidFarm(ctx, liquidFarm)

	return nil
}
