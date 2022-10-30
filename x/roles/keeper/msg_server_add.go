package keeper

import (
	"context"

	"SmplChain/x/roles/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Add(goCtx context.Context, msg *types.MsgAdd) (*types.MsgAddResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgAddResponse{}, nil
}
