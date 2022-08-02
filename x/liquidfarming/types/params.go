package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter store keys
var (
	KeyLiquidFarms       = []byte("LiquidFarms")
	KeyDelayedFarmGasFee = []byte("DelayedFarmGasFee")

	DefaultLiquidFarms       = []LiquidFarm{}
	DefaultDelayedFarmGasFee = sdk.Gas(60000)
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return Params{
		LiquidFarms:       DefaultLiquidFarms,
		DelayedFarmGasFee: DefaultDelayedFarmGasFee,
	}
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyLiquidFarms, &p.LiquidFarms, validateLiquidFarms),
		paramtypes.NewParamSetPair(KeyDelayedFarmGasFee, &p.DelayedFarmGasFee, validateDelayedFarmGasFee),
	}
}

// Validate validates the set of parameters
func (p Params) Validate() error {
	for _, v := range []struct {
		value     interface{}
		validator func(interface{}) error
	}{
		{p.LiquidFarms, validateLiquidFarms},
		{p.DelayedFarmGasFee, validateDelayedFarmGasFee},
	} {
		if err := v.validator(v.value); err != nil {
			return err
		}
	}
	return nil
}

func validateLiquidFarms(i interface{}) error {
	liquidFarms, ok := i.([]LiquidFarm)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	for _, liquidFarm := range liquidFarms {
		if liquidFarm.PoolId == 0 {
			return fmt.Errorf("pool id must not be 0")
		}
		if liquidFarm.MinimumBidAmount.IsNegative() {
			return fmt.Errorf("minimum bid amount must be 0 or positive value: %s", liquidFarm.MinimumBidAmount)
		}
		if liquidFarm.MinimumFarmAmount.IsNegative() {
			return fmt.Errorf("minimum farm amount must be 0 or positive value: %s", liquidFarm.MinimumFarmAmount)
		}
	}

	return nil
}

func validateDelayedFarmGasFee(i interface{}) error {
	_, ok := i.(sdk.Gas)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}
