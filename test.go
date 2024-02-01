package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func test(c *gin.Context) {

	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Test Works"})

}
