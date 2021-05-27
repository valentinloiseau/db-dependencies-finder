package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/raphamorim/go-rainbow"
)

func findBrokenDependenciesForAllTables() {
	tables := findAllTables()

	for i := range tables {
		broken_tables := findBrokenDependenciesForTable(tables[i])

		if len(broken_tables) > 0 {
			fmt.Printf(
				"Found %v broken dependencies for table %v:\n",
				rainbow.Bold(fmt.Sprint(len(broken_tables))),
				rainbow.Bold(tables[i]),
			)

			for y := range broken_tables {
				fmt.Printf(
					" -> %v errors in table %v\n",
					rainbow.Bold(fmt.Sprint(broken_tables[y].count)),
					rainbow.Red(rainbow.Bold(broken_tables[y].name)),
				)
			}

			fmt.Println()
		}
	}
}
