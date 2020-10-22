package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var (
		id      string
		balance int
	)

	rows, err := db.Query("select _id, balance from accounts where _id  = ?", "59f6ace10b0fb62461985888")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &balance)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, balance)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}
