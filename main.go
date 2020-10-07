package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/mattn/go-runewidth"
)

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

func readTableMD(r io.Reader) (Table, error) {
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

type Table [][]string

func (t Table) Pretty() {
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
		for i, v := range row {
			fv := runewidth.FillRight(v, colSizes[i])
			fmt.Printf("| %s ", fv)
		}
		fmt.Printf("|\n")
	}
}

var table Table

func main() {
	table, err := readTableMD(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	table.Pretty()
}
