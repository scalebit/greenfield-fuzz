package app_test

import (
	"testing"

	dbm "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/bnb-chain/greenfield/sdk/client/test"
	"github.com/bnb-chain/greenfield/testutil"
)

func TestExportAppStateAndValidators(t *testing.T) {
	logger := log.NewNopLogger()
	db := dbm.NewMemDB()
	app, _, _ := testutil.NewTestApp(logger, db, nil, true, test.TEST_CHAIN_ID)

	_, err := app.ExportAppStateAndValidators(false, nil, []string{banktypes.ModuleName})
	if err != nil {
		t.Fatalf("error exporting state: %s", err)
	}
}

// func FuzzExportAppStateAndValidators(f *testing.F) {
// 	logger := log.NewNopLogger()
// 	db := dbm.NewMemDB()
// 	app, _, _ := testutil.NewTestApp(logger, db, nil, true, test.TEST_CHAIN_ID)
// 	testcases := []string{""}
// 	for _, tc := range testcases {
// 		f.Add(tc) // Use f.Add to provide a seed corpus
// 	}

// 	f.Fuzz(func(t *testing.T, a string) {
// 		var data []string
// 		fuzz.NewFromGoFuzz([]byte(a)).NilChance(0.1).Fuzz(&data)

// 		_, err := app.ExportAppStateAndValidators(false, data, []string{banktypes.ModuleName})
// 		if err != nil {
// 			t.Fatalf("error exporting state: %s", err)
// 		}
// 	})

// }
