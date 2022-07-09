package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func transaction(id, birth int) {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/test")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	ctx := context.Background()

	tx, err := db.BeginTx(ctx, nil)

	if err != nil {
		panic(err.Error())
	}

	_, err = tx.ExecContext(ctx, "UPDATE user SET birth = ? WHERE id = ?", birth, id)

	if err != nil {
		tx.Rollback()
		return
	}

	_, err = tx.ExecContext(ctx, "UPDATE point SET points = ? WHERE id = ?", 20, id)

	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
