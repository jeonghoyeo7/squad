package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"
	liquiditytypes "github.com/cosmosquad-labs/squad/v2/x/liquidity/types"
)

// Querier is used as Keeper will have duplicate methods if used directly, and gRPC names take precedence over keeper.
type Querier struct {
	Keeper
}

var _ types.QueryServer = Querier{}

// Params queries the parameters of the module.
func (k Querier) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	params := k.GetParams(ctx)
	return &types.QueryParamsResponse{Params: params}, nil
}

// LiquidFarms queries all liquidfarms.
func (k Querier) LiquidFarms(c context.Context, req *types.QueryLiquidFarmsRequest) (*types.QueryLiquidFarmsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	liquidFarmsRes := []types.LiquidFarmResponse{}
	for _, liquidFarm := range k.GetAllLiquidFarms(ctx) {
		reserveAddr := types.LiquidFarmReserveAddress(liquidFarm.PoolId)
		poolCoinDenom := liquiditytypes.PoolCoinDenom(liquidFarm.PoolId)
		queuedAmt := k.farmingKeeper.GetAllQueuedStakingAmountByFarmerAndDenom(ctx, reserveAddr, poolCoinDenom)
		stakedAmt := sdk.ZeroInt()
		staking, found := k.farmingKeeper.GetStaking(ctx, poolCoinDenom, reserveAddr)
		if found {
			stakedAmt = staking.Amount
		}

		liquidFarmsRes = append(liquidFarmsRes, types.LiquidFarmResponse{
			PoolId:                   liquidFarm.PoolId,
			LiquidFarmReserveAddress: reserveAddr.String(),
			LFCoinDenom:              types.LiquidFarmCoinDenom(liquidFarm.PoolId),
			MinFarmAmount:            liquidFarm.MinFarmAmount,
			MinBidAmount:             liquidFarm.MinBidAmount,
			QueuedCoin:               sdk.NewCoin(poolCoinDenom, queuedAmt),
			StakedCoin:               sdk.NewCoin(poolCoinDenom, stakedAmt),
		})
	}

	return &types.QueryLiquidFarmsResponse{LiquidFarms: liquidFarmsRes}, nil
}

// LiquidFarm queries the specific liquidfarm.
func (k Querier) LiquidFarm(c context.Context, req *types.QueryLiquidFarmRequest) (*types.QueryLiquidFarmResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if req.PoolId == 0 {
		return nil, status.Error(codes.InvalidArgument, "pool id cannot be 0")
	}

	ctx := sdk.UnwrapSDKContext(c)

	liquidFarm, found := k.GetLiquidFarm(ctx, req.PoolId)
	if !found {
		return nil, status.Errorf(codes.NotFound, "liquid farm by pool id %d not found", req.PoolId)
	}

	reserveAddr := types.LiquidFarmReserveAddress(liquidFarm.PoolId)
	poolCoinDenom := liquiditytypes.PoolCoinDenom(liquidFarm.PoolId)
	queuedAmt := k.farmingKeeper.GetAllQueuedStakingAmountByFarmerAndDenom(ctx, reserveAddr, poolCoinDenom)
	stakedAmt := sdk.ZeroInt()
	staking, found := k.farmingKeeper.GetStaking(ctx, poolCoinDenom, reserveAddr)
	if found {
		stakedAmt = staking.Amount
	}

	liquidFarmRes := types.LiquidFarmResponse{
		PoolId:                   liquidFarm.PoolId,
		LiquidFarmReserveAddress: reserveAddr.String(),
		LFCoinDenom:              types.LiquidFarmCoinDenom(liquidFarm.PoolId),
		MinFarmAmount:            liquidFarm.MinFarmAmount,
		MinBidAmount:             liquidFarm.MinBidAmount,
		QueuedCoin:               sdk.NewCoin(poolCoinDenom, queuedAmt),
		StakedCoin:               sdk.NewCoin(poolCoinDenom, stakedAmt),
	}

	return &types.QueryLiquidFarmResponse{LiquidFarm: liquidFarmRes}, nil
}

// RewardsAuctions queries all rewards auctions
func (k Querier) RewardsAuctions(c context.Context, req *types.QueryRewardsAuctionsRequest) (*types.QueryRewardsAuctionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if req.PoolId == 0 {
		return nil, status.Error(codes.InvalidArgument, "pool id cannot be 0")
	}

	ctx := sdk.UnwrapSDKContext(c)
	store := ctx.KVStore(k.storeKey)
	auctionStore := prefix.NewStore(store, types.RewardsAuctionKeyPrefix)

	// Filter auctions by descending order to show an ongoing auction first
	pageReq := &query.PageRequest{
		Reverse: true,
	}
	req.Pagination = pageReq

	var auctions []types.RewardsAuction
	pageRes, err := query.FilteredPaginate(auctionStore, req.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		auction, err := types.UnmarshalRewardsAuction(k.cdc, value)
		if err != nil {
			return false, err
		}

		if auction.PoolId != req.PoolId {
			return false, err
		}

		if accumulate {
			auctions = append(auctions, auction)
		}

		return true, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryRewardsAuctionsResponse{RewardAuctions: auctions, Pagination: pageRes}, nil
}

// RewardsAuction queries rewards auction
func (k Querier) RewardsAuction(c context.Context, req *types.QueryRewardsAuctionRequest) (*types.QueryRewardsAuctionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if req.PoolId == 0 {
		return nil, status.Error(codes.InvalidArgument, "pool id cannot be 0")
	}

	if req.AuctionId == 0 {
		return nil, status.Error(codes.InvalidArgument, "auction id cannot be 0")
	}

	ctx := sdk.UnwrapSDKContext(c)

	auction, found := k.GetRewardsAuction(ctx, req.PoolId, req.AuctionId)
	if !found {
		return nil, status.Errorf(codes.NotFound, "auction with pool %d and auction %d doesn't exist", req.PoolId, req.AuctionId)
	}

	return &types.QueryRewardsAuctionResponse{RewardAuction: auction}, nil
}

// Bids queries all bids.
func (k Querier) Bids(c context.Context, req *types.QueryBidsRequest) (*types.QueryBidsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.PoolId == 0 {
		return nil, status.Error(codes.InvalidArgument, "pool id cannot be 0")
	}

	ctx := sdk.UnwrapSDKContext(c)
	store := ctx.KVStore(k.storeKey)
	bidStore := prefix.NewStore(store, types.BidKeyPrefix)

	var bids []types.Bid
	pageRes, err := query.FilteredPaginate(bidStore, req.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		bid, err := types.UnmarshalBid(k.cdc, value)
		if err != nil {
			return false, err
		}

		if bid.PoolId != req.PoolId {
			return false, nil
		}

		if accumulate {
			bids = append(bids, bid)
		}

		return true, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryBidsResponse{Bids: bids, Pagination: pageRes}, nil
}
