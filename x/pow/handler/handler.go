package handler

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ibtcore/ibtcore/x/pow/keeper"
)

// NewHandler returns a minimal handler function for the pow module.
// It intentionally avoids referencing sdk.Handler or sdkerrors so it compiles
// across SDK versions. Replace with real message handling later.
func NewHandler(k keeper.Keeper) func(sdk.Context, sdk.Msg) (*sdk.Result, error) {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		return nil, fmt.Errorf("pow module: no message handlers implemented")
	}
}

// New is an alias for NewHandler to match different wiring styles.
func New(k keeper.Keeper) func(sdk.Context, sdk.Msg) (*sdk.Result, error) {
	return NewHandler(k)
}
