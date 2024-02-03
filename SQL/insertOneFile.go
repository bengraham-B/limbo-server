package SQL_LIMBO

import (
	"fmt"
)

func SQL_insertData(fileName string, fileType string, fileContent string) {
	db, err := connectDB() //^ Connect to the db

	if err != nil {
		fmt.Println("Cannot connect to database in [_insertData.go_]:", err)
	}

	query := fmt.Sprintf(`INSERT INTO limbo(file_name, file_type, file_content) VALUES('%v', '%v', '%v')`, fileName, fileType, fileContent)

	rows, err := db.Query(query)

	if err != nil {
		fmt.Printf("Error inserting data file name: (%v) -  [_insertData.go_]: %v", fileName, err)
	}

	fmt.Println(rows)

	defer db.Close()
}
