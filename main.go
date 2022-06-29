package main

import (
	"os"
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

	// port := os.Getenv("PORT")
	apiUrl := os.Getenv("APIURL")
	router.Run(apiUrl)
}