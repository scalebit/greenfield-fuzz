package keeper_test

import (
	"testing"

	"github.com/bnb-chain/greenfield/testutil/sample"
	"github.com/bnb-chain/greenfield/x/payment/types"
)

func (s *TestSuite) TestCreatePaymentAccount() {
	creator := sample.RandAccAddress()

	// create first one
	msg := types.NewMsgCreatePaymentAccount(creator.String())
	_, err := s.msgServer.CreatePaymentAccount(s.ctx, msg)
	s.Require().NoError(err)

	record, _ := s.paymentKeeper.GetPaymentAccountCount(s.ctx, creator)
	s.Require().True(record.Count == 1)

	// create another one
	msg = types.NewMsgCreatePaymentAccount(creator.String())
	_, err = s.msgServer.CreatePaymentAccount(s.ctx, msg)
	s.Require().NoError(err)

	record, _ = s.paymentKeeper.GetPaymentAccountCount(s.ctx, creator)
	s.Require().True(record.Count == 2)

	// limit the number of payment account
	params := s.paymentKeeper.GetParams(s.ctx)
	params.PaymentAccountCountLimit = 2
	_ = s.paymentKeeper.SetParams(s.ctx, params)

	msg = types.NewMsgCreatePaymentAccount(creator.String())
	_, err = s.msgServer.CreatePaymentAccount(s.ctx, msg)
	s.Require().Error(err)
}

func FuzzCreatePaymentAccount(f *testing.F) {
	f.Add(uint64(2))
	f.Fuzz(func(t *testing.T, a uint64) {
		s := &TestSuite{}
		s.SetT(t)
		s.SetupTest()

		creator := sample.RandAccAddress()

		// create first one
		msg := types.NewMsgCreatePaymentAccount(creator.String())
		_, err := s.msgServer.CreatePaymentAccount(s.ctx, msg)
		s.Require().NoError(err)

		record, _ := s.paymentKeeper.GetPaymentAccountCount(s.ctx, creator)
		s.Require().True(record.Count == uint64(1))

		// limit the number of payment account
		params := s.paymentKeeper.GetParams(s.ctx)
		params.PaymentAccountCountLimit = a
		_ = s.paymentKeeper.SetParams(s.ctx, params)

		for i := uint64(2); i <= a; i++ {
			// create another one
			msg = types.NewMsgCreatePaymentAccount(creator.String())
			_, err = s.msgServer.CreatePaymentAccount(s.ctx, msg)

			s.Require().NoError(err)

			record, _ = s.paymentKeeper.GetPaymentAccountCount(s.ctx, creator)
			s.Require().True(record.Count == i)
		}

		msg = types.NewMsgCreatePaymentAccount(creator.String())
		_, err = s.msgServer.CreatePaymentAccount(s.ctx, msg)
		t.Log(a, err)
		if a == 0 {
			return
		}
		s.Require().Error(err)
	})

}
