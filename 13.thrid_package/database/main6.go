package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DbConnection *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	cmd := "DELETE FROM persons WHERE name = ?"
	_, err := DbConnection.Exec(cmd, "Nancy")
	if err != nil {
		log.Fatalln(err)
	}
}
