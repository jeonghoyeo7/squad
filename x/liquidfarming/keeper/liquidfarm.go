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
		ctx.GasMeter().ConsumeGas(sdk.Gas(numQueuedFarmings)*k.GetParams(ctx).DelayedFarmGasFee, "DelayedFarmGasFee")
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
		),
	})

	return nil
}

// UnfarmInfo holds information about unfarm.
type UnfarmInfo struct {
	Farmer       sdk.AccAddress
	UnfarmedCoin sdk.Coin
}

// Unfarm handles types.MsgUnfarm to unfarm LFCoin.
// The logic doesn't check whether or not liquid farm exists because it can be removed for some reason and
// farmers still need to be able to unfarm their pool coin.
func (k Keeper) Unfarm(ctx sdk.Context, poolId uint64, farmer sdk.AccAddress, unfarmingCoin sdk.Coin) (UnfarmInfo, error) {
	reserveAddr := types.LiquidFarmReserveAddress(poolId)
	lfCoinDenom := types.LiquidFarmCoinDenom(poolId)

	lfCoinBalance := k.bankKeeper.SpendableCoins(ctx, farmer).AmountOf(lfCoinDenom)
	if lfCoinBalance.LT(unfarmingCoin.Amount) {
		return UnfarmInfo{},
			sdkerrors.Wrapf(types.ErrInsufficientUnfarmingAmount, "%s is smaller than %s", lfCoinBalance, unfarmingCoin.Amount)
	}

	pool, found := k.liquidityKeeper.GetPool(ctx, poolId)
	if !found {
		return UnfarmInfo{},
			sdkerrors.Wrapf(sdkerrors.ErrNotFound, "pool %d not found", poolId)
	}

	lfCoinTotalSupply := k.bankKeeper.GetSupply(ctx, lfCoinDenom).Amount
	lpCoinTotalStaked := k.farmingKeeper.GetAllStakedCoinsByFarmer(ctx, reserveAddr).AmountOf(pool.PoolCoinDenom)
	unfarmFee := sdk.ZeroInt() // TODO: TBD

	// UnfarmedAmount = TotalStakedLPAmount / TotalSupplyLFAmount * UnfarmingLFAmount * (1 - UnfarmFee)
	unfarmedAmt := lpCoinTotalStaked.Quo(lfCoinTotalSupply).Mul(unfarmingCoin.Amount).Mul(sdk.OneInt().Sub(unfarmFee))
	unfarmedCoin := sdk.NewCoin(pool.PoolCoinDenom, unfarmedAmt)

	// Send the unfarming LFCoin to module account and burn them
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, farmer, types.ModuleName, sdk.NewCoins(unfarmingCoin)); err != nil {
		return UnfarmInfo{}, err
	}
	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(unfarmingCoin)); err != nil {
		return UnfarmInfo{}, err
	}

	// Unstake with the reserve account and release corresponding pool coin amount back to the farmer
	if err := k.farmingKeeper.Unstake(ctx, reserveAddr, sdk.NewCoins(unfarmedCoin)); err != nil {
		return UnfarmInfo{}, err
	}
	if err := k.bankKeeper.SendCoins(ctx, reserveAddr, farmer, sdk.NewCoins(unfarmedCoin)); err != nil {
		return UnfarmInfo{}, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUnfarm,
			sdk.NewAttribute(types.AttributeKeyPoolId, strconv.FormatUint(poolId, 10)),
			sdk.NewAttribute(types.AttributeKeyFarmer, farmer.String()),
			sdk.NewAttribute(types.AttributeKeyUnfarmingCoin, unfarmingCoin.String()),
			sdk.NewAttribute(types.AttributeKeyUnfarmedCoin, unfarmedCoin.String()),
		),
	})

	info := UnfarmInfo{
		Farmer:       farmer,
		UnfarmedCoin: unfarmedCoin,
	}

	return info, nil
}

// UnfarmAndWithdraw handles types.MsgUnfarmAndWithdraw to unfarm LFCoin and withdraw pool coin from the pool.
func (k Keeper) UnfarmAndWithdraw(ctx sdk.Context, msg *types.MsgUnfarmAndWithdraw) error {
	unfarmInfo, err := k.Unfarm(ctx, msg.PoolId, msg.GetFarmer(), msg.UnfarmingCoin)
	if err != nil {
		return sdkerrors.Wrapf(err, "unable to unfarm")
	}

	balance := k.bankKeeper.SpendableCoins(ctx, msg.GetFarmer()).AmountOf(unfarmInfo.UnfarmedCoin.Denom)
	if balance.LT(unfarmInfo.UnfarmedCoin.Amount) {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "%s is smaller than %s", balance, unfarmInfo.UnfarmedCoin.Amount)
	}

	_, err = k.liquidityKeeper.Withdraw(ctx, &liquiditytypes.MsgWithdraw{
		PoolId:     msg.PoolId,
		Withdrawer: msg.Farmer,
		PoolCoin:   unfarmInfo.UnfarmedCoin,
	})
	if err != nil {
		return sdkerrors.Wrapf(err, "unable to withdraw unfarmed coin from the pool")
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUnfarmAndWithdraw,
			sdk.NewAttribute(types.AttributeKeyPoolId, strconv.FormatUint(msg.PoolId, 10)),
			sdk.NewAttribute(types.AttributeKeyFarmer, msg.Farmer),
			sdk.NewAttribute(types.AttributeKeyUnfarmedCoin, unfarmInfo.UnfarmedCoin.String()),
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
