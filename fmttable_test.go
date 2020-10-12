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
	"bytes"
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
			want:     Table{},
		},
		"single line": {
			rawTable: `| a | b | c |`,
			want:     Table{{"a", "b", "c"}},
		},
		"multi line": {
			rawTable: `
			| a | b | c |
			| golang | hello | gopher |`,
			want: Table{{"a", "b", "c"}, {"golang", "hello", "gopher"}},
		},
		"uneven table": {
			rawTable: `
			| a | b |
			| foo | bar | baz | qux | quux |
			| golang | hello | gopher |`,
			want: Table{
				{"a", "b"},
				{"foo", "bar", "baz", "qux", "quux"},
				{"golang", "hello", "gopher"},
			},
		},
	}

	for name, test := range tests {
		test := test
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

func TestPrintMD(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		table Table
		want  string
	}{
		"empty": {
			table: Table{},
			want:  "",
		},
		"skip empty line": {
			table: Table{{"a", "b"}, {}, {"c", "d"}},
			want:  "| a | b |\n| c | d |\n",
		},
		"single line": {
			table: Table{{"a", "b", "c"}},
			want:  "| a | b | c |\n",
		},
		"multi line": {
			table: Table{{"a", "b", "c"}, {"golang", "hello", "gopher"}},
			want:  "| a      | b     | c      |\n| golang | hello | gopher |\n",
		},
		"uneven table": {
			table: Table{
				{"a", "b"},
				{"foo", "bar", "baz", "qux", "quux"},
				{"golang", "hello", "gopher"},
			},
			want: "| a      | b     |\n| foo    | bar   | baz    | qux | quux |\n| golang | hello | gopher |\n",
		},
		"include japanese": {
			table: Table{
				{"a", "b"},
				{"hello", "world"},
				{"こんにちは", "世界"},
			},
			want: "| a          | b     |\n| hello      | world |\n| こんにちは | 世界  |\n",
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			buf := new(bytes.Buffer)
			test.table.Pretty(buf)
			got := buf.String()
			if got != test.want {
				t.Fatalf("%s (want: %#v, got: %#v)", name, test.want, got)
			}
		})
	}
}
