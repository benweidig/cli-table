package main

import (
	"fmt"

	"github.com/benweidig/cli-table"
)

func main() {

	table := clitable.New()

	table.AddRow("COL1", "COL2", "COL3", "COL4")
	table.AddRow("This is", "the first", "row?", 42)
	table.AddRow("A", "shorter", "row")
	table.AddRow("And an even", "longer", "one", "with", "more", "columns than the first")

	fmt.Println(table.String())
}
