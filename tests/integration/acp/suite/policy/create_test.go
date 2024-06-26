package policy

import (
	"testing"

	test "github.com/sourcenetwork/sourcehub/tests/integration/acp"
	"github.com/sourcenetwork/sourcehub/x/acp/policy"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

func TestCreatePolicy_ValidPolicyIsCreated(t *testing.T) {
	ctx := test.NewTestCtx(t)

	policyStr := `
name: policy
description: ok
resources:
  file:
    relations:
      owner:
        doc: owner owns
        types:
          - actor-resource
      reader:
      admin:
        manages:
          - reader
    permissions:
      own:
        expr: owner
        doc: own doc
      read:
        expr: owner + reader
actor:
  name: actor-resource
  doc: my actor
`
	want := &types.Policy{
		Id:           "d431eabc84aa4d2e782b31b37382aeb73fe435c10c35769756ee8302d8371b1c",
		Name:         "policy",
		Description:  "ok",
		CreationTime: test.TimeToProto(ctx.Timestamp),
		Creator:      ctx.TxSigner.SourceHubAddr,
		Resources: []*types.Resource{
			&types.Resource{
				Name: "file",
				Relations: []*types.Relation{
					&types.Relation{
						Name: "admin",
						Manages: []string{
							"reader",
						},
						VrTypes: []*types.Restriction{},
					},
					&types.Relation{
						Name: "owner",
						Doc:  "owner owns",
						VrTypes: []*types.Restriction{
							&types.Restriction{
								ResourceName: "actor-resource",
								RelationName: "",
							},
						},
					},
					&types.Relation{
						Name: "reader",
					},
				},
				Permissions: []*types.Permission{
					&types.Permission{
						Name:       "own",
						Expression: "owner",
						Doc:        "own doc",
					},
					&types.Permission{
						Name:       "read",
						Expression: "owner + reader",
					},
					&types.Permission{
						Name:       "_can_manage_admin",
						Expression: "owner",
						Doc:        "permission controls actors which are allowed to create relationships for the admin relation (permission was auto-generated by SourceHub).",
					},
					&types.Permission{
						Name:       "_can_manage_owner",
						Expression: "owner",
						Doc:        "permission controls actors which are allowed to create relationships for the owner relation (permission was auto-generated by SourceHub).",
					},
					&types.Permission{
						Name:       "_can_manage_reader",
						Expression: "(admin + owner)",
						Doc:        "permission controls actors which are allowed to create relationships for the reader relation (permission was auto-generated by SourceHub).",
					},
				},
			},
		},
		ActorResource: &types.ActorResource{
			Name: "actor-resource",
			Doc:  "my actor",
		},
	}

	action := test.CreatePolicyAction{
		Policy:   policyStr,
		Expected: want,
		Creator:  ctx.TxSigner,
	}
	action.Run(ctx)

	event := &types.EventPolicyCreated{
		Creator:    ctx.TxSigner.SourceHubAddr,
		PolicyId:   "4419a8abb886c641bc794b9b3289bc2118ab177542129627b6b05d540de03e46",
		PolicyName: "policy",
	}
	_ = event
	//.AssertEventEmmited(t, ctx, event)
}

func TestCreatePolicy_PolicyResourcesRequiresOwnerRelation(t *testing.T) {
	ctx := test.NewTestCtx(t)

	action := test.CreatePolicyAction{
		Policy: `
name: policy
description: ok
resources:
  file:
    relations:
      reader:
    permissions:
  foo:
    relations:
      owner:
    permissions:
`,
		Creator:     ctx.TxSigner,
		ExpectedErr: policy.ErrResourceMissingOwnerRelation,
	}
	action.Run(ctx)
}

func TestCreatePolicy_ManagementReferencingUndefinedRelationReturnsError(t *testing.T) {
	ctx := test.NewTestCtx(t)

	action := test.CreatePolicyAction{
		Policy: `
name: policy
description: ok
resources:
  file:
    relations:
      owner:
      admin:
        manages:
          - deleter
    permissions:
`,
		Creator:     ctx.TxSigner,
		ExpectedErr: policy.ErrInvalidManagementRule,
	}
	action.Run(ctx)
}

/*
func TestCreatePolicy_UnamedPolicyCausesError(t *testing.T) {
	ctx := test.NewTestCtx(t)

	a := test.CreatePolicyAction{
		Policy: `
resources:
`,
		Creator:     ctx.TxSigner,
		ExpectedErr: policy.ErrInvalidPolicy,
	}
	a.Run(ctx)
}

func TestCreatePolicy_CreatingMultipleEqualPoliciesProduceDifferentIDs(t *testing.T) {
	ctx := test.NewTestCtx(t)
	ctx.SetPrincipal("creator")

	pol := `
name: test
description: A Valid Defra Policy Interface (DPI)

actor:
  name: actor

resources:
  users:
    permissions:
      read:
        expr: owner + reader
      write:
        expr: owner

    relations:
      owner:
        types:
          - actor
      reader:
        types:
          - actor
      admin:
        manages:
          - reader
        types:
          - actor
`

	req := types.CreatePolicyRequest{
		Policy:      pol,
		MarshalType: types.PolicyMarshalingType_SHORT_YAML,
	}
	resp1, err1 := ctx.Engine.CreatePolicy(ctx, &req)
	resp2, err2 := ctx.Engine.CreatePolicy(ctx, &req)

	want1 := "94eb195c0e459aa79e02a1986c7e731c5015721c18a373f2b2a0ed140a04b454"
	want2 := "7bcb558ef8dac6b744a11ea144a61a756ea38475554097ac04612037c36ffe52"
	require.NoError(t, err1)
	require.NoError(t, err2)
	require.Equal(t, want1, resp1.Policy.Id)
	require.Equal(t, want2, resp2.Policy.Id)
}

*/
