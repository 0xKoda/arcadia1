package keeper_test

import (
	"testing"

	testkeeper "arcadia/testutil/keeper"
	"arcadia/x/arcadia/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ArcadiaKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
