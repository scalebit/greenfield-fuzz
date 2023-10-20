package keeper_test

import (
	"testing"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"

	"github.com/bnb-chain/greenfield/x/challenge/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := makeKeeper(t)
	params := types.DefaultParams()

	err := k.SetParams(ctx, params)
	require.NoError(t, err)
	require.EqualValues(t, params, k.GetParams(ctx))

	params.AttestationKeptCount = 100
	err = k.SetParams(ctx, params)
	require.NoError(t, err)
	require.EqualValues(t, params, k.GetParams(ctx))
}

func FuzzGetParams(f *testing.F) {
	f.Add("")
	f.Fuzz(func(t *testing.T, a string) {
		k, ctx := makeKeeper(t)
		params := types.DefaultParams()
		fuzz.NewFromGoFuzz([]byte(a)).Fuzz(&params)
		err := k.SetParams(ctx, params)
		// require.NoError(t, err)
		if err != nil {
			return
		}
		require.EqualValues(t, params, k.GetParams(ctx))

		params.AttestationKeptCount = 100
		fuzz.NewFromGoFuzz([]byte(a)).Fuzz(&params.AttestationKeptCount)
		err = k.SetParams(ctx, params)
		// require.NoError(t, err)
		if err != nil {
			return
		}
		require.EqualValues(t, params, k.GetParams(ctx))
	})
}
