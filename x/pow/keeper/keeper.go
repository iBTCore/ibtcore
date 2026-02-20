package keeper

import (
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ibtcore/ibtcore/x/pow/types"
)

// Keeper defines the pow module Keeper
type Keeper struct {
	storeKey storetypes.StoreKey
	cdc      codec.BinaryCodec

	bankKeeper  types.BankKeeper // interface to bank keeper
	founderAddr sdk.AccAddress   // founder / treasury address
}

// NewKeeper creates a new Keeper instance
func NewKeeper(cdc codec.BinaryCodec, key storetypes.StoreKey, bk types.BankKeeper, founder sdk.AccAddress) Keeper {
	return Keeper{
		storeKey:    key,
		cdc:         cdc,
		bankKeeper:  bk,
		founderAddr: founder,
	}
}

// SetParams stores params (pointer) to avoid copying proto mutex
func (k Keeper) SetParams(ctx sdk.Context, params *types.Params) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(params)
	store.Set([]byte(types.ModuleName), bz)
}

// GetParams returns a pointer to Params (reads from KV store; falls back to defaults)
func (k Keeper) GetParams(ctx sdk.Context) *types.Params {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte(types.ModuleName))
	if bz == nil {
		dp := types.DefaultParams()
		return &dp
	}
	var params types.Params
	k.cdc.MustUnmarshal(bz, &params)
	return &params
}

// MintCoins mints coins into the module account (requires mint permission)
func (k Keeper) MintCoins(ctx sdk.Context, coins sdk.Coins) error {
	return k.bankKeeper.MintCoins(ctx, types.ModuleName, coins)
}

// SendCoinsFromModuleToAccount sends coins from a module account to an account
func (k Keeper) SendCoinsFromModuleToAccount(ctx sdk.Context, module string, addr sdk.AccAddress, coins sdk.Coins) error {
	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, module, addr, coins)
}

// SendCoinsFromAccountToModule sends coins from an account to a module account
func (k Keeper) SendCoinsFromAccountToModule(ctx sdk.Context, addr sdk.AccAddress, module string, coins sdk.Coins) error {
	return k.bankKeeper.SendCoinsFromAccountToModule(ctx, addr, module, coins)
}

// GetFounderAddress returns the configured founder address
func (k Keeper) GetFounderAddress(ctx sdk.Context) sdk.AccAddress {
	return k.founderAddr
}
