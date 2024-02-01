package SQL_LIMBO

import "fmt"

func SQL_readOneFile(fileName string) (file, error) {
	db, err := connectDB()

	if err != nil {
		fmt.Println("Cannot connect to database in [_readOneFile.go_]:", err)
	}

	query := `SELECT * FROM limbo(file_name, file_type, file_content) WHERE file_name =$1`

	rows, err := db.Query(query, fileName)

	if err != nil {
		fmt.Printf("Error reading data file name: (%v) -  [_insertData.go_]: %v", fileName, err)
	}

	fmt.Println(rows)

	defer db.Close()

	return FileData{}, fmt.Errorf("No data found for file name %s", fileName)

}
