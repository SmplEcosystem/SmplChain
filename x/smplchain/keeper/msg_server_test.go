package keeper_test

import (
	"context"
	"testing"

	keepertest "SmplChain/testutil/keeper"
	"SmplChain/x/smplchain/keeper"
	"SmplChain/x/smplchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.SmplchainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
