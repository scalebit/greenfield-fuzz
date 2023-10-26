package types

import (
	"math/rand"
	"testing"
	"time"

	"cosmossdk.io/math"
)

func TestParsePolicyIdFromQueueKey(t *testing.T) {
	policyIds := []math.Uint{math.NewUint(rand.Uint64()), math.NewUint(rand.Uint64()), math.NewUint(rand.Uint64())}

	expiration := time.Now()
	for _, policyId := range policyIds {
		key := PolicyPrefixQueue(&expiration, policyId.Bytes())
		recoverId := ParsePolicyIdFromQueueKey(key)
		if !recoverId.Equal(policyId) {
			t.Errorf("ParseIdFromQueueKey failed to recover policy id: %s", policyId.String())
		}
	}
}

func FuzzParsePolicyIdFromQueueKey(f *testing.F) {
	f.Add(1)
	f.Fuzz(func(t *testing.T, a int) {
		if a < 0 {
			a = -a
		}

		policyIds := []math.Uint{}
		for i := 0; i < a; i++ {
			policyIds = append(policyIds, math.NewUint(rand.Uint64()))
		}

		expiration := time.Now()
		for _, policyId := range policyIds {
			key := PolicyPrefixQueue(&expiration, policyId.Bytes())
			recoverId := ParsePolicyIdFromQueueKey(key)
			if !recoverId.Equal(policyId) {
				t.Errorf("ParseIdFromQueueKey failed to recover policy id: %s", policyId.String())
			}
		}
	})
}
