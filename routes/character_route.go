package routes

import (
	"stranger-things-gin/controllers"

	"github.com/gin-gonic/gin"
)

func CharacterRoute(router *gin.Engine) {
	r := router.Group("/api")
	r.GET("/characters", controllers.GetAllCharacters())
	r.GET("/character/:characterId", controllers.GetCharacter())
	r.POST("/character", controllers.CreateCharacter())
	r.PUT("/character/:characterId", controllers.EditCharacter())
	r.DELETE("/character/:characterId", controllers.DeleteCharacter())
}