package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type Parameters struct {
	foreign_key_pattern string
	db_user             string
	db_pwd              string
	db_host             string
	db_port             string
	db_name             string
}

type Application struct {
	parameters Parameters
	db         *sql.DB
}

func (a *Application) Load() func() {
	a.loadParameters()
	a.connectDB()

	return func() {
		a.db.Close()
	}
}

func (a *Application) loadParameters() {
	a.loadEnvParameters()

	if a.parameters.foreign_key_pattern == "" {
		fmt.Print("\nWhat is the foreign key pattern (replace the table name by '%v'. eg: 'id_%v')?\n> ")
		fmt.Scanf("%s", &a.parameters.foreign_key_pattern)
	}
	if a.parameters.db_user == "" {
		fmt.Print("\nDB user:\n> ")
		fmt.Scanf("%s", &a.parameters.db_user)
	}
	if a.parameters.db_pwd == "" {
		fmt.Print("\nDB password\n> ")
		fmt.Scanf("%s", &a.parameters.db_pwd)
	}
	if a.parameters.db_host == "" {
		fmt.Print("\nDB host:\n> ")
		fmt.Scanf("%s", &a.parameters.db_host)
	}
	if a.parameters.db_port == "" {
		fmt.Print("\nDB port:\n> ")
		fmt.Scanf("%s", &a.parameters.db_port)
	}
	if a.parameters.db_name == "" {
		fmt.Print("\nDB name:\n> ")
		fmt.Scanf("%s", &a.parameters.db_name)
	}

	fmt.Println()
}

func (a *Application) loadEnvParameters() {
	a.parameters.foreign_key_pattern = os.Getenv("FOREIGN_KEY_PATTERN")
	a.parameters.db_user = os.Getenv("DB_USER")
	a.parameters.db_pwd = os.Getenv("DB_PWD")
	a.parameters.db_host = os.Getenv("DB_HOST")
	a.parameters.db_port = os.Getenv("DB_PORT")
	a.parameters.db_name = os.Getenv("DB_NAME")
}

func (a *Application) connectDB() {
	db, err := sql.Open("mysql", fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v",
		a.parameters.db_user,
		a.parameters.db_pwd,
		a.parameters.db_host,
		a.parameters.db_port,
		a.parameters.db_name,
	))

	if err != nil {
		log.Fatal(err)
	}

	a.db = db
}
