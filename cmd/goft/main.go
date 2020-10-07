package main

import (
	"log"
	"os"

	"github.com/entooone/go-ftable"
)

func main() {
	table, err := ftable.ReadTableMD(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	table.Pretty()
}
