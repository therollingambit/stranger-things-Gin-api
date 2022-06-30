package main

import (
	"net/http"
	"os"
	"stranger-things-gin/configs"
	"stranger-things-gin/routes"

	"github.com/gin-gonic/gin"
)



func main() {
	router := gin.Default()

	// run db
	configs.ConnectDB()

	// health check
	router.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"data": "hello world",})
  })

	// routes
	routes.CharacterRoute(router)

	port := os.Getenv("PORT")
	router.Run(":" + port)
}