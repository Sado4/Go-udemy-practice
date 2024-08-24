package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DbConnection *sql.DB

func main() {
	cmd := "UPDATE persons SET age = ? WHERE name = ?"
	_, err := DbConnection.Exec(cmd, 25, "Mike")
	if err != nil {
		log.Fatalln(err)
	}
}
