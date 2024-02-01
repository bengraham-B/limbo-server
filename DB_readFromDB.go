package main

import (
	"fmt"
	SQL_LIMBO "main/SQL"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DB_readOneFile(c *gin.Context) {
	var requestBody file

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fileName := requestBody.FileName

	fmt.Println("Recieved:", fileName)

	SQL_LIMBO.SQL_readOneFile(fileName)

	c.JSON(http.StatusOK, gin.H{"message": "Request [Read] successfull", "rows": ""})

}
