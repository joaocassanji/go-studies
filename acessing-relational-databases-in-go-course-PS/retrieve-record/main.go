package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	dsn := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "sakila",
		AllowNativePasswords: true,
	}

	var err error
	db, err = sql.Open("mysql", dsn.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	if err != nil {
		log.Fatal(err)
	}

	actors, err := getActor(201)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Actor found: %v\n", actors)

}

type Actor struct {
	actor_id   int64
	first_name string
	last_name  string
}

func getActor(actorID int32) ([]Actor, error) {
	var actors []Actor

	result, err := db.Query("SELECT actor_id, first_name, last_name FROM actor WHERE actor_id = ?", actorID)
	if err != nil {
		return nil, fmt.Errorf("getActor: %v, %v", actorID, err)
	}
	defer result.Close()

	for result.Next() {
		var acts Actor
		if err := result.Scan(&acts.actor_id, &acts.first_name, &acts.last_name); err != nil {
			return nil, fmt.Errorf("GetActor %v: %v", actorID, err)
		}
		actors = append(actors, acts)
		if err := result.Err(); err != nil {
			return nil, fmt.Errorf("GetActor %v: %v", actorID, err)
		}
	}

	return actors, nil
}
