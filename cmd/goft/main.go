package main

import (
	"log"
	"os"

	"github.com/entooone/go-fmttable"
)

func main() {
	table, err := fmttable.ReadTableMD(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	table.Pretty(os.Stdout)
}
