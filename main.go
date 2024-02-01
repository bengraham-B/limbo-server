package main

import "github.com/gin-gonic/gin"

func main() {
	// connectDB()
	createTable()
	insertData("Test.txt", "98, 345, 21, 43, 45, 678, 775, 34, 224, 556, 7,8 ,8")
	router := gin.Default()

	router.GET("/test", test)

	router.Run("localhost:8080")

}
