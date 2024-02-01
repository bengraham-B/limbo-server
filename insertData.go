package main

import "fmt"

func insertData(fileName string, fileContent string) {
	db, err := connectDB() //^ Connect to the db

	if err != nil {
		fmt.Println("Cannot connect to database in [_insertData.go_]:", err)
	}

	query := fmt.Sprintf(`INSERT INTO limbo(file_name, file_content) VALUES('%v', '%v')`, fileName, fileContent)

	rows, err := db.Query(query)

	if err != nil {
		fmt.Printf("Error inserting data file name: (%v) -  [_insertData.go_]: %v", fileName, err)
	}

	fmt.Println(rows)

	defer db.Close()
}
