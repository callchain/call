package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "call/testutil/keeper"
	"call/x/call/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.CallKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
