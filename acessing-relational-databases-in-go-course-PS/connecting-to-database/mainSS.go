package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/microsoft/go-mssqldb"
)

var db *sql.DB

func main() {

	ConnString := "server=localhost;user id=root;password=root;port=1433;database=sakila;"

	var err error
	db, err = sql.Open("sqlserver", ConnString)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

}
