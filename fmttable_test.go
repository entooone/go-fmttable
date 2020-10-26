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
	"errors"
	"strings"
	"testing"
)

func TestEqual(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		t1   Table
		t2   Table
		want bool
	}{
		"same table": {
			t1:   Table{{"a"}, {"b", "c", "d"}, {"e", "f"}},
			t2:   Table{{"a"}, {"b", "c", "d"}, {"e", "f"}},
			want: true,
		},
		"different length table": {
			t1:   Table{},
			t2:   Table{{}},
			want: false,
		},
		"different length table 2": {
			t1:   Table{{}},
			t2:   Table{{"a"}},
			want: false,
		},
		"has different element": {
			t1:   Table{{"a"}},
			t2:   Table{{"b"}},
			want: false,
		},
	}
	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := test.t1.Equal(test.t2)
			if got != test.want {
				t.Fatalf("%s (want: %v, got: %v)", name, test.want, got)
			}
		})
	}
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

			if !got.Equal(test.want) {
				t.Fatalf("%s (want: %v, got: %v)", name, test.want, got)
			}
		})
	}
}

type MockFailReader struct{}

func (m *MockFailReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("this mock always returns an error")
}

func TestReadTableMDFail(t *testing.T) {
	t.Parallel()
	t.Run("error", func(t *testing.T) {
		t.Parallel()
		_, err := ReadTableMD(&MockFailReader{})
		if err == nil {
			t.Fatalf("the test should have error")
		}
	})
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
			test.table.WriteMDFormat(buf)
			got := buf.String()
			if got != test.want {
				t.Fatalf("%s (want: %#v, got: %#v)", name, test.want, got)
			}
		})
	}
}

func TestPrintCSV(t *testing.T) {
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
			want:  "a, b\nc, d\n",
		},
		"single line": {
			table: Table{{"a", "b", "c"}},
			want:  "a, b, c\n",
		},
		"multi line": {
			table: Table{{"a", "b", "c"}, {"golang", "hello", "gopher"}},
			want:  "a     , b    , c     \ngolang, hello, gopher\n",
		},
		"uneven table": {
			table: Table{
				{"a", "b"},
				{"foo", "bar", "baz", "qux", "quux"},
				{"golang", "hello", "gopher"},
			},
			want: "a     , b    \nfoo   , bar  , baz   , qux, quux\ngolang, hello, gopher\n",
		},
		"include japanese": {
			table: Table{
				{"a", "b"},
				{"hello", "world"},
				{"こんにちは", "世界"},
			},
			want: "a         , b    \nhello     , world\nこんにちは, 世界 \n",
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			buf := new(bytes.Buffer)
			test.table.WriteCSVFormat(buf)
			got := buf.String()
			if got != test.want {
				t.Fatalf("%s (want: %#v, got: %#v)", name, test.want, got)
			}
		})
	}
}
