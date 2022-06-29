package main

import (
	"stranger-things-gin/configs"
	"stranger-things-gin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// run db
	configs.ConnectDB()

	// routes
	routes.CharacterRoute(router)

	router.Run("localhost:6000")
}