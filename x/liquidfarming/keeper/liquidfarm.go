package keeper

import (
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmosquad-labs/squad/v3/x/liquidfarming/types"
	liquiditytypes "github.com/cosmosquad-labs/squad/v3/x/liquidity/types"
)

// LiquidFarm handles types.MsgFarm to liquid farm.
func (k Keeper) LiquidFarm(ctx sdk.Context, poolId uint64, farmer sdk.AccAddress, farmingCoin sdk.Coin) error {
	pool, found := k.liquidityKeeper.GetPool(ctx, poolId)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "pool %d not found", poolId)
	}

	liquidFarm, found := k.GetLiquidFarm(ctx, pool.Id)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "liquid farm by pool %d not found", pool.Id)
	}

	if farmingCoin.Amount.LT(liquidFarm.MinFarmAmount) {
		return sdkerrors.Wrapf(types.ErrSmallerThanMinimumAmount, "%s is smaller than %s", farmingCoin.Amount, liquidFarm.MinFarmAmount)
	}

	reserveAddr := types.LiquidFarmReserveAddress(pool.Id)
	if err := k.bankKeeper.SendCoins(ctx, farmer, reserveAddr, sdk.NewCoins(farmingCoin)); err != nil {
		return err
	}

	lfCoinDenom := types.LiquidFarmCoinDenom(pool.Id)
	lfCoinTotalSupplyAmt := k.bankKeeper.GetSupply(ctx, lfCoinDenom).Amount
	lpCoinTotalFarmingAmt := sdk.ZeroInt()
	farm, found := k.farmKeeper.GetFarm(ctx, farmingCoin.Denom)
	if found {
		lpCoinTotalFarmingAmt = farm.TotalFarmingAmount
	}

	mintingAmt := types.CalculateFarmMintingAmount(
		lfCoinTotalSupplyAmt,
		lpCoinTotalFarmingAmt,
		farmingCoin.Amount,
	)
	mintingCoin := sdk.NewCoin(lfCoinDenom, mintingAmt)

	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(mintingCoin)); err != nil {
		return err
	}

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, farmer, sdk.NewCoins(mintingCoin)); err != nil {
		return err
	}

	if _, err := k.farmKeeper.Farm(ctx, reserveAddr, farmingCoin); err != nil {
		return err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeFarm,
			sdk.NewAttribute(types.AttributeKeyPoolId, strconv.FormatUint(poolId, 10)),
			sdk.NewAttribute(types.AttributeKeyFarmer, farmer.String()),
			sdk.NewAttribute(types.AttributeKeyFarmingCoin, farmingCoin.String()),
			sdk.NewAttribute(types.AttributeKeyMintedCoin, mintingCoin.String()),
			sdk.NewAttribute(types.AttributeKeyLiquidFarmReserveAddress, reserveAddr.String()),
		),
	})

	return nil
}

// UnfarmInfo holds information about unfarm.
type UnfarmInfo struct {
	Farmer       sdk.AccAddress
	UnfarmedCoin sdk.Coin
}

// LiquidUnfarm handles types.MsgUnfarm to unfarm LFCoin.
// It doesn't validate if the liquid farm exists because farmers still need to be able to
// unfarm their LFCoin although the liquid farm object is removed in params.
func (k Keeper) LiquidUnfarm(ctx sdk.Context, poolId uint64, farmer sdk.AccAddress, burningCoin sdk.Coin) (UnfarmInfo, error) {
	pool, found := k.liquidityKeeper.GetPool(ctx, poolId)
	if !found {
		return UnfarmInfo{}, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "pool %d not found", poolId)
	}

	reserveAddr := types.LiquidFarmReserveAddress(pool.Id)
	poolCoinDenom := liquiditytypes.PoolCoinDenom(pool.Id)
	lfCoinDenom := types.LiquidFarmCoinDenom(pool.Id)
	lfCoinTotalSupplyAmt := k.bankKeeper.GetSupply(ctx, lfCoinDenom).Amount
	lpCoinTotalFarmingAmt := sdk.ZeroInt()
	compoundingRewardsAmt := sdk.ZeroInt()
	farm, found := k.farmKeeper.GetFarm(ctx, poolCoinDenom)
	if found {
		lpCoinTotalFarmingAmt = farm.TotalFarmingAmount
	}
	compoundingRewards, found := k.GetCompoundingRewards(ctx, pool.Id)
	if found {
		compoundingRewardsAmt = compoundingRewards.Amount
	}

	_, found = k.GetLiquidFarm(ctx, poolId)
	if !found {
		// TODO: check if this is necessary any more
		if !lpCoinTotalFarmingAmt.IsZero() {
			panic(fmt.Errorf("unexpected lp coin total farming amount %s", lpCoinTotalFarmingAmt))
		}
		// Handle a case when liquid farm is removed in params
		// Since the reserve account must have unstaked all staked coins from the farming module,
		// the module must use the reserve account balance (staked + queued) and make queued amount zero
		lpCoinTotalFarmingAmt = k.bankKeeper.SpendableCoins(ctx, reserveAddr).AmountOf(poolCoinDenom)
	}

	unfarmingAmt := types.CalculateUnfarmingAmount(
		lfCoinTotalSupplyAmt,
		lpCoinTotalFarmingAmt,
		burningCoin.Amount,
		compoundingRewardsAmt,
	)
	unfarmingCoin := sdk.NewCoin(poolCoinDenom, unfarmingAmt)

	// TODO: check if this if statement is necessary any more
	if found {
		// Unstake unfarm coin in the farming module and release it to the farmer
		if _, err := k.farmKeeper.Unfarm(ctx, reserveAddr, unfarmingCoin); err != nil {
			return UnfarmInfo{}, err
		}
	}

	if err := k.bankKeeper.SendCoins(ctx, reserveAddr, farmer, sdk.NewCoins(unfarmingCoin)); err != nil {
		return UnfarmInfo{}, err
	}

	// Burn the unfarming LFCoin by sending it to module account
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, farmer, types.ModuleName, sdk.NewCoins(burningCoin)); err != nil {
		return UnfarmInfo{}, err
	}

	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(burningCoin)); err != nil {
		return UnfarmInfo{}, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUnfarm,
			sdk.NewAttribute(types.AttributeKeyPoolId, strconv.FormatUint(poolId, 10)),
			sdk.NewAttribute(types.AttributeKeyFarmer, farmer.String()),
			sdk.NewAttribute(types.AttributeKeyBurningCoin, burningCoin.String()),
			sdk.NewAttribute(types.AttributeKeyUnfarmedCoin, unfarmingCoin.String()),
		),
	})

	return UnfarmInfo{Farmer: farmer, UnfarmedCoin: unfarmingCoin}, nil
}

// LiquidUnfarmAndWithdraw handles types.MsgUnfarmAndWithdraw to unfarm LFCoin and withdraw pool coin from the pool.
func (k Keeper) LiquidUnfarmAndWithdraw(ctx sdk.Context, poolId uint64, farmer sdk.AccAddress, burningCoin sdk.Coin) error {
	unfarmInfo, err := k.LiquidUnfarm(ctx, poolId, farmer, burningCoin)
	if err != nil {
		return err
	}

	_, err = k.liquidityKeeper.Withdraw(ctx, &liquiditytypes.MsgWithdraw{
		PoolId:     poolId,
		Withdrawer: farmer.String(),
		PoolCoin:   unfarmInfo.UnfarmedCoin,
	})
	if err != nil {
		return err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUnfarmAndWithdraw,
			sdk.NewAttribute(types.AttributeKeyPoolId, strconv.FormatUint(poolId, 10)),
			sdk.NewAttribute(types.AttributeKeyFarmer, farmer.String()),
			sdk.NewAttribute(types.AttributeKeyBurningCoin, burningCoin.String()),
			sdk.NewAttribute(types.AttributeKeyUnfarmedCoin, unfarmInfo.UnfarmedCoin.String()),
		),
	})

	return nil
}

// HandleRemovedLiquidFarm unstakes all staked pool coins from the farm module and
// remove the liquid farm object in the store
func (k Keeper) HandleRemovedLiquidFarm(ctx sdk.Context, liquidFarm types.LiquidFarm) {
	reserveAddr := types.LiquidFarmReserveAddress(liquidFarm.PoolId)
	poolCoinDenom := liquiditytypes.PoolCoinDenom(liquidFarm.PoolId)
	farmingAmt := sdk.ZeroInt()
	position, found := k.farmKeeper.GetPosition(ctx, reserveAddr, poolCoinDenom)
	if found {
		farmingAmt = position.FarmingAmount
	}
	stakedCoin := sdk.NewCoin(poolCoinDenom, farmingAmt)
	if !stakedCoin.IsZero() {
		// Unstake all staked coins so that there will be no rewards accumulating
		if _, err := k.farmKeeper.Unfarm(ctx, reserveAddr, stakedCoin); err != nil {
			panic(err)
		}
	}

	// Handle a case when the last rewards auction id isn't set in the store
	auctionId := k.GetLastRewardsAuctionId(ctx, liquidFarm.PoolId)
	auction, found := k.GetRewardsAuction(ctx, liquidFarm.PoolId, auctionId)
	if found {
		if err := k.RefundAllBids(ctx, auction); err != nil {
			panic(err)
		}
		k.DeleteWinningBid(ctx, liquidFarm.PoolId, auctionId)
		auction.SetStatus(types.AuctionStatusFinished)
		k.SetRewardsAuction(ctx, auction)
	}

	k.SetCompoundingRewards(ctx, liquidFarm.PoolId, types.CompoundingRewards{Amount: sdk.ZeroInt()})
	k.DeleteLiquidFarm(ctx, liquidFarm)
}
