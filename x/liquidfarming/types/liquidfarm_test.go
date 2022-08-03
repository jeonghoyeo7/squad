package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"
)

func TestLiquidFarm(t *testing.T) {
	liquidFarm := types.LiquidFarm{
		PoolId:            1,
		MinimumFarmAmount: sdk.ZeroInt(),
		MinimumBidAmount:  sdk.ZeroInt(),
	}
	require.Equal(t, `minimum_bid_amount: "0"
minimum_farm_amount: "0"
pool_id: "1"
`, liquidFarm.String())
}

func TestQueuedFarming(t *testing.T) {
	msg := types.QueuedFarming{
		PoolId: 1,
		Amount: sdk.NewInt(100_000_000),
	}

	registry := codectypes.NewInterfaceRegistry()
	types.RegisterInterfaces(registry)
	cdc := codec.NewProtoCodec(registry)

	bz, err := types.MarshalQueuedFarming(cdc, msg)
	require.NoError(t, err)

	queuedFarming, err := types.UnmarshalQueuedFarming(cdc, bz)
	require.NoError(t, err)

	require.EqualValues(t, msg, queuedFarming)
}
