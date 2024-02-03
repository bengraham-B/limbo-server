package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (F *FILE_STRUCT) SQL_InsertOneFile() int64 {
	db, err := CONN_DB() //& Connecting to the database

	if err != nil {
		gooseLogger("Error", 1, "Openning conection to DB", "DB_insertFile.go", err)
		return 0
	}

	query := fmt.Sprintf(`INSERT INTO limbo(user_id, void, file_name, file_type, file_content) VALUES('%v', '%v', '%v', '%v', '%v')`, F.UserID, F.Void, F.FileName, F.FileType, F.FileContent)

	result, err := db.Exec(query)

	if err != nil {
		gooseLogger("Error", 2, "Inserting data into to DB", "DB_insertFile.go", err)

	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		gooseLogger("Error", 3, "Getting data inserted into to DB", "DB_insertFile.go", err)

	}

	gooseLogger("Success", 4, fmt.Sprintf("Inserted %d rows into DB", rowsAffected), "DB_insertFile.go", nil)

	defer db.Close()

	return rowsAffected //^ Returns 1 if success

}

func (F *FILE_STRUCT) DB_InsertOneFile(c *gin.Context) {
	var requestBody FILE_STRUCT

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	F.updateFileStruct(requestBody.UserID, requestBody.Void, requestBody.FileName, requestBody.FileType, requestBody.FileContent)

	queryResult := F.SQL_InsertOneFile()

	c.JSON(http.StatusOK, gin.H{"message": "Request processed successfully", "Result": queryResult})

}
