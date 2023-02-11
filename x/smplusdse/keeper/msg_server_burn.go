package keeper

import (
	"context"
	"errors"

	"SmplChain/x/smplusdse/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Burn(goCtx context.Context, msg *types.MsgBurn) (*types.MsgBurnResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	user, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	coin, err := sdk.ParseCoinsNormalized(msg.Amount + "busdse")
	if err != nil {
		return nil, err
	}

	if k.bankKeeper.SpendableCoins(ctx, user).IsAllGTE(coin) {
		k.bankKeeper.SendCoinsFromAccountToModule(ctx, user, types.ModuleName, coin)
		err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, coin)
		if err != nil {
			return nil, err
		}

	} else {
		return nil, errors.New("not enough balance")
	}

	return &types.MsgBurnResponse{}, nil
}
