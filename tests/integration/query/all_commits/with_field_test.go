// Copyright 2022 Democratized Data Foundation
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package all_commits

import (
	"testing"

	testUtils "github.com/sourcenetwork/defradb/tests/integration"
)

// This test is for documentation reasons only. This is not
// desired behaviour (should return all commits for field).
func TestQueryAllCommitsWithField(t *testing.T) {
	test := testUtils.QueryTestCase{
		Description: "Simple all commits query with field",
		Query: `query {
					allCommits (field: "Age") {
						cid
					}
				}`,
		Docs: map[int][]string{
			0: {
				`{
					"Name": "John",
					"Age": 21
				}`,
			},
		},
		ExpectedError: "Field \"allCommits\" argument \"dockey\" of type \"ID!\" is required but not provided.",
	}

	executeTestCase(t, test)
}

// This test is for documentation reasons only. This is not
// desired behaviour (should return all commits for field).
func TestQueryAllCommitsWithFieldId(t *testing.T) {
	test := testUtils.QueryTestCase{
		Description: "Simple all commits query with field id",
		Query: `query {
					allCommits (field: "1") {
						cid
					}
				}`,
		Docs: map[int][]string{
			0: {
				`{
					"Name": "John",
					"Age": 21
				}`,
			},
		},
		ExpectedError: "Field \"allCommits\" argument \"dockey\" of type \"ID!\" is required but not provided.",
	}

	executeTestCase(t, test)
}
