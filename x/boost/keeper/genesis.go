package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmosquad-labs/squad/v2/x/boost/types"
)

// InitGenesis initializes the capability module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	if err := genState.Validate(); err != nil {
		panic(err)
	}

	k.SetParams(ctx, genState.Params)

	// for _, liquidFarm := range genState.LiquidFarms {
	// 	k.SetLiquidFarm(ctx, liquidFarm)
	// }

	// for _, record := range genState.LastRewardsAuctionIdRecord {
	// 	k.SetLastRewardsAuctionId(ctx, record.PoolId, record.AuctionId)
	// }

	// for _, auction := range genState.RewardsAuctions {
	// 	k.SetRewardsAuction(ctx, auction)
	// }

	// for _, bid := range genState.Bids {
	// 	k.SetBid(ctx, bid)
	// }

	// for _, record := range genState.WinningBidRecords {
	// 	k.SetWinningBid(ctx, record.WinningBid, record.AuctionId)
	// }
}

// ExportGenesis returns the module's exported genesis.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	params := k.GetParams(ctx)

	// Initialize objects to prevent from having nil slice
	if params.LiquidFarms == nil {
		params.LiquidFarms = []types.LiquidFarm{}
	}

	// poolIds := []uint64{}
	// for _, liquidFarm := range params.LiquidFarms {
	// 	poolIds = append(poolIds, liquidFarm.PoolId)
	// }

	return &types.GenesisState{
		Params: params,
		// LastRewardsAuctionIdRecord: lastRewardsAuctionIdRecords,
		// LiquidFarms:                k.GetAllLiquidFarms(ctx),
		// RewardsAuctions:            k.GetAllRewardsAuctions(ctx),
		// Bids:                       bids,
		// WinningBidRecords:          winningBidRecords,
	}
}
