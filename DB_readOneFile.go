package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (F *FILE_STRUCT) SQL_ReadOneFile() map[string]string {
	db, err := CONN_DB() //& Connecting to the database

	if err != nil {
		gooseLogger("Error", 1, "Openning conection to DB", "DB_readOneFile.go", err)
		return nil
	}

	query := fmt.Sprintf(`SELECT * FROM limbo WHERE user_id='%v' AND void='%v' AND file_name='%v'`, F.UserID, F.Void, F.FileName)

	queryResult := map[string]interface{}{
		"userID":      new(string),
		"void":        new(string),
		"fileName":    new(string),
		"fileType":    new(string),
		"fileContent": new(string),
	}

	err = db.QueryRow(query).Scan(queryResult["userID"], queryResult["void"], queryResult["void"], queryResult["fileName"], queryResult["fileType"], queryResult["fileContent"])
	fmt.Println(err)
	if err != nil {
		if err == sql.ErrNoRows {
			gooseLogger("Error", 2, "No rows found with ID:", "DB_readOneFile.go", err)
			return nil
		}
		gooseLogger("Error", 3, "Cannot read row from DB", "DB_readOneFile.go", err)
		return nil
	}

	// Access the scanned values using type assertion
	result := map[string]string{
		"userID":      *queryResult["userID"].(*string),
		"void":        *queryResult["void"].(*string),
		"fileName":    *queryResult["fileName"].(*string),
		"fileType":    *queryResult["fileType"].(*string),
		"fileContent": *queryResult["fileContent"].(*string),
	}

	gooseLogger("Success", 4, "Read from DB", "DB_readOneFile.go", err)
	return result
}

func (F *FILE_STRUCT) DB_ReadoneFile(c *gin.Context) {
	var requestBody FILE_STRUCT

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	F.updateFileStruct(requestBody.UserID, requestBody.Void, requestBody.FileName, requestBody.FileType, requestBody.FileContent)

	row := F.SQL_ReadOneFile()
	fmt.Println(row)

	c.JSON(http.StatusOK, gin.H{"message": "Request processed successfully", "Result": row})

}
