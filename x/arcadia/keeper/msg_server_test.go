package keeper_test

import (
	"context"
	"testing"

	keepertest "arcadia/testutil/keeper"
	"arcadia/x/arcadia/keeper"
	"arcadia/x/arcadia/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ArcadiaKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
