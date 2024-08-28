package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}

func main() {
	r := gin.Default()

	r.GET("/hello", HelloWorld)

	r.Run(":8080")
}
