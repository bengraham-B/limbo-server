package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (F *FILE_STRUCT) SQL_InsertOneFile() {
	db, err := CONN_DB() //& Connecting to the database

	if err != nil {
		gooseLogger("Error", 1, "Openning conection to DB", "DB_insertFile.go", err)
		return
	}

	query := fmt.Sprintf(`INSERT INTO limbo(file_name, file_type, file_content) VALUES('%v', '%v', '%v')`, F.FileName, F.FileType, F.FileContent)
	rows, err := db.Query(query)

	gooseLogger("Error", 2, "Inserting data into to DB", "DB_insertFile.go", err)

	fmt.Println(rows)

	defer db.Close()

}

func (F *FILE_STRUCT) DB_InsertOneFile(c *gin.Context) {
	var requestBody FILE_STRUCT

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	F.updateFileStruct(requestBody.FileName, requestBody.FileType, requestBody.FileContent)

	F.SQL_InsertOneFile()

	c.JSON(http.StatusOK, gin.H{"message": "Request processed successfully"})

}
