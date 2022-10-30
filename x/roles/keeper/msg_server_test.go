package keeper_test

import (
	"context"
	"testing"

	keepertest "SmplChain/testutil/keeper"
	"SmplChain/x/roles/keeper"
	"SmplChain/x/roles/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.RolesKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
