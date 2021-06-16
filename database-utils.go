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
		primary_name := fmt.Sprintf(app.parameters.primary_key_pattern, table)
		foreign_name := fmt.Sprintf(app.parameters.foreign_key_pattern, table)

		var nb int

		row := app.db.QueryRow(`
			SELECT COUNT(*) count
			FROM ` + deps[i] + ` t
			LEFT JOIN ` + table + ` d ON t.` + foreign_name + ` = d.` + primary_key_pattern + `
			WHERE d.` + primary_key_pattern + ` IS NULL
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
