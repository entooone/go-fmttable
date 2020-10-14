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
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/mattn/go-runewidth"
)

// Table represents a table of strings in two dimensions
type Table [][]string

// Equal reports whether x and y are equal
func (x Table) Equal(y Table) bool {
	if len(x) != len(y) {
		return false
	}

	for i := 0; i < len(y); i++ {
		if len(x[i]) != len(y[i]) {
			return false
		}

		for j := 0; j < len(y[i]); j++ {
			if x[i][j] != y[i][j] {
				return false
			}
		}
	}
	return true
}

// Pretty writes the Markdown table to w
func (t Table) Pretty(w io.Writer) {
	if len(t) == 0 {
		return
	}

	maxCol := len(t[0])
	for _, row := range t {
		if len(row) > maxCol {
			maxCol = len(row)
		}
	}

	colSizes := make([]int, maxCol)
	for _, row := range t {
		for i, v := range row {
			vlen := runewidth.StringWidth(v)
			if colSizes[i] < vlen {
				colSizes[i] = vlen
			}
		}
	}

	for _, row := range t {
		if len(row) == 0 {
			continue
		}

		for i, v := range row {
			fv := runewidth.FillRight(v, colSizes[i])
			fmt.Fprintf(w, "| %s ", fv)
		}
		fmt.Fprintf(w, "|\n")
	}
}

func readLineMD(line string) []string {
	if line == "" {
		return []string{}
	}
	line = strings.Trim(line, "\t |")
	es := strings.Split(line, "|")
	for i, e := range es {
		es[i] = strings.Trim(e, " ")
	}

	return es
}

// ReadTableMD loads Table from a Markdown table
func ReadTableMD(r io.Reader) (Table, error) {
	s, err := ioutil.ReadAll(r)
	if err != nil {
		return Table{}, err
	}

	lines := strings.Split(string(s), "\n")
	table := make(Table, 0, len(lines))

	for _, line := range lines {
		lm := readLineMD(line)
		if len(lm) == 0 {
			continue
		}
		table = append(table, lm)
	}

	return table, nil
}
