package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"
)

// NewDecodeStore returns a decoder function closure that unmarshals the KVPair's
// Value to the corresponding liquidfarming type.
func NewDecodeStore(cdc codec.Codec) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.Equal(kvA.Key[:1], types.LiquidFarmKeyPrefix):
			var lA, lB types.LiquidFarm
			cdc.MustUnmarshal(kvA.Value, &lA)
			cdc.MustUnmarshal(kvB.Value, &lB)
			return fmt.Sprintf("%v\n%v", lA, lB)

			// TODO: not implemented yet

		default:
			panic(fmt.Sprintf("invalid liquid farm key prefix %X", kvA.Key[:1]))
		}
	}
}
