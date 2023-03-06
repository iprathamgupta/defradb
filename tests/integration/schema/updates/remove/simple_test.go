// Copyright 2023 Democratized Data Foundation
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package remove

import (
	"testing"

	testUtils "github.com/sourcenetwork/defradb/tests/integration"
)

func TestSchemaUpdatesRemoveCollectionNameErrors(t *testing.T) {
	test := testUtils.TestCase{
		Description: "Test schema update, remove collection name",
		Actions: []any{
			testUtils.SchemaUpdate{
				Schema: `
					type Users {
						Name: String
						Email: String
					}
				`,
			},
			testUtils.SchemaPatch{
				Patch: `
					[
						{ "op": "remove", "path": "/Users/Name" }
					]
				`,
				ExpectedError: "collection name can't be empty",
			},
		},
	}
	testUtils.ExecuteTestCase(t, []string{"Users"}, test)
}

func TestSchemaUpdatesRemoveCollectionIDErrors(t *testing.T) {
	test := testUtils.TestCase{
		Description: "Test schema update, remove collection id",
		Actions: []any{
			testUtils.SchemaUpdate{
				Schema: `
					type Users {
						Name: String
						Email: String
					}
				`,
			},
			testUtils.SchemaPatch{
				Patch: `
					[
						{ "op": "remove", "path": "/Users/ID" }
					]
				`,
				ExpectedError: "CollectionID does not match existing. Name: Users, ExistingID: 1, ProposedID: 0",
			},
		},
	}
	testUtils.ExecuteTestCase(t, []string{"Users"}, test)
}

func TestSchemaUpdatesRemoveSchemaIDErrors(t *testing.T) {
	test := testUtils.TestCase{
		Description: "Test schema update, remove schema ID",
		Actions: []any{
			testUtils.SchemaUpdate{
				Schema: `
					type Users {
						Name: String
						Email: String
					}
				`,
			},
			testUtils.SchemaPatch{
				Patch: `
					[
						{ "op": "remove", "path": "/Users/Schema/SchemaID" }
					]
				`,
				ExpectedError: "SchemaID does not match existing",
			},
		},
	}
	testUtils.ExecuteTestCase(t, []string{"Users"}, test)
}

func TestSchemaUpdatesRemoveSchemaVersionIDErrors(t *testing.T) {
	test := testUtils.TestCase{
		Description: "Test schema update, remove schema version id",
		Actions: []any{
			testUtils.SchemaUpdate{
				Schema: `
					type Users {
						Name: String
						Email: String
					}
				`,
			},
			testUtils.SchemaPatch{
				// This should do nothing
				Patch: `
					[
						{ "op": "remove", "path": "/Users/Schema/VersionID" }
					]
				`,
			},
			testUtils.Request{
				Request: `query {
					Users {
						Name
						Email
					}
				}`,
				Results: []map[string]any{},
			},
		},
	}
	testUtils.ExecuteTestCase(t, []string{"Users"}, test)
}

func TestSchemaUpdatesRemoveSchemaNameErrors(t *testing.T) {
	test := testUtils.TestCase{
		Description: "Test schema update, remove schema name",
		Actions: []any{
			testUtils.SchemaUpdate{
				Schema: `
					type Users {
						Name: String
						Email: String
					}
				`,
			},
			testUtils.SchemaPatch{
				Patch: `
					[
						{ "op": "remove", "path": "/Users/Schema/Name" }
					]
				`,
				ExpectedError: "modifying the schema name is not supported. ExistingName: Users, ProposedName: ",
			},
		},
	}
	testUtils.ExecuteTestCase(t, []string{"Users"}, test)
}
