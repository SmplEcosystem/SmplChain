package roles

import (
	"SmplChain/x/roles/keeper"
	"SmplChain/x/roles/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
	k.SetAdminAccount(ctx, genState.Adminaccount)

}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.Adminaccount = k.GetAdminAccount(ctx)

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
