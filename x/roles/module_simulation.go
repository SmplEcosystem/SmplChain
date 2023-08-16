package roles

import (
	"math/rand"

	"SmplChain/testutil/sample"
	rolessimulation "SmplChain/x/roles/simulation"
	"SmplChain/x/roles/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = rolessimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgAdd = "op_weight_msg_add"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAdd int = 100

	opWeightMsgRemove = "op_weight_msg_remove"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRemove int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	rolesGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&rolesGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgAdd int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAdd, &weightMsgAdd, nil,
		func(_ *rand.Rand) {
			weightMsgAdd = defaultWeightMsgAdd
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAdd,
		rolessimulation.SimulateMsgAdd(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRemove int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRemove, &weightMsgRemove, nil,
		func(_ *rand.Rand) {
			weightMsgRemove = defaultWeightMsgRemove
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRemove,
		rolessimulation.SimulateMsgRemove(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
