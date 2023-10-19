package types_test

import (
	"testing"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"

	"github.com/bnb-chain/greenfield/x/challenge/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),

				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func FuzzGenesisState_Validate(f *testing.F) {
	f.Add("2")
	tc := struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		desc:     "default is valid",
		genState: types.DefaultGenesis(),
		valid:    true,
	}
	f.Fuzz(func(t *testing.T, a string) {
		fuzz.NewFromGoFuzz([]byte(a)).Fuzz(&tc.genState)
		if tc.genState == nil {
			return
		}
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				// require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	})

}
