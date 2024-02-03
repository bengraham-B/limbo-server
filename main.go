package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Limbo Server")
	log := gooseLogger("TEST", 1, "Testing Logger", "main.go", nil)
	fmt.Println(log)

	//^ Creating file Struct
	fileStruct := createFileStruct()

	router := gin.Default()

	router.GET("/test", test)

	router.POST("/create", func(c *gin.Context) {
		fileStruct.DB_InsertOneFile(c)

	})

	router.Run("localhost:8001")

}
