package keeper

import (
	"encoding/json"
	"fmt"

	"SmplChain/x/roles/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace
	}
)

type KeeperI interface {
	HasRole(ctx sdk.Context, senderModule string, account sdk.AccAddress, role string) (bool, error)
	IsAdmin(ctx sdk.Context, senderModule string, account sdk.AccAddress) (bool, error)
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) HasRole(ctx sdk.Context, senderModule string, account sdk.AccAddress, role string) (bool, error) {

	var isRoleExist bool
	isRoleExist = false

	isExist := ctx.KVStore(k.storeKey).Get([]byte(account.String() + role))

	if len(isExist) > 0 {
		json.Unmarshal(isExist, &isRoleExist)

	}

	return isRoleExist, nil
}

func (k Keeper) IsAdmin(ctx sdk.Context, senderModule string, account sdk.AccAddress) (bool, error) {

	var isRoleExist bool
	isRoleExist = false

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(""))

	adminaccount := store.Get([]byte("admin"))

	if account.String() == string(adminaccount) {
		isRoleExist = true
	}

	// if !isRoleExist {
	// 	if account.String() == k.GetParams(ctx).Adminaccount {
	// 		isRoleExist = true
	// 	}

	// }

	return isRoleExist, nil
}

func (k Keeper) SetAdminAccount(ctx sdk.Context, account string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(""))
	key := []byte("admin")
	value := []byte(account)

	store.Set(key, value)
}

func (k Keeper) GetAdminAccount(ctx sdk.Context) string {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(""))
	key := []byte("admin")
	value := store.Get(key)
	return string(value)

}
