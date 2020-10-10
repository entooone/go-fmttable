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

func isValid(got, want Table) bool {
	if len(got) != len(want) {
		return false
	}

	for i := 0; i < len(want); i++ {
		if len(got[i]) != len(want[i]) {
			return false
		}

		for j := 0; j < len(want[i]); j++ {
			if got[i][j] != want[i][j] {
				return false
			}
		}
	}
	return true
}

func TestReadTableMD(t *testing.T) {
	data := map[string]struct {
		rawTable string
		want     Table
	}{
		"Read Markdown Table": {
			rawTable: `| a | b | c |`,
			want:     Table([][]string{{"a", "b", "c"}}),
		},
	}

	for s, v := range data {
		got, err := ReadTableMD(strings.NewReader(v.rawTable))
		if err != nil {
			t.Fatalf("%s (error: %v)", s, err)
		}

		if !isValid(got, v.want) {
			t.Fatalf("%s (want: %v, got: %v)", s, v.want, got)
		}
	}
}
