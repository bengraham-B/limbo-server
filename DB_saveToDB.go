package main

import (
	"fmt"
	SQL_LIMBO "main/SQL"
	"net/http"

	"github.com/gin-gonic/gin"
)

type file struct {
	FileName    string `json: "fileName"`
	FileType    string `json: "fileType"`
	FileContent string `json: "fileContent"`
}

func saveToDB(c *gin.Context) {
	var requestBody file

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fileName := requestBody.FileName
	fileType := requestBody.FileType
	fileContent := requestBody.FileContent

	fmt.Println("Recieved:", fileName, fileType, fileContent)

	SQL_LIMBO.SQL_insertData(fileName, fileType, fileContent)

	c.JSON(http.StatusOK, gin.H{"message": "Request processed successfully"})

}
