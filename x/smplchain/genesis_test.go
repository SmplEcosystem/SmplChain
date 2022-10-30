package smplchain_test

import (
	"testing"

	keepertest "SmplChain/testutil/keeper"
	"SmplChain/testutil/nullify"
	"SmplChain/x/smplchain"
	"SmplChain/x/smplchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SmplchainKeeper(t)
	smplchain.InitGenesis(ctx, *k, genesisState)
	got := smplchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
