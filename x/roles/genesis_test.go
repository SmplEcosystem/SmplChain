package roles_test

import (
	"testing"

	keepertest "SmplChain/testutil/keeper"
	"SmplChain/testutil/nullify"
	"SmplChain/x/roles"
	"SmplChain/x/roles/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:       types.DefaultParams(),
		Adminaccount: "hhh",

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RolesKeeper(t)
	roles.InitGenesis(ctx, *k, genesisState)
	got := roles.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.Equal(t, got.Adminaccount, "hhh")

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
