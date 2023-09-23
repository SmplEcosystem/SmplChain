package app

import (
	rolestypes "SmplChain/x/roles/types"
	smplchaintypes "SmplChain/x/smplchain/types"
	smplusdsetypes "SmplChain/x/smplusdse/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

// UpgradeName defines the on-chain upgrade name for the sample SimApp upgrade
// from v046 to v047.
//
// NOTE: This upgrade defines a reference implementation of what an upgrade
// could look like when an application is migrating from Cosmos SDK version
// v0.46.x to v0.47.x.
const UpgradeName = "v0.46-to-v047"

func (app *App) RegisterUpgradeHandlers() {
	// Set param key table for params module migration
	for _, subspace := range app.ParamsKeeper.GetSubspaces() {
		subspace := subspace

		var keyTable paramstypes.KeyTable
		switch subspace.Name() {
		case authtypes.ModuleName:
			keyTable = authtypes.ParamKeyTable() //nolint:staticcheck
		case banktypes.ModuleName:
			keyTable = banktypes.ParamKeyTable() //nolint:staticcheck
		case stakingtypes.ModuleName:
			keyTable = stakingtypes.ParamKeyTable() //nolint:staticcheck
		case minttypes.ModuleName:
			keyTable = minttypes.ParamKeyTable() //nolint:staticcheck
		case distrtypes.ModuleName:
			keyTable = distrtypes.ParamKeyTable() //nolint:staticcheck
		case slashingtypes.ModuleName:
			keyTable = slashingtypes.ParamKeyTable() //nolint:staticcheck
		case govtypes.ModuleName:
			keyTable = govv1.ParamKeyTable() //nolint:staticcheck
		case crisistypes.ModuleName:
			keyTable = crisistypes.ParamKeyTable() //nolint:staticcheck
		case smplchaintypes.ModuleName:
			keyTable = smplchaintypes.ParamKeyTable()
		case rolestypes.ModuleName:
			keyTable = rolestypes.ParamKeyTable()
		case smplusdsetypes.ModuleName:
			keyTable = smplusdsetypes.ParamKeyTable()
		}

		//statelessModules := []string{"smplusdse", "roles", "smplchain"}

		if !subspace.HasKeyTable() {
			app.Logger().Info("keyTable", "keytable", keyTable, "subspace", subspace, "subpace.name", subspace.Name())
			subspace.WithKeyTable(keyTable)
		}
	}

	baseAppLegacySS := app.ParamsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramstypes.ConsensusParamsKeyTable())

	app.UpgradeKeeper.SetUpgradeHandler(
		UpgradeName,
		func(ctx sdk.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			// Migrate Tendermint consensus parameters from x/params module to a dedicated x/consensus module.
			baseapp.MigrateParams(ctx, baseAppLegacySS, &app.ConsensusParamsKeeper)

			// Note: this migration is optional,
			// You can include x/gov proposal migration documented in [UPGRADING.md](https://github.com/cosmos/cosmos-sdk/blob/main/UPGRADING.md)

			return app.mm.RunMigrations(ctx, app.Configurator(), fromVM)
		},
	)

	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(err)
	}

	if upgradeInfo.Name == UpgradeName && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := storetypes.StoreUpgrades{
			Added: []string{
				consensustypes.ModuleName,
				crisistypes.ModuleName,
			},
		}

		// configure store loader that checks if version == upgradeHeight and applies store upgrades
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}
}
