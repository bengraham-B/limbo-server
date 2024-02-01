package main

import (
	"fmt"
	"log"
)

func createTable() {
	db, err := connectDB()

	if err != nil {
		fmt.Println(err)
		return
	}

	rows, err := db.Query(`CREATE TABLE IF NOT EXISTS limbo(
		pk SERIAL PRIMARY KEY,
		id UUID,
	
		file_name VARCHAR(255),
		file_Content TEXT
	)`)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rows)
	defer db.Close()

}
