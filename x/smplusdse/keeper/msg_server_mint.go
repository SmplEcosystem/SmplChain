package keeper

import (
	"context"
	"errors"

	"SmplChain/x/smplusdse/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Mint(goCtx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	coin, err := sdk.ParseCoinsNormalized(msg.Amount + "usdse")
	if err != nil {
		return nil, err
	}

	user, err := sdk.AccAddressFromBech32(msg.Creator)

	isBanker, err := k.rolesKeeper.HasRole(ctx, types.ModuleName, user, "bank")
	if !isBanker {
		return &types.MsgMintResponse{}, errors.New("bank role missing")
	}

	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, coin)
	if err != nil {
		return nil, err
	}

	k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, user, coin)
	return &types.MsgMintResponse{}, nil
}
