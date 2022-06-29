package routes

import (
	"stranger-things-gin/controllers"

	"github.com/gin-gonic/gin"
)

func CharacterRoute(router *gin.Engine) {
	router.GET("/characters", controllers.GetAllCharacters())
	router.GET("/character/:characterId", controllers.GetCharacter())
	router.POST("/character", controllers.CreateCharacter())
	router.PUT("/character/:characterId", controllers.EditCharacter())
	router.DELETE("/character/:characterId", controllers.DeleteCharacter())
}