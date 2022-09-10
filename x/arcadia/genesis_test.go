package arcadia_test

import (
	"testing"

	keepertest "arcadia/testutil/keeper"
	"arcadia/testutil/nullify"
	"arcadia/x/arcadia"
	"arcadia/x/arcadia/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ArcadiaKeeper(t)
	arcadia.InitGenesis(ctx, *k, genesisState)
	got := arcadia.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
