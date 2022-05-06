package v2_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/cosmosquad-labs/squad/app"
	"github.com/cosmosquad-labs/squad/x/mint/legacy/v2"
	"github.com/cosmosquad-labs/squad/x/mint/types"
)

func TestStoreMigration(t *testing.T) {
	encCfg := app.MakeTestEncodingConfig()
	key := sdk.NewKVStoreKey(types.ModuleName)
	tKey := sdk.NewTransientStoreKey("transient_test")
	ctx := testutil.DefaultContext(key, tKey)
	paramstore := paramtypes.NewSubspace(encCfg.Marshaler, encCfg.Amino, key, tKey, types.ModuleName)

	// Check no params
	require.False(t, paramstore.Has(ctx, types.KeyMintPoolAddress))

	// Run migrations.
	err := v2.MigrateStore(ctx, paramstore)
	require.NoError(t, err)

	// Make sure the new params are set.
	require.True(t, paramstore.Has(ctx, types.KeyMintPoolAddress))
}
