package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/test", test)

	router.POST("/create", func(c *gin.Context) {
		saveToDB(c)

	})

	router.POST("/read", func(c *gin.Context) {
		DB_readOneFile(c)
	})

	router.Run("localhost:8080")

}
