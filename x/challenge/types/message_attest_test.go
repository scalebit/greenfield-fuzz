package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"

	"github.com/bnb-chain/greenfield/testutil/sample"
)

func TestMsgAttest_ValidateBasic(t *testing.T) {
	var sig [96]byte
	tests := []struct {
		name string
		msg  MsgAttest
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgAttest{
				Submitter: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid vote result",
			msg: MsgAttest{
				Submitter:         sample.RandAccAddressHex(),
				SpOperatorAddress: sample.RandAccAddressHex(),
				VoteResult:        100,
			},
			err: ErrInvalidVoteResult,
		}, {
			name: "invalid vote result",
			msg: MsgAttest{
				Submitter:         sample.RandAccAddressHex(),
				SpOperatorAddress: sample.RandAccAddressHex(),
				VoteResult:        CHALLENGE_SUCCEED,
				VoteValidatorSet:  make([]uint64, 0),
			},
			err: ErrInvalidVoteValidatorSet,
		}, {
			name: "invalid vote aggregated signature",
			msg: MsgAttest{
				Submitter:         sample.RandAccAddressHex(),
				SpOperatorAddress: sample.RandAccAddressHex(),
				VoteResult:        CHALLENGE_SUCCEED,
				VoteValidatorSet:  []uint64{1},
				VoteAggSignature:  []byte{1, 2, 3},
			},
			err: ErrInvalidVoteAggSignature,
		}, {
			name: "valid message",
			msg: MsgAttest{
				Submitter:         sample.RandAccAddressHex(),
				SpOperatorAddress: sample.RandAccAddressHex(),
				VoteResult:        CHALLENGE_SUCCEED,
				VoteValidatorSet:  []uint64{1},
				VoteAggSignature:  sig[:],
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func FuzzMsgAttest_ValidateBasic(f *testing.F) {
	f.Add("")
	var sig [96]byte
	tt := struct {
		name string
		msg  MsgAttest
		err  error
	}{
		name: "valid message",
		msg: MsgAttest{
			Submitter:         sample.RandAccAddressHex(),
			SpOperatorAddress: sample.RandAccAddressHex(),
			VoteResult:        CHALLENGE_SUCCEED,
			VoteValidatorSet:  []uint64{1},
			VoteAggSignature:  sig[:],
		}}
	f.Fuzz(func(t *testing.T, a string) {
		t.Run(tt.name, func(t *testing.T) {
			vv := []uint64{}
			fuzz.NewFromGoFuzz([]byte(a)).NumElements(5, 5).Fuzz(&vv)
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	})
}
