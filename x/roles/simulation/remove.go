package simulation

import (
	"math/rand"

	"SmplChain/x/roles/keeper"
	"SmplChain/x/roles/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgRemove(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgRemove{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the Remove simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Remove simulation not implemented"), nil, nil
	}
}
