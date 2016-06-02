package main

import (
	"log"
	"os"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})

	router.Run(":" + port)
}
