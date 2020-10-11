// Copyright 2020 entooone
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fmttable

import (
	"strings"
	"testing"
)

func isValid(a, b Table) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(b); i++ {
		if len(a[i]) != len(b[i]) {
			return false
		}

		for j := 0; j < len(b[i]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func TestReadTableMD(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		rawTable string
		want     Table
	}{
		"empty": {
			rawTable: "",
			want:     Table([][]string{}),
		},
		"single line": {
			rawTable: `| a | b | c |`,
			want:     Table([][]string{{"a", "b", "c"}}),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := ReadTableMD(strings.NewReader(test.rawTable))
			if err != nil {
				t.Fatalf("%s (error: %v)", name, err)
			}

			if !isValid(got, test.want) {
				t.Fatalf("%s (want: %v, got: %v)", name, test.want, got)
			}
		})
	}
}
