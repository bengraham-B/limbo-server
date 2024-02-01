package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// ^ Connect to Database
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "limbo"
)

func connectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	//^ Oppenning the connection
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Println("Error openning DB connection", err)
		return nil, err
	}

	fmt.Println("")
	fmt.Printf("Connected to Database: [%v] 🐳\n", dbname)
	fmt.Println("")

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db, nil
}
