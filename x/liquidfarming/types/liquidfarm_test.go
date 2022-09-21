package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	utils "github.com/cosmosquad-labs/squad/v3/types"
	"github.com/cosmosquad-labs/squad/v3/x/liquidfarming/types"
)

func TestLiquidFarm(t *testing.T) {
	liquidFarm := types.LiquidFarm{
		PoolId:        1,
		MinFarmAmount: sdk.ZeroInt(),
		MinBidAmount:  sdk.ZeroInt(),
		FeeRate:       sdk.ZeroDec(),
	}
	require.Equal(t, `fee_rate: "0.000000000000000000"
min_bid_amount: "0"
min_farm_amount: "0"
pool_id: "1"
`, liquidFarm.String())
}

func TestLiquidFarmCoinDenom(t *testing.T) {
	for _, tc := range []struct {
		denom      string
		expectsErr bool
	}{
		{"lf1", false},
		{"lf10", false},
		{"lf18446744073709551615", false},
		{"lf18446744073709551616", true},
		{"lfabc", true},
		{"lf01", true},
		{"lf-10", true},
		{"lf+10", true},
		{"ucre", true},
		{"denom1", true},
	} {
		t.Run("", func(t *testing.T) {
			poolId, err := types.ParseLiquidFarmCoinDenom(tc.denom)
			if tc.expectsErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.denom, types.LiquidFarmCoinDenom(poolId))
			}
		})
	}
}

func TestLiquidFarmReserveAddress(t *testing.T) {
	config := sdk.GetConfig()
	addrPrefix := config.GetBech32AccountAddrPrefix()

	for _, tc := range []struct {
		poolId   uint64
		expected string
	}{
		{1, addrPrefix + "1zyyf855slxure4c8dr06p00qjnkem95d2lgv8wgvry2rt437x6tsaf9tcf"},
		{2, addrPrefix + "1d2csu4ynxpuxll8wk72n9z98ytm649u78paj9efskjwrlc2wyhpq8h886j"},
	} {
		t.Run("", func(t *testing.T) {
			require.Equal(t, tc.expected, types.LiquidFarmReserveAddress(tc.poolId).String())
		})
	}
}

func TestCalculateMintingFarmAmount(t *testing.T) {
	for _, tc := range []struct {
		name              string
		lfTotalSupplyAmt  sdk.Int
		lpTotalFarmingAmt sdk.Int
		newFarmingAmt     sdk.Int
		expectedAmt       sdk.Int
	}{
		{
			name:              "initial minting",
			lfTotalSupplyAmt:  sdk.ZeroInt(),
			lpTotalFarmingAmt: sdk.ZeroInt(),
			newFarmingAmt:     sdk.NewInt(1_000_00_000),
			expectedAmt:       sdk.NewInt(1_000_00_000),
		},
		{
			name:              "case #1",
			lfTotalSupplyAmt:  sdk.NewInt(1_000_000_000),
			lpTotalFarmingAmt: sdk.NewInt(2_000_000_000),
			newFarmingAmt:     sdk.NewInt(1_000_000_000),
			expectedAmt:       sdk.NewInt(500_000_000),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			mintingAmt := types.CalculateFarmMintingAmount(
				tc.lfTotalSupplyAmt,
				tc.lpTotalFarmingAmt,
				tc.newFarmingAmt,
			)
			require.Equal(t, tc.expectedAmt, mintingAmt)
		})
	}
}

func TestCalculateUnfarmAmount(t *testing.T) {
	for _, tc := range []struct {
		name               string
		lfTotalSupplyAmt   sdk.Int
		lpTotalFarmingAmt  sdk.Int
		unfarmingAmt       sdk.Int
		compoundingRewards sdk.Int
		expectedAmt        sdk.Int
	}{
		{
			name:               "supply equals to unfarming amount",
			lfTotalSupplyAmt:   sdk.NewInt(100_000_000),
			lpTotalFarmingAmt:  sdk.NewInt(50_000_000),
			unfarmingAmt:       sdk.NewInt(100_000_000),
			compoundingRewards: sdk.ZeroInt(),
			expectedAmt:        sdk.NewInt(100_000_000),
		},
		{
			name:               "small unfarming amount",
			lfTotalSupplyAmt:   sdk.NewInt(100_000_000),
			lpTotalFarmingAmt:  sdk.NewInt(50_000_000),
			unfarmingAmt:       sdk.NewInt(1),
			compoundingRewards: sdk.ZeroInt(),
			expectedAmt:        sdk.NewInt(1),
		},
		{
			name:               "case #1: bidding amount is auto staked",
			lfTotalSupplyAmt:   sdk.NewInt(2000000000),
			lpTotalFarmingAmt:  sdk.NewInt(2200000000),
			unfarmingAmt:       sdk.NewInt(1000000000),
			compoundingRewards: sdk.NewInt(30000000),
			expectedAmt:        sdk.NewInt(1100000000),
		},
		{
			name:               "case #2: bidding amount is auto staked",
			lfTotalSupplyAmt:   sdk.NewInt(1000000000),
			lpTotalFarmingAmt:  sdk.NewInt(1130000000),
			unfarmingAmt:       sdk.NewInt(1000000000),
			compoundingRewards: sdk.NewInt(30000000),
			expectedAmt:        sdk.NewInt(1130000000),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			unfarmingAmt := types.CalculateUnfarmingAmount(
				tc.lfTotalSupplyAmt,
				tc.lpTotalFarmingAmt,
				tc.unfarmingAmt,
				tc.compoundingRewards,
			)
			require.Equal(t, tc.expectedAmt, unfarmingAmt)
		})
	}
}

func TestDeductFees(t *testing.T) {
	for _, tc := range []struct {
		name     string
		feeRate  sdk.Dec
		rewards  sdk.Coins
		deducted sdk.Coins
	}{
		{
			name:     "zero fee rate",
			feeRate:  sdk.ZeroDec(),
			rewards:  utils.ParseCoins("100denom1"),
			deducted: utils.ParseCoins("100denom1"),
		},
		{
			name:     "fee rate - 10%",
			feeRate:  sdk.MustNewDecFromStr("0.1"),
			rewards:  utils.ParseCoins("100denom1"),
			deducted: utils.ParseCoins("90denom1"),
		},
		{
			name:     "fee rate - 0.066666666666666",
			feeRate:  sdk.MustNewDecFromStr("0.066666666666666"),
			rewards:  utils.ParseCoins("100000denom1"),
			deducted: utils.ParseCoins("93333denom1"),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			deducted, err := types.DeductFees(1, tc.rewards, tc.feeRate)
			require.NoError(t, err)
			require.Equal(t, tc.deducted, deducted)
		})
	}
}
