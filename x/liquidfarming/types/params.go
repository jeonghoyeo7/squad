package types

import (
	"fmt"
	time "time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter store keys
var (
	KeyLiquidFarms            = []byte("LiquidFarms")
	KeyRewardsAuctionDuration = []byte("RewardsAuctionDuration")
	KeyFeeCollector           = []byte("FeeCollector")
)

// Default parameters
var (
	DefaultLiquidFarms                          = []LiquidFarm{}
	DefaultRewardsAuctionDuration time.Duration = time.Hour * 12
	DefaultFeeCollector                         = sdk.AccAddress(address.Module(ModuleName, []byte("FeeCollector")))
)

var _ paramstypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return Params{
		LiquidFarms:            DefaultLiquidFarms,
		RewardsAuctionDuration: DefaultRewardsAuctionDuration,
		FeeCollector:           DefaultFeeCollector.String(),
	}
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(KeyLiquidFarms, &p.LiquidFarms, validateLiquidFarms),
		paramstypes.NewParamSetPair(KeyRewardsAuctionDuration, &p.RewardsAuctionDuration, validateRewardsAuctionDuration),
		paramstypes.NewParamSetPair(KeyFeeCollector, &p.FeeCollector, validateFeeCollector),
	}
}

// Validate validates the set of parameters
func (p Params) Validate() error {
	for _, v := range []struct {
		value     interface{}
		validator func(interface{}) error
	}{
		{p.LiquidFarms, validateLiquidFarms},
		{p.RewardsAuctionDuration, validateRewardsAuctionDuration},
		{p.FeeCollector, validateFeeCollector},
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
	for _, l := range liquidFarms {
		if err := l.Validate(); err != nil {
			return fmt.Errorf("invalid liquid farm: %v", err)
		}
	}
	return nil
}

func validateRewardsAuctionDuration(i interface{}) error {
	v, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v <= 0 {
		return fmt.Errorf("invalid rewards auction duration: %d must be positive value", v)
	}
	return nil
}

func validateFeeCollector(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	_, err := sdk.AccAddressFromBech32(v)
	if err != nil {
		return fmt.Errorf("invalid fee collector address: %v", v)
	}
	return nil
}
