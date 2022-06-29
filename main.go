package main

import (
	"net/http"
	"stranger-things-gin/configs"
	"stranger-things-gin/routes"

	"github.com/gin-gonic/gin"
)



func main() {
	router := gin.Default()

	// health check
	router.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"data": "hello world"})
  })

	// run db
	configs.ConnectDB()

	// routes
	routes.CharacterRoute(router)

	router.Run()
}