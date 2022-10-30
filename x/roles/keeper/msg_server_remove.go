package keeper

import (
	"context"

	"SmplChain/x/roles/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Remove(goCtx context.Context, msg *types.MsgRemove) (*types.MsgRemoveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRemoveResponse{}, nil
}
