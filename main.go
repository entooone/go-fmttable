package main

import (
	"fmt"

	"github.com/mattn/go-runewidth"
)

type Table [][]string

func (t Table) Pretty() {
	colSizes := make([]int, len(t[0]))
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
	table = Table([][]string{
		{"abc", "b", "c"},
		{"d", "e", "hello"},
		{"", "abcあいう", "日本語"},
		{"こんにちは", "hello, world", "hello"},
	})
	table.Pretty()
}
