package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

var err error

type Person struct {
	Name string
	Age  int
}

func main() {
	Db, err = sql.Open("postgres", "user=test_user dbname=testdb password=password sslmode=disable")
	if err != nil {
		log.Panicln(err)
	}
	defer Db.Close()

	/*
		//C
		cmd := "INSERT INTO persons (name, age) VALUES ($1, $2)"
		_, err := Db.Exec(cmd, "Nancy", 20)
		if err != nil {
			log.Fatalln(err)
		}
	*/

	/*
		//R
		cmd := "SELECT * FROM persons where age = $1"
		row := Db.QueryRow(cmd, 20)
		var p Person
		err = row.Scan(&p.Name, &p.Age)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Println("No row")
			} else {
				log.Println(err)
			}
		}
		fmt.Println(p.Name, p.Age)

		cmd = "SELECT * FROM persons"
		rows, _ := Db.Query(cmd)
		defer rows.Close()
		var pp []Person
		for rows.Next() {
			var p Person
			err := rows.Scan(&p.Name, &p.Age)
			if err != nil {
				log.Println(err)
			}
			pp = append(pp, p)
		}
		err = rows.Err()
		if err != nil {
			log.Fatalln(err)
		}
		for _, p := range pp {
			fmt.Println(p.Name, p.Age)
		}
	*/

	/*
		//U
		cmd := "UPDATE persons SET age = $1 WHERE name = $2"
		_, err := Db.Exec(cmd, 25, "Nancy")
		if err != nil {
			log.Fatalln(err)
		}
	*/

	//D
	cmd := "DELETE FROM persons WHERE name = $1"
	_, err := Db.Exec(cmd, "Nancy")
	if err != nil {
		log.Fatalln(err)
	}

}
