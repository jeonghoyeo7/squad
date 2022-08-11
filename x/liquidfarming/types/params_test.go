package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"
)

func TestParams_Validate(t *testing.T) {
	require.IsType(t, paramstypes.KeyTable{}, types.ParamKeyTable())

	for _, tc := range []struct {
		name     string
		malleate func(*types.Params)
		errStr   string
	}{
		{
			"default params",
			func(params *types.Params) {},
			"",
		},
		{
			"default params",
			func(params *types.Params) {
				params.LiquidFarms = []types.LiquidFarm{
					types.NewLiquidFarm(1, sdk.ZeroInt(), sdk.ZeroInt()),
				}
				params.DelayedFarmGasFee = sdk.Gas(0)
			},
			"",
		},
		{
			"invalid params: pool id",
			func(params *types.Params) {
				params.LiquidFarms = []types.LiquidFarm{
					types.NewLiquidFarm(0, sdk.ZeroInt(), sdk.ZeroInt()),
				}
				params.DelayedFarmGasFee = sdk.Gas(0)
			},
			"pool id must not be 0",
		},
		{
			"negative UnfarmFeeRate",
			func(params *types.Params) {
				params.UnfarmFeeRate = sdk.NewDec(-1)
			},
			"unfarm fee rate must not be negative: -1.000000000000000000",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			params := types.DefaultParams()
			tc.malleate(&params)
			err := params.Validate()
			if tc.errStr == "" {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tc.errStr)
			}
		})
	}
}
