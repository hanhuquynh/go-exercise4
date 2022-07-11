package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func createDB(name string) {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/")

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connect succcessfully")
	}

	defer db.Close()

	_, err = db.Exec("CREATE DATABASE " + name)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Create database succcessfully!")
	}

	_, err = db.Exec("USE " + name)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE user (id int, name varchar(40), birth int, created int, updated_at int)")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Create table user succcessfully!")
	}

	_, err = db.Exec("CREATE TABLE point (user_id int, points int, max_points int)")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Create table point succcessfully!")
	}

}

func insertUser(id int, name string, birth, created, updated_at int) {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/test")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	ins, err := db.Query("INSERT INTO user (id, name, birth, created, updated_at) VALUES(?, ?, ?, ?, ?)", id, name, birth, created, updated_at)

	if err != nil {
		log.Fatal(err)
	}

	afterIns, err := db.Query("INSERT INTO point (user_id, points) VALUES(?, 10)", id)

	if err != nil {
		log.Fatal(err)
	}

	defer ins.Close()
	defer afterIns.Close()

}

func updateUser(id int, name string, birth int, created int, updated_at int) {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/test")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	upd, err := db.Query("UPDATE user SET id = ?, name = ?, birth = ?, created = ?, updated_at = ? WHERE id = ?", id, name, birth, created, updated_at, id)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Update user successfully")
	}

	defer upd.Close()
}

func listUser() {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/test")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM user ")

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Birth, &user.Created, &user.Updated_at)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(user)
	}
}

func listUserById(id int) {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/test")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM user WHERE id = ?", id)

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Birth, &user.Created, &user.Updated_at)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(user)
	}
}
