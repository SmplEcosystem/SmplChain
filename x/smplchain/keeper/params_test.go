package keeper_test

import (
	"testing"

	testkeeper "SmplChain/testutil/keeper"
	"SmplChain/x/smplchain/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.SmplchainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
