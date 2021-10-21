// Copyright 2020 Source Inc.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.
package simple

import (
	"testing"

	testUtils "github.com/sourcenetwork/defradb/db/tests"
)

func TestQuerySimpleWithEmbeddedLatestCommit(t *testing.T) {
	test := testUtils.QueryTestCase{
		Description: "Embedded latest commits query within object query",
		Query: `query {
					users {
						Name
						Age
						_version {
							cid
							links {
								cid
								name
							}
						}
					}
				}`,
		Docs: map[int][]string{
			0: {
				(`{
				"Name": "John",
				"Age": 21
			}`)},
		},
		Results: []map[string]interface{}{
			{
				"Name": "John",
				"Age":  uint64(21),
				"_version": []map[string]interface{}{
					{
						"cid": "QmaXdKKsc5GRWXtMytZj4PEf5hFgFxjZaKToQpDY8cAocV",
						"links": []map[string]interface{}{
							{
								"cid":  "QmPaY2DNmd7LtRDpReswc5UTGoU5Q32Py1aEVG7Shq6Np1",
								"name": "Age",
							},
							{
								"cid":  "Qmag2zKKGGQwVSss9pQn3hjTu9opdF5mkJXUR9rt2A651h",
								"name": "Name",
							},
						},
					},
				},
			},
		},
	}

	executeTestCase(t, test)
}
