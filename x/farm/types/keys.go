package types

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"

	utils "github.com/cosmosquad-labs/squad/v3/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "farm"

	// StoreKey defines the primary module store key
	StoreKey = "f4rm" // To avoid store key collision with "farming"

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName
)

var (
	LastPlanIdKey              = []byte{0xd0}
	LastBlockTimeKey           = []byte{0xd1}
	PlanKeyPrefix              = []byte{0xd2}
	FarmKeyPrefix              = []byte{0xd3}
	PositionKeyPrefix          = []byte{0xd4}
	HistoricalRewardsKeyPrefix = []byte{0xd5}
)

func GetPlanKey(id uint64) []byte {
	return append(PlanKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func GetFarmKey(denom string) []byte {
	return append(FarmKeyPrefix, denom...)
}

func GetPositionKey(farmerAddr sdk.AccAddress, denom string) []byte {
	return append(append(PositionKeyPrefix, address.MustLengthPrefix(farmerAddr)...), denom...)
}

func GetHistoricalRewardsKey(denom string, period uint64) []byte {
	return append(append(HistoricalRewardsKeyPrefix, utils.LengthPrefixString(denom)...), sdk.Uint64ToBigEndian(period)...)
}

func ParseFarmKey(key []byte) (denom string) {
	if !bytes.HasPrefix(key, FarmKeyPrefix) {
		panic("key does not have proper prefix")
	}
	denom = string(key[1:])
	return
}

func ParseHistoricalRewardsKey(key []byte) (denom string, period uint64) {
	if !bytes.HasPrefix(key, HistoricalRewardsKeyPrefix) {
		panic("key does not have proper prefix")
	}
	denomLen := key[1]
	denom = string(key[2 : 2+denomLen])
	period = sdk.BigEndianToUint64(key[2+denomLen:])
	return
}
