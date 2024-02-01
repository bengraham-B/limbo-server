package main

import "github.com/gin-gonic/gin"

func main() {
	connectDB()
	router := gin.Default()

	router.GET("/test", test)

	router.Run("localhost:8080")

}
