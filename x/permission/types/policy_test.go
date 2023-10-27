package types_test

import (
	"math/rand"
	"regexp"
	"testing"
	"time"

	"cosmossdk.io/math"
	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"

	"github.com/bnb-chain/greenfield/testutil/sample"
	"github.com/bnb-chain/greenfield/testutil/storage"
	types2 "github.com/bnb-chain/greenfield/types"
	"github.com/bnb-chain/greenfield/types/common"
	"github.com/bnb-chain/greenfield/types/resource"
	"github.com/bnb-chain/greenfield/types/s3util"
	"github.com/bnb-chain/greenfield/x/permission/types"
)

func TestPolicy_BucketBasic(t *testing.T) {
	tests := []struct {
		name          string
		policyAction  types.ActionType
		policyEffect  types.Effect
		operateAction types.ActionType
		expectEffect  types.Effect
	}{
		{
			name:          "basic_update_bucket_info",
			policyAction:  types.ACTION_UPDATE_BUCKET_INFO,
			policyEffect:  types.EFFECT_ALLOW,
			operateAction: types.ACTION_UPDATE_BUCKET_INFO,
			expectEffect:  types.EFFECT_ALLOW,
		},
		{
			name:          "basic_delete_bucket",
			policyAction:  types.ACTION_DELETE_BUCKET,
			policyEffect:  types.EFFECT_ALLOW,
			operateAction: types.ACTION_DELETE_BUCKET,
			expectEffect:  types.EFFECT_ALLOW,
		},
		{
			name:          "basic_delete_bucket_deny",
			policyAction:  types.ACTION_DELETE_BUCKET,
			policyEffect:  types.EFFECT_DENY,
			operateAction: types.ACTION_DELETE_BUCKET,
			expectEffect:  types.EFFECT_DENY,
		},
		{
			name:          "basic_delete_bucket_pass",
			policyAction:  types.ACTION_UPDATE_BUCKET_INFO,
			policyEffect:  types.EFFECT_ALLOW,
			operateAction: types.ACTION_DELETE_BUCKET,
			expectEffect:  types.EFFECT_UNSPECIFIED,
		},
		{
			name:          "basic_create_object",
			policyAction:  types.ACTION_CREATE_OBJECT,
			policyEffect:  types.EFFECT_ALLOW,
			operateAction: types.ACTION_CREATE_OBJECT,
			expectEffect:  types.EFFECT_ALLOW,
		},
		{
			name:          "basic_create_object_deny",
			policyAction:  types.ACTION_CREATE_OBJECT,
			policyEffect:  types.EFFECT_DENY,
			operateAction: types.ACTION_CREATE_OBJECT,
			expectEffect:  types.EFFECT_DENY,
		},
		{
			name:          "basic_create_object_pass",
			policyAction:  types.ACTION_COPY_OBJECT,
			policyEffect:  types.EFFECT_ALLOW,
			operateAction: types.ACTION_CREATE_OBJECT,
			expectEffect:  types.EFFECT_UNSPECIFIED,
		},
		{
			name:          "basic_delete_object",
			policyAction:  types.ACTION_DELETE_OBJECT,
			policyEffect:  types.EFFECT_ALLOW,
			operateAction: types.ACTION_DELETE_OBJECT,
			expectEffect:  types.EFFECT_ALLOW,
		},
		{
			name:          "basic_delete_object_deny",
			policyAction:  types.ACTION_DELETE_OBJECT,
			policyEffect:  types.EFFECT_DENY,
			operateAction: types.ACTION_DELETE_OBJECT,
			expectEffect:  types.EFFECT_DENY,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := sample.RandAccAddress()
			policy := types.Policy{
				Principal:    types.NewPrincipalWithAccount(user),
				ResourceType: resource.RESOURCE_TYPE_BUCKET,
				ResourceId:   math.OneUint(),
				Statements: []*types.Statement{
					{
						Effect:  tt.policyEffect,
						Actions: []types.ActionType{tt.policyAction},
					},
				},
			}
			effect, _ := policy.Eval(tt.operateAction, time.Now(), nil)
			require.Equal(t, effect, tt.expectEffect)
		})
	}
}

func FuzzPolicy_BucketBasic(f *testing.F) {
	f.Add(1)
	f.Fuzz(func(t *testing.T, a int) {
		tests := []struct {
			name          string
			policyAction  types.ActionType
			policyEffect  types.Effect
			operateAction types.ActionType
			expectEffect  types.Effect
		}{
			{
				name:          "basic_update_bucket_info",
				policyAction:  types.ACTION_UPDATE_BUCKET_INFO,
				policyEffect:  types.EFFECT_ALLOW,
				operateAction: types.ACTION_UPDATE_BUCKET_INFO,
				expectEffect:  types.EFFECT_ALLOW,
			},
			{
				name:          "basic_delete_bucket",
				policyAction:  types.ACTION_DELETE_BUCKET,
				policyEffect:  types.EFFECT_ALLOW,
				operateAction: types.ACTION_DELETE_BUCKET,
				expectEffect:  types.EFFECT_ALLOW,
			},
			{
				name:          "basic_delete_bucket_deny",
				policyAction:  types.ACTION_DELETE_BUCKET,
				policyEffect:  types.EFFECT_DENY,
				operateAction: types.ACTION_DELETE_BUCKET,
				expectEffect:  types.EFFECT_DENY,
			},
			{
				name:          "basic_delete_bucket_pass",
				policyAction:  types.ACTION_UPDATE_BUCKET_INFO,
				policyEffect:  types.EFFECT_ALLOW,
				operateAction: types.ACTION_DELETE_BUCKET,
				expectEffect:  types.EFFECT_UNSPECIFIED,
			},
			{
				name:          "basic_create_object",
				policyAction:  types.ACTION_CREATE_OBJECT,
				policyEffect:  types.EFFECT_ALLOW,
				operateAction: types.ACTION_CREATE_OBJECT,
				expectEffect:  types.EFFECT_ALLOW,
			},
			{
				name:          "basic_create_object_deny",
				policyAction:  types.ACTION_CREATE_OBJECT,
				policyEffect:  types.EFFECT_DENY,
				operateAction: types.ACTION_CREATE_OBJECT,
				expectEffect:  types.EFFECT_DENY,
			},
			{
				name:          "basic_create_object_pass",
				policyAction:  types.ACTION_COPY_OBJECT,
				policyEffect:  types.EFFECT_ALLOW,
				operateAction: types.ACTION_CREATE_OBJECT,
				expectEffect:  types.EFFECT_UNSPECIFIED,
			},
			{
				name:          "basic_delete_object",
				policyAction:  types.ACTION_DELETE_OBJECT,
				policyEffect:  types.EFFECT_ALLOW,
				operateAction: types.ACTION_DELETE_OBJECT,
				expectEffect:  types.EFFECT_ALLOW,
			},
			{
				name:          "basic_delete_object_deny",
				policyAction:  types.ACTION_DELETE_OBJECT,
				policyEffect:  types.EFFECT_DENY,
				operateAction: types.ACTION_DELETE_OBJECT,
				expectEffect:  types.EFFECT_DENY,
			},
		}

		tt := tests[0]
		rand.Seed(int64(a))
		r1 := rand.Intn(8)
		r2 := rand.Intn(8)
		r3 := rand.Intn(8)
		r4 := rand.Intn(8)
		t.Log(r1, r2, r3, r4)
		tt.policyAction = tests[r1].policyAction
		tt.policyEffect = tests[r2].policyEffect
		tt.operateAction = tests[r3].operateAction
		tt.expectEffect = tests[r4].expectEffect

		t.Run(tt.name, func(t *testing.T) {
			user := sample.RandAccAddress()
			policy := types.Policy{
				Principal:    types.NewPrincipalWithAccount(user),
				ResourceType: resource.RESOURCE_TYPE_BUCKET,
				ResourceId:   math.OneUint(),
				Statements: []*types.Statement{
					{
						Effect:  tt.policyEffect,
						Actions: []types.ActionType{tt.policyAction},
					},
				},
			}
			_, _ = policy.Eval(tt.operateAction, time.Now(), nil)
			// require.Equal(t, effect, tt.expectEffect)
		})

	})
}

func TestPolicy_BucketExpirationBasic(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name                    string
		policyAction            types.ActionType
		policyEffect            types.Effect
		policyExpirationTime    *time.Time
		statementExpirationTime *time.Time
		operateAction           types.ActionType
		operateTime             time.Time
		expectEffect            types.Effect
	}{
		{
			name:                 "policy_expired",
			policyAction:         types.ACTION_UPDATE_BUCKET_INFO,
			policyEffect:         types.EFFECT_ALLOW,
			policyExpirationTime: &now,
			operateAction:        types.ACTION_UPDATE_BUCKET_INFO,
			expectEffect:         types.EFFECT_UNSPECIFIED,
			operateTime:          time.Now().Add(time.Duration(1 * time.Second)),
		},
		{
			name:                 "policy_not_expired",
			policyAction:         types.ACTION_UPDATE_BUCKET_INFO,
			policyEffect:         types.EFFECT_ALLOW,
			policyExpirationTime: &now,
			operateAction:        types.ACTION_UPDATE_BUCKET_INFO,
			expectEffect:         types.EFFECT_ALLOW,
			operateTime:          time.Now().Add(-time.Duration(1 * time.Second)),
		},
		{
			name:                    "statement_expired",
			policyAction:            types.ACTION_UPDATE_BUCKET_INFO,
			policyEffect:            types.EFFECT_ALLOW,
			statementExpirationTime: &now,
			operateAction:           types.ACTION_UPDATE_BUCKET_INFO,
			expectEffect:            types.EFFECT_UNSPECIFIED,
			operateTime:             time.Now().Add(time.Duration(1 * time.Second)),
		},
		{
			name:                 "statement_not_expired",
			policyAction:         types.ACTION_UPDATE_BUCKET_INFO,
			policyEffect:         types.EFFECT_ALLOW,
			policyExpirationTime: &now,
			operateAction:        types.ACTION_UPDATE_BUCKET_INFO,
			expectEffect:         types.EFFECT_ALLOW,
			operateTime:          time.Now().Add(-time.Duration(1 * time.Second)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := sample.RandAccAddress()
			policy := types.Policy{
				Principal:      types.NewPrincipalWithAccount(user),
				ResourceType:   resource.RESOURCE_TYPE_BUCKET,
				ResourceId:     math.OneUint(),
				ExpirationTime: tt.policyExpirationTime,
				Statements: []*types.Statement{
					{
						Effect:         tt.policyEffect,
						Actions:        []types.ActionType{tt.policyAction},
						ExpirationTime: tt.statementExpirationTime,
					},
				},
			}
			effect, _ := policy.Eval(tt.operateAction, tt.operateTime, nil)
			require.Equal(t, effect, tt.expectEffect)
		})
	}
}

func FuzzPolicy_BucketExpirationBasic(f *testing.F) {
	f.Add(1)
	f.Fuzz(func(t *testing.T, a int) {
		now := time.Now()
		tests := []struct {
			name                    string
			policyAction            types.ActionType
			policyEffect            types.Effect
			policyExpirationTime    *time.Time
			statementExpirationTime *time.Time
			operateAction           types.ActionType
			operateTime             time.Time
			expectEffect            types.Effect
		}{
			{
				name:                 "policy_expired",
				policyAction:         types.ACTION_UPDATE_BUCKET_INFO,
				policyEffect:         types.EFFECT_ALLOW,
				policyExpirationTime: &now,
				operateAction:        types.ACTION_UPDATE_BUCKET_INFO,
				expectEffect:         types.EFFECT_UNSPECIFIED,
				operateTime:          time.Now().Add(time.Duration(1 * time.Second)),
			},
			{
				name:                 "policy_not_expired",
				policyAction:         types.ACTION_UPDATE_BUCKET_INFO,
				policyEffect:         types.EFFECT_ALLOW,
				policyExpirationTime: &now,
				operateAction:        types.ACTION_UPDATE_BUCKET_INFO,
				expectEffect:         types.EFFECT_ALLOW,
				operateTime:          time.Now().Add(-time.Duration(1 * time.Second)),
			},
			{
				name:                    "statement_expired",
				policyAction:            types.ACTION_UPDATE_BUCKET_INFO,
				policyEffect:            types.EFFECT_ALLOW,
				statementExpirationTime: &now,
				operateAction:           types.ACTION_UPDATE_BUCKET_INFO,
				expectEffect:            types.EFFECT_UNSPECIFIED,
				operateTime:             time.Now().Add(time.Duration(1 * time.Second)),
			},
			{
				name:                 "statement_not_expired",
				policyAction:         types.ACTION_UPDATE_BUCKET_INFO,
				policyEffect:         types.EFFECT_ALLOW,
				policyExpirationTime: &now,
				operateAction:        types.ACTION_UPDATE_BUCKET_INFO,
				expectEffect:         types.EFFECT_ALLOW,
				operateTime:          time.Now().Add(-time.Duration(1 * time.Second)),
			},
		}

		tt := tests[0]
		rand.Seed(int64(a))
		r1 := rand.Intn(4)
		r2 := rand.Intn(4)
		r3 := rand.Intn(4)
		r4 := rand.Intn(4)
		r5 := rand.Intn(4)
		r6 := rand.Intn(4)
		r7 := rand.Intn(4)
		tt.name = tests[r1].name
		tt.policyAction = tests[r2].policyAction
		tt.policyEffect = tests[r3].policyEffect
		tt.policyExpirationTime = tests[r4].policyExpirationTime
		tt.operateAction = tests[r5].operateAction
		tt.expectEffect = tests[r6].expectEffect
		tt.operateTime = tests[r7].operateTime

		t.Run(tt.name, func(t *testing.T) {
			user := sample.RandAccAddress()
			policy := types.Policy{
				Principal:      types.NewPrincipalWithAccount(user),
				ResourceType:   resource.RESOURCE_TYPE_BUCKET,
				ResourceId:     math.OneUint(),
				ExpirationTime: tt.policyExpirationTime,
				Statements: []*types.Statement{
					{
						Effect:         tt.policyEffect,
						Actions:        []types.ActionType{tt.policyAction},
						ExpirationTime: tt.statementExpirationTime,
					},
				},
			}
			_, _ = policy.Eval(tt.operateAction, tt.operateTime, nil)
			// require.Equal(t, effect, tt.expectEffect)
		})

	})
}

func TestPolicy_CreateObjectLimitSize(t *testing.T) {
	tests := []struct {
		name         string
		limitSize    uint64
		wantedSize   uint64
		expectEffect types.Effect
	}{
		{
			name:         "limit_size_not_exceed",
			limitSize:    2 * 1024,
			wantedSize:   1 * 1024,
			expectEffect: types.EFFECT_ALLOW,
		},
		{
			name:         "limit_size_exceed",
			limitSize:    2 * 1024,
			wantedSize:   3 * 1024,
			expectEffect: types.EFFECT_DENY,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := sample.RandAccAddress()
			policy := types.Policy{
				Principal:    types.NewPrincipalWithAccount(user),
				ResourceType: resource.RESOURCE_TYPE_BUCKET,
				ResourceId:   math.OneUint(),
				Statements: []*types.Statement{
					{
						Effect:    types.EFFECT_ALLOW,
						Actions:   []types.ActionType{types.ACTION_CREATE_OBJECT},
						LimitSize: &common.UInt64Value{Value: tt.limitSize},
					},
				},
			}
			effect, p := policy.Eval(types.ACTION_CREATE_OBJECT, time.Now(), &types.VerifyOptions{WantedSize: &tt.wantedSize})
			if effect == types.EFFECT_ALLOW && tt.limitSize > tt.wantedSize {
				require.Equal(t, p.Statements[0].LimitSize.GetValue(), tt.limitSize-tt.wantedSize)
			}
			require.Equal(t, effect, tt.expectEffect)
		})
	}
}

func FuzzPolicy_CreateObjectLimitSize(f *testing.F) {
	f.Add(uint64(1))
	f.Fuzz(func(t *testing.T, a uint64) {
		rand.Seed(int64(a))
		b := rand.Uint64()
		var name string
		fuzz.New().Fuzz(&name)
		tt := struct {
			name         string
			limitSize    uint64
			wantedSize   uint64
			expectEffect types.Effect
		}{
			name:         "limit_size_not_exceed",
			limitSize:    b * 1024,
			wantedSize:   a * 1024,
			expectEffect: types.EFFECT_ALLOW,
		}

		t.Run(tt.name, func(t *testing.T) {
			user := sample.RandAccAddress()
			policy := types.Policy{
				Principal:    types.NewPrincipalWithAccount(user),
				ResourceType: resource.RESOURCE_TYPE_BUCKET,
				ResourceId:   math.OneUint(),
				Statements: []*types.Statement{
					{
						Effect:    types.EFFECT_ALLOW,
						Actions:   []types.ActionType{types.ACTION_CREATE_OBJECT},
						LimitSize: &common.UInt64Value{Value: tt.limitSize},
					},
				},
			}
			effect, p := policy.Eval(types.ACTION_CREATE_OBJECT, time.Now(), &types.VerifyOptions{WantedSize: &tt.wantedSize})
			if effect == types.EFFECT_ALLOW && tt.limitSize > tt.wantedSize {
				require.Equal(t, p.Statements[0].LimitSize.GetValue(), tt.limitSize-tt.wantedSize)
			}
			// require.Equal(t, effect, tt.expectEffect)
		})

	})
}

func TestPolicy_SubResource(t *testing.T) {
	bucketName := storage.GenRandomBucketName()
	tests := []struct {
		name            string
		policyAction    types.ActionType
		policyEffect    types.Effect
		policyResource  string
		operateAction   types.ActionType
		operateResource string
		expectEffect    types.Effect
	}{
		{
			name:            "policy_resource_matched_allow",
			policyAction:    types.ACTION_GET_OBJECT,
			policyEffect:    types.EFFECT_ALLOW,
			policyResource:  types2.NewObjectGRN(bucketName, "*").String(),
			operateAction:   types.ACTION_GET_OBJECT,
			operateResource: types2.NewObjectGRN(bucketName, "xxxx").String(),
			expectEffect:    types.EFFECT_ALLOW,
		},
		{
			name:            "policy_resource_matched_deny",
			policyAction:    types.ACTION_GET_OBJECT,
			policyEffect:    types.EFFECT_DENY,
			policyResource:  types2.NewObjectGRN(bucketName, "*").String(),
			operateAction:   types.ACTION_GET_OBJECT,
			operateResource: types2.NewObjectGRN(bucketName, "xxxx").String(),
			expectEffect:    types.EFFECT_DENY,
		},
		{
			name:            "policy_resource_not_matched",
			policyAction:    types.ACTION_GET_OBJECT,
			policyEffect:    types.EFFECT_ALLOW,
			policyResource:  types2.NewObjectGRN(bucketName, "xxx").String(),
			operateAction:   types.ACTION_GET_OBJECT,
			operateResource: types2.NewObjectGRN(bucketName, "1111").String(),
			expectEffect:    types.EFFECT_UNSPECIFIED,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := sample.RandAccAddress()
			policy := types.Policy{
				Principal:    types.NewPrincipalWithAccount(user),
				ResourceType: resource.RESOURCE_TYPE_BUCKET,
				ResourceId:   math.OneUint(),
				Statements: []*types.Statement{
					{
						Effect:    tt.policyEffect,
						Actions:   []types.ActionType{tt.policyAction},
						Resources: []string{tt.policyResource},
					},
				},
			}
			effect, _ := policy.Eval(tt.operateAction, time.Now(), &types.VerifyOptions{Resource: tt.operateResource})
			require.Equal(t, effect, tt.expectEffect)
		})
	}
}

// func FuzzPolicy_SubResource(f *testing.F) {
// 	f.Add("\"(\"") //crash "("
// 	f.Fuzz(func(t *testing.T, a string) {
// 		bucketName := storage.GenRandomBucketName()
// 		tests := []struct {
// 			name            string
// 			policyAction    types.ActionType
// 			policyEffect    types.Effect
// 			policyResource  string
// 			operateAction   types.ActionType
// 			operateResource string
// 			expectEffect    types.Effect
// 		}{
// 			{
// 				name:            "policy_resource_matched_allow",
// 				policyAction:    types.ACTION_GET_OBJECT,
// 				policyEffect:    types.EFFECT_ALLOW,
// 				policyResource:  types2.NewObjectGRN(bucketName, a).String(),
// 				operateAction:   types.ACTION_GET_OBJECT,
// 				operateResource: types2.NewObjectGRN(bucketName, "xxxx").String(),
// 				expectEffect:    types.EFFECT_ALLOW,
// 			},
// 			{
// 				name:            "policy_resource_matched_deny",
// 				policyAction:    types.ACTION_GET_OBJECT,
// 				policyEffect:    types.EFFECT_DENY,
// 				policyResource:  types2.NewObjectGRN(bucketName, a).String(),
// 				operateAction:   types.ACTION_GET_OBJECT,
// 				operateResource: types2.NewObjectGRN(bucketName, "xxxx").String(),
// 				expectEffect:    types.EFFECT_DENY,
// 			},
// 			{
// 				name:            "policy_resource_not_matched",
// 				policyAction:    types.ACTION_GET_OBJECT,
// 				policyEffect:    types.EFFECT_ALLOW,
// 				policyResource:  types2.NewObjectGRN(bucketName, a).String(),
// 				operateAction:   types.ACTION_GET_OBJECT,
// 				operateResource: types2.NewObjectGRN(bucketName, "1111").String(),
// 				expectEffect:    types.EFFECT_UNSPECIFIED,
// 			},
// 		}

// 		for _, tt := range tests {
// 			t.Run(tt.name, func(t *testing.T) {
// 				user := sample.RandAccAddress()
// 				policy := types.Policy{
// 					Principal:    types.NewPrincipalWithAccount(user),
// 					ResourceType: resource.RESOURCE_TYPE_BUCKET,
// 					ResourceId:   math.OneUint(),
// 					Statements: []*types.Statement{
// 						{
// 							Effect:    tt.policyEffect,
// 							Actions:   []types.ActionType{tt.policyAction},
// 							Resources: []string{tt.policyResource},
// 						},
// 					},
// 				}

// 				t.Log(tt.operateResource)
// 				t.Log(tt.operateAction)
// 				_, _ = policy.Eval(tt.operateAction, time.Now(), &types.VerifyOptions{Resource: tt.operateResource})
// 				t.Log(123)
// 				// require.Equal(t, effect, tt.expectEffect)
// 			})
// 		}
// 	})
// }

func FuzzCheckValidBucketName(f *testing.F) {
	f.Add("a\x28-123b")
	// f.Add("ab123-a.a-1")
	f.Fuzz(func(t *testing.T, a string) {
		err := s3util.CheckValidBucketName(a)
		t.Log(err)
		if err != nil {
			return
		}
		_, err = regexp.Compile(a)
		require.NoError(t, err)

	})

}
