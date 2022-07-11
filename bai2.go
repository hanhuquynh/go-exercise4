package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func transaction(id, birth int) {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/test")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	tx, err := db.Begin()

	if err != nil {
		log.Fatal(err)
	}

	_, err = tx.Exec("UPDATE user SET birth = ? WHERE id = ?", birth, id)

	if err != nil {
		tx.Rollback()
		fmt.Println("Update birth failed:: ", err)
		return
	} else {
		fmt.Println("Update birth successfully")
	}

	point, err := db.Query("SELECT points FROM point WHERE user_id = ?", id)

	if err != nil {
		log.Fatal(err)
	}

	var getPoint int

	for point.Next() {
		err = point.Scan(&getPoint)

		if err != nil {
			log.Fatal(err)
		}
	}

	_, err = tx.Exec("UPDATE point SET points = ? WHERE user_id = ?", getPoint+10, id)

	if err != nil {
		tx.Rollback()
		fmt.Println("Update point failed: ", err)
		return
	} else {
		fmt.Println("Update point successfully")
	}

	rows, err := db.Query("SELECT name, updated_at FROM user WHERE id = ?", id)

	if err != nil {
		log.Fatal(err)
	}

	var name string
	var updated_at int

	for rows.Next() {
		err = rows.Scan(&name, &updated_at)

		if err != nil {
			log.Fatal(err)
		}
	}

	var newName = name + " - " + strconv.Itoa(updated_at)

	_, err = tx.Exec("UPDATE user SET name = ? WHERE id = ?", newName, id)

	if err != nil {
		tx.Rollback()
		fmt.Println("Update name failed: ", err)
		return
	} else {
		fmt.Println("Update name successfully")
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
