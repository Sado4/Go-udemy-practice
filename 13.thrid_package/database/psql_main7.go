package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var Db *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	Db, _ = sql.Open("postgres", "user=aoyagimasanori dbname=test_db password=genius0610 sslmode=disable")

	defer Db.Close()

}
