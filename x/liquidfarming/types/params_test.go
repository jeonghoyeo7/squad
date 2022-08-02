package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"
)

func TestParams_Validate(t *testing.T) {
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
