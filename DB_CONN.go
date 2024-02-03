package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// ^ Connect to Database
const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "root"
	dbname   = "limbo"
)

func CONN_DB() (*sql.DB, error) {

	//^ Connection String
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	//^ Openning the Connection
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Printf("[ERROR: 1] - [Cannot open DB: %v] - [File: DB_CONN] - (ERROR: %v)", dbname, err)
		gooseLogger("Error", 1, "Cannot Open DB: Limbo", "DB_CONN.go", err)

		return nil, err
	}

	fmt.Println("")
	fmt.Printf("Connected to Database: [%v] 🐳\n", dbname)
	fmt.Println("")

	//^ Pinning the DB
	err = db.Ping()
	if err != nil {
		gooseLogger("Error", 2, "Cannot ping DB: Limbo", "DB_CONN.go", err)
		return nil, err
	}

	gooseLogger("Success", 3, "Returning DB variable", "DB_CONN.go", nil)
	return db, err

}
