package main

import (
	"fmt"
	"log"
)

type BrokenDependingTable struct {
	name  string
	count int
}

func findAllTables() []string {
	var t []string

	res, err := app.db.Query(`
		SELECT DISTINCT table_name
		FROM information_schema.tables
		WHERE table_schema = database()
	`)

	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {
		var n string
		err := res.Scan(&n)

		if err != nil {
			log.Fatal(err)
		}

		t = append(t, n)
	}

	return t
}

func findDependingTables(table string) []string {
	var t []string

	column_name := fmt.Sprintf(app.parameters.foreign_key_pattern, table)

	res, err := app.db.Query(`
		SELECT DISTINCT table_name
		FROM information_schema.columns
		WHERE
			column_name LIKE '` + column_name + `' AND
			table_name <> '` + table + `' AND
			table_schema = database()
	`)

	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {
		var n string
		err := res.Scan(&n)

		if err != nil {
			log.Fatal(err)
		}

		t = append(t, n)
	}

	return t
}

func findBrokenDependenciesForTable(table string) []BrokenDependingTable {
	var t []BrokenDependingTable
	deps := findDependingTables(table)

	for i := range deps {
		foreign_name := fmt.Sprintf(app.parameters.foreign_key_pattern, table)

		var nb int

		row := app.db.QueryRow(`
			SELECT COUNT(*) count
			FROM ` + deps[i] + ` t
			LEFT JOIN ` + table + ` d ON t.` + foreign_name + ` = d.` + foreign_name + `
			WHERE d.` + foreign_name + ` IS NULL
		`)
		err := row.Scan(&nb)

		if err != nil {
			fmt.Printf("%v", err)
		}

		if nb > 0 {
			t = append(t, BrokenDependingTable{
				name:  deps[i],
				count: nb,
			})
		}
	}

	return t
}
