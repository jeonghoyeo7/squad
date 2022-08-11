package keeper

import (
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"
	liquiditytypes "github.com/cosmosquad-labs/squad/v2/x/liquidity/types"
)

// ValidateMsgFarm validates types.MsgFarm.
func (k Keeper) ValidateMsgFarm(ctx sdk.Context, msg *types.MsgFarm) error {
	pool, found := k.liquidityKeeper.GetPool(ctx, msg.PoolId)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "pool %d not found", msg.PoolId)
	}

	liquidFarm, found := k.GetLiquidFarm(ctx, msg.PoolId)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "liquid farm by pool %d not found", msg.PoolId)
	}

	if msg.FarmingCoin.Amount.LT(liquidFarm.MinimumFarmAmount) {
		return sdkerrors.Wrapf(types.ErrSmallerThanMinimumAmount, "%s is smaller than %s", msg.FarmingCoin.Amount, liquidFarm.MinimumFarmAmount)
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

	poolCoinDenom := liquiditytypes.PoolCoinDenom(msg.PoolId)
	lfCoinDenom := types.LiquidFarmCoinDenom(msg.PoolId)
	lfCoinTotalSupplyAmt := k.bankKeeper.GetSupply(ctx, types.LiquidFarmCoinDenom(msg.PoolId)).Amount
	lpCoinTotalStakedAmt := k.farmingKeeper.GetAllStakedCoinsByFarmer(ctx, reserveAddr).AmountOf(poolCoinDenom)
	lpCoinTotalQueuedAmt := k.farmingKeeper.GetAllQueuedCoinsByFarmer(ctx, reserveAddr).AmountOf(poolCoinDenom)

	mintingAmt := types.CalculateFarmMintingAmount(lfCoinTotalSupplyAmt, lpCoinTotalStakedAmt, lpCoinTotalQueuedAmt, msg.FarmingCoin.Amount)
	mintingCoins := sdk.NewCoins(sdk.NewCoin(lfCoinDenom, mintingAmt))

	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, mintingCoins); err != nil {
		return err
	}

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, msg.GetFarmer(), mintingCoins); err != nil {
		return err
	}

	if err := k.farmingKeeper.Stake(ctx, reserveAddr, sdk.NewCoins(msg.FarmingCoin)); err != nil {
		return err
	}

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
	_, found := k.liquidityKeeper.GetPool(ctx, poolId)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "pool %d not found", poolId)
	}

	lfCoinBalance := k.bankKeeper.SpendableCoins(ctx, farmer).AmountOf(types.LiquidFarmCoinDenom(poolId))
	if lfCoinBalance.LT(unfarmingCoin.Amount) {
		return sdkerrors.Wrapf(types.ErrInsufficientUnfarmingAmount, "%s is smaller than %s", lfCoinBalance, unfarmingCoin.Amount)
	}

	return nil
}

// Unfarm handles types.MsgUnfarm to unfarm LFCoin.
func (k Keeper) Unfarm(ctx sdk.Context, poolId uint64, farmer sdk.AccAddress, unfarmingCoin sdk.Coin) (UnfarmInfo, error) {
	if err := k.ValidateMsgUnfarm(ctx, poolId, farmer, unfarmingCoin); err != nil {
		return UnfarmInfo{}, err
	}

	reserveAddr := types.LiquidFarmReserveAddress(poolId)
	poolCoinDenom := liquiditytypes.PoolCoinDenom(poolId)
	lfCoinTotalSupplyAmt := k.bankKeeper.GetSupply(ctx, types.LiquidFarmCoinDenom(poolId)).Amount
	lpCoinTotalStakedAmt := k.farmingKeeper.GetAllStakedCoinsByFarmer(ctx, reserveAddr).AmountOf(poolCoinDenom)
	lpCoinTotalQueuedAmt := k.farmingKeeper.GetAllQueuedCoinsByFarmer(ctx, reserveAddr).AmountOf(poolCoinDenom)
	compoundingRewards, found := k.GetCompoundingRewards(ctx, poolId)
	if !found {
		compoundingRewards.Amount = sdk.ZeroInt()
	}

	_, found = k.GetLiquidFarm(ctx, poolId)
	if !found {
		if !lpCoinTotalStakedAmt.IsZero() || !lpCoinTotalQueuedAmt.IsZero() {
			panic(fmt.Errorf("unexpected amount; staked amount: %s; queued amount: %s", lpCoinTotalStakedAmt, lpCoinTotalQueuedAmt))
		}
		// Handle a case when liquid farm is removed in params
		// Since the reserve account must have unstaked all staked coins from the farming module,
		// the module must use the reserve account balance (staked + queued) and make queued amount zero
		lpCoinTotalStakedAmt = k.bankKeeper.SpendableCoins(ctx, reserveAddr).AmountOf(poolCoinDenom)
	}

	unfarmAmt := types.CalculateUnfarmAmount(
		lfCoinTotalSupplyAmt,
		lpCoinTotalStakedAmt,
		lpCoinTotalQueuedAmt,
		unfarmingCoin.Amount,
		compoundingRewards.Amount,
	)
	unfarmCoin := sdk.NewCoin(poolCoinDenom, unfarmAmt)

	// Unstake unfarm coin in the farming module and release it to the farmer
	if err := k.farmingKeeper.Unstake(ctx, reserveAddr, sdk.NewCoins(unfarmCoin)); err != nil {
		return UnfarmInfo{}, err
	}

	if err := k.bankKeeper.SendCoins(ctx, reserveAddr, farmer, sdk.NewCoins(unfarmCoin)); err != nil {
		return UnfarmInfo{}, err
	}

	// Burn the unfarming LFCoin by sending it to module account
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, farmer, types.ModuleName, sdk.NewCoins(unfarmingCoin)); err != nil {
		return UnfarmInfo{}, err
	}

	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(unfarmingCoin)); err != nil {
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
		return err
	}

	spendable := k.bankKeeper.SpendableCoins(ctx, msg.GetFarmer()).AmountOf(unfarmInfo.UnfarmCoin.Denom)
	if spendable.LT(unfarmInfo.UnfarmCoin.Amount) {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "%s is smaller than %s", spendable, unfarmInfo.UnfarmCoin.Amount)
	}

	_, err = k.liquidityKeeper.Withdraw(ctx, &liquiditytypes.MsgWithdraw{
		PoolId:     msg.PoolId,
		Withdrawer: msg.Farmer,
		PoolCoin:   unfarmInfo.UnfarmCoin,
	})
	if err != nil {
		return err
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

// RemoveLiquidFarm unstakes all staked pool coins from the farming module and
// remove the liquid farm object in the store
func (k Keeper) RemoveLiquidFarm(ctx sdk.Context, liquidFarm types.LiquidFarm) {
	reserveAddr := types.LiquidFarmReserveAddress(liquidFarm.PoolId)
	stakedCoins := k.farmingKeeper.GetAllStakedCoinsByFarmer(ctx, reserveAddr)
	if !stakedCoins.IsZero() {
		// Unstake all staked coins so that there will be no rewards accumulating
		if err := k.farmingKeeper.Unstake(ctx, reserveAddr, stakedCoins); err != nil {
			panic(err)
		}
	}

	auctionId := k.GetLastRewardsAuctionId(ctx, liquidFarm.PoolId)
	auction, found := k.GetRewardsAuction(ctx, liquidFarm.PoolId, auctionId)
	if !found {
		panic(fmt.Errorf("rewards auction %d must exist, but somehow it is not found", auctionId))
	}

	if err := k.RefundAllBids(ctx, auction, types.Bid{}); err != nil {
		panic(err)
	}

	auction.SetStatus(types.AuctionStatusFinished)
	k.SetRewardsAuction(ctx, auction)
	k.SetCompoundingRewards(ctx, liquidFarm.PoolId, types.CompoundingRewards{Amount: sdk.ZeroInt()})
	k.DeleteWinningBid(ctx, liquidFarm.PoolId, auctionId)
	k.DeleteLiquidFarm(ctx, liquidFarm)
}
