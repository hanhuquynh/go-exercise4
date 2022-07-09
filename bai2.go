package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func transaction() {
	db, err := sql.Open("mysql", "root:hanhuquynh@tcp(127.0.0.1:3306)/test")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}
