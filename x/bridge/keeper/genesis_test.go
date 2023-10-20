package keeper_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"

	"github.com/bnb-chain/greenfield/x/bridge/keeper"
	"github.com/bnb-chain/greenfield/x/bridge/types"
	fuzz "github.com/google/gofuzz"
)

func (s *TestSuite) TestExportGenesis() {
	ctx := s.ctx

	s.Require().NoError(s.bridgeKeeper.SetParams(ctx, types.DefaultParams()))
	exportGenesis := keeper.ExportGenesis(ctx, *s.bridgeKeeper)

	s.Require().Equal(types.DefaultParams().BscTransferOutRelayerFee, exportGenesis.Params.BscTransferOutRelayerFee)
	s.Require().Equal(types.DefaultParams().BscTransferOutAckRelayerFee, exportGenesis.Params.BscTransferOutAckRelayerFee)
}

func FuzzExportGenesis(f *testing.F) {
	s := &TestSuite{}
	f.Add("")
	f.Fuzz(func(t *testing.T, a string) {
		s.SetT(t)
		s.SetupTest()
		ctx := s.ctx
		params := types.DefaultParams()
		fuzz.NewFromGoFuzz([]byte(a)).Fuzz(&params)
		s.Require().NoError(s.bridgeKeeper.SetParams(ctx, params))
		exportGenesis := keeper.ExportGenesis(ctx, *s.bridgeKeeper)
		s.Require().Equal(params.BscTransferOutRelayerFee, exportGenesis.Params.BscTransferOutRelayerFee)
		s.Require().Equal(params.BscTransferOutAckRelayerFee, exportGenesis.Params.BscTransferOutAckRelayerFee)
	})
}

func (s *TestSuite) TestInitGenesis() {
	g := types.DefaultGenesis()
	k := s.bridgeKeeper
	keeper.InitGenesis(s.ctx, *k, *g)

	// Check that the genesis state was set correctly.
	params := k.GetParams(s.ctx)
	s.Require().Equal(sdkmath.NewInt(780000000000000), params.BscTransferOutRelayerFee)
	s.Require().Equal(sdkmath.NewInt(0), params.BscTransferOutAckRelayerFee)
}

// func FuzzInitGenesis(f *testing.F) {
// 	s := &TestSuite{}
// 	f.Add("")
// 	f.Fuzz(func(t *testing.T, a string) {
// 		s.SetT(t)
// 		s.SetupTest()

// 		g := types.DefaultGenesis()
// 		k := s.bridgeKeeper
// 		keeper.InitGenesis(s.ctx, *k, *g)

// 		// Check that the genesis state was set correctly.
// 		params := k.GetParams(s.ctx)
// 		s.Require().Equal(sdkmath.NewInt(780000000000000), params.BscTransferOutRelayerFee)
// 		s.Require().Equal(sdkmath.NewInt(0), params.BscTransferOutAckRelayerFee)
// 	})
// }
