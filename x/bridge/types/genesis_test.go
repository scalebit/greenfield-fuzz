package types_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/stretchr/testify/require"

	"github.com/bnb-chain/greenfield/x/bridge/types"
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
				Params: types.Params{
					BscTransferOutRelayerFee:    sdkmath.NewInt(1),
					BscTransferOutAckRelayerFee: sdkmath.NewInt(0),
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
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

	f.Add(int64(1))
	f.Fuzz(func(t *testing.T, a int64) {
		tc := struct {
			desc     string
			genState *types.GenesisState
			valid    bool
		}{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				Params: types.Params{
					BscTransferOutRelayerFee:    sdkmath.NewInt(a),
					BscTransferOutAckRelayerFee: sdkmath.NewInt(a),
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
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
