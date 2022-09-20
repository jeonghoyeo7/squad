package liquidfarming

import (
	"sort"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmosquad-labs/squad/v3/x/liquidfarming/keeper"
	"github.com/cosmosquad-labs/squad/v3/x/liquidfarming/types"
)

// BeginBlocker compares all LiquidFarms stored in the store with all LiquidFarms registered in params.
// Execute an appropriate operation when either adding new LiquidFarm or removing one through governance proposal.
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	// TODO: create RewardsAuction
	for _, liquidFarm := range k.GetAllLiquidFarms(ctx) {
		auctionId := k.GetLastRewardsAuctionId(ctx, liquidFarm.PoolId)
		auction, found := k.GetRewardsAuction(ctx, liquidFarm.PoolId, auctionId)
		if found {
			if err := k.FinishRewardsAuction(ctx, auction, liquidFarm.FeeRate); err != nil {
				panic(err)
			}
		}
		k.CreateRewardsAuction(ctx, liquidFarm.PoolId)
	}

	///////////////////////////////

	liquidFarmByPoolId := map[uint64]types.LiquidFarm{} // PoolId => LiquidFarm
	for _, liquidFarm := range k.GetAllLiquidFarms(ctx) {
		liquidFarmByPoolId[liquidFarm.PoolId] = liquidFarm
	}

	// Compare all liquid farms stored in KVStore with the ones registered in params
	// If new liquid farm is added through governance proposal, store it in KVStore.
	// Otherwise, delete from the liquidFarmByPoolId
	for _, liquidFarm := range k.GetParams(ctx).LiquidFarms {
		_, found := liquidFarmByPoolId[liquidFarm.PoolId]
		if !found { // new LiquidFarm is added
			k.SetLiquidFarm(ctx, liquidFarm)
		} else {
			delete(liquidFarmByPoolId, liquidFarm.PoolId)
		}
	}

	// Sort map keys for deterministic execution
	var poolIds []uint64
	for poolId := range liquidFarmByPoolId {
		poolIds = append(poolIds, poolId)
	}
	sort.Slice(poolIds, func(i, j int) bool {
		return poolIds[i] < poolIds[j]
	})
	for _, poolId := range poolIds {
		k.HandleRemovedLiquidFarm(ctx, liquidFarmByPoolId[poolId])
	}
}
