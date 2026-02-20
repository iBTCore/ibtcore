package pow

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/ibtcore/ibtcore/x/pow/handler"
	"github.com/ibtcore/ibtcore/x/pow/keeper"
	"github.com/ibtcore/ibtcore/x/pow/types"
)

type AppModule struct {
	keeper keeper.Keeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(k keeper.Keeper) AppModule {
	return AppModule{
		keeper: k,
	}
}

func (am AppModule) Name() string { return types.ModuleName }

func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// Route returns the module message route name
func (am AppModule) Route() string { return types.RouterKey }

// InitGenesis uses encoding/json directly
func (am AppModule) InitGenesis(ctx sdk.Context, _ interface{}, data json.RawMessage) []abci.ValidatorUpdate {
	var gs types.GenesisState
	_ = json.Unmarshal(data, &gs)

	am.keeper.SetParams(ctx, &gs.Params)
	return []abci.ValidatorUpdate{}
}

// ExportGenesis uses encoding/json directly
func (am AppModule) ExportGenesis(ctx sdk.Context, _ interface{}) json.RawMessage {
	gs := types.GenesisState{Params: *am.keeper.GetParams(ctx)}
	bz, _ := json.Marshal(&gs)
	return bz
}

// NewHandler returns the module handler.
// Some SDK versions use the concrete type sdk.Handler; others use a function signature.
// To remain compatible with the handler stub we added, return the function signature.
func (am AppModule) NewHandler() func(sdk.Context, sdk.Msg) (*sdk.Result, error) {
	return handler.NewHandler(am.keeper)
}

func (am AppModule) BeginBlock(ctx sdk.Context, _ abci.RequestBeginBlock) {
	params := am.keeper.GetParams(ctx)
	reward := params.BlockReward
	if reward <= 0 {
		return
	}

	founderReward := reward * params.FounderShare / 100
	minerReward := reward * params.MinerShare / 100

	coins := sdk.NewCoins(sdk.NewInt64Coin("ibtc", reward))

	// Mint to module account
	if err := am.keeper.MintCoins(ctx, coins); err != nil {
		ctx.EventManager().EmitEvent(sdk.NewEvent(types.ModuleName,
			sdk.NewAttribute("action", "mint_failed"),
			sdk.NewAttribute("error", err.Error()),
		))
		return
	}

	// Send founder share
	if founderReward > 0 {
		founderCoins := sdk.NewCoins(sdk.NewInt64Coin("ibtc", founderReward))
		founderAddr := am.keeper.GetFounderAddress(ctx)
		if err := am.keeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, founderAddr, founderCoins); err != nil {
			ctx.EventManager().EmitEvent(sdk.NewEvent(types.ModuleName,
				sdk.NewAttribute("action", "send_founder_failed"),
				sdk.NewAttribute("error", err.Error()),
			))
		}
	}

	// Send miner share to block proposer
	if minerReward > 0 {
		proposer := ctx.BlockHeader().ProposerAddress
		if len(proposer) > 0 {
			proposerAcc := sdk.AccAddress(proposer)
			minerCoins := sdk.NewCoins(sdk.NewInt64Coin("ibtc", minerReward))
			if err := am.keeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, proposerAcc, minerCoins); err != nil {
				ctx.EventManager().EmitEvent(sdk.NewEvent(types.ModuleName,
					sdk.NewAttribute("action", "send_miner_failed"),
					sdk.NewAttribute("error", err.Error()),
				))
			}
		}
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.ModuleName,
			sdk.NewAttribute("action", "block_reward"),
			sdk.NewAttribute("founder", fmt.Sprintf("%d", founderReward)),
			sdk.NewAttribute("miner", fmt.Sprintf("%d", minerReward)),
		),
	)
}

func (am AppModule) EndBlock(_ sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}
