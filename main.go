package main

import (
	"github.com/joho/godotenv"
)

var app = Application{}

func init() {
	godotenv.Load(".env")
}

func main() {
	close_db := app.Load()
	defer close_db()

	findBrokenDependenciesForAllTables()
}
