package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter store keys
var (
	KeyLiquidFarms        = []byte("LiquidFarms")
	KeyAuctionPeriodHours = []byte("AuctionPeriodHours")

	DefaultLiquidFarms               = []LiquidFarm{}
	DefaultAuctionPeriodHours uint32 = 12
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return Params{
		LiquidFarms:        DefaultLiquidFarms,
		AuctionPeriodHours: DefaultAuctionPeriodHours,
	}
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyLiquidFarms, &p.LiquidFarms, validateLiquidFarms),
		paramtypes.NewParamSetPair(KeyAuctionPeriodHours, &p.AuctionPeriodHours, validateAuctionPeriodHours),
	}
}

// Validate validates the set of parameters
func (p Params) Validate() error {
	for _, v := range []struct {
		value     interface{}
		validator func(interface{}) error
	}{
		{p.LiquidFarms, validateLiquidFarms},
		{p.AuctionPeriodHours, validateAuctionPeriodHours},
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
		if err := liquidFarm.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func validateAuctionPeriodHours(i interface{}) error {
	_, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return nil
}
