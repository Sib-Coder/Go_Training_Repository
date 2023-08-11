package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	db, err := sqlx.Open(
		"postgres",
		"user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	stmt, err := db.Prepare("INSERT INTO  words (name_task, translation, image, phrase, weight, date_insert, date_done) VALUES ('HelloWorld', 'Привет Мир!', 'https://mleern:80/images/test.png','Hello World!', '1', '2005-04-02', '2005-04-01') RETURNING id_task;")
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	var id int
	err = stmt.QueryRow().Scan(&id)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	fmt.Printf("inserted object's ID: %d\n", id)
}
