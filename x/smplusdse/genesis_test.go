package smplusdse_test

import (
	"testing"

	keepertest "SmplChain/testutil/keeper"
	"SmplChain/testutil/nullify"
	"SmplChain/x/smplusdse"
	"SmplChain/x/smplusdse/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SmplusdseKeeper(t)
	smplusdse.InitGenesis(ctx, *k, genesisState)
	got := smplusdse.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
