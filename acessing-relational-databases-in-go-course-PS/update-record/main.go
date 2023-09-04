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

	rowsUpdated, err := updateActor("JAMES", 201)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Total actor affected: %d\n", rowsUpdated)

}

func updateActor(firstname string, actorid int32) (int64, error) {

	result, err := db.Exec("UPDATE actor SET first_name =? WHERE actor_id = ?", firstname, actorid)
	if err != nil {
		return 0, fmt.Errorf("updateActor: %v", err)
	}
	id, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("updateActor: %v", err)
	}

	return id, nil
}
