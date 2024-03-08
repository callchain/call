package call_test

import (
	"testing"

	keepertest "call/testutil/keeper"
	"call/testutil/nullify"
	"call/x/call/module"
	"call/x/call/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CallKeeper(t)
	call.InitGenesis(ctx, k, genesisState)
	got := call.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
