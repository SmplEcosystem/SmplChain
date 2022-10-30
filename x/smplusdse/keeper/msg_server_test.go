package keeper_test

import (
	"context"
	"testing"

	keepertest "SmplChain/testutil/keeper"
	"SmplChain/x/smplusdse/keeper"
	"SmplChain/x/smplusdse/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.SmplusdseKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
