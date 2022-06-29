package controllers

import (
	"context"
	"net/http"
	"stranger-things-gin/configs"
	"stranger-things-gin/models"
	"stranger-things-gin/responses"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var characterCollection *mongo.Collection = configs.GetCollection(configs.DB, "characters")
var validate = validator.New()

// POST
func CreateCharacter() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
		var character models.Character
		defer cancel()

		// validate request body
		if err := c.BindJSON(&character); err != nil {
			c.JSON(http.StatusBadRequest, responses.CharacterResponse{
				Status: http.StatusBadRequest,
				Message: "error",
				Data: map[string]interface{}{"data": err.Error()},
			})
		}

		// use validator library to validate required fields
		if validationErr := validate.Struct(&character); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.CharacterResponse{
				Status: http.StatusBadRequest,
				Message: "error",
				Data: map[string]interface{}{"data": validationErr.Error()},
			})
		}

		newCharacter := models.Character{
			Id: primitive.NewObjectID(),
			Name: character.Name,
			Nicknames: character.Nicknames,
			Born: character.Born,
		}

		result, err := characterCollection.InsertOne(ctx, newCharacter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.CharacterResponse{
				Status: http.StatusInternalServerError,
				Message: "error",
				Data: map[string]interface{}{"data": err.Error()},
			})
			return
		}

		c.JSON(http.StatusCreated, responses.CharacterResponse{
			Status: http.StatusCreated,
			Message: "success",
			Data: map[string]interface{}{"data": result},
		})
	}
}

// GET
func GetCharacter() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
		characterId := c.Param("characterId")
		var character models.Character
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(characterId)

		err := characterCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&character)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.CharacterResponse{
				Status: http.StatusInternalServerError,
				Message: "error",
				Data: map[string]interface{}{"data": err.Error()},
			})
			return
		}

		c.JSON(http.StatusOK, responses.CharacterResponse{
			Status: http.StatusOK,
			Message: "success",
			Data: map[string]interface{}{"data": character},
		})
	}
}

func GetAllCharacters() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var characters []models.Character
		defer cancel()

		results, err := characterCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.CharacterResponse{
				Status: http.StatusInternalServerError,
				Message: "error",
				Data: map[string]interface{}{"data": err.Error()},
			})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleCharacter models.Character
			if err = results.Decode(&singleCharacter); err != nil {
				c.JSON(http.StatusInternalServerError, responses.CharacterResponse{
					Status: http.StatusInternalServerError,
					Message: "error",
					Data: map[string]interface{}{"data": err.Error()},
				})
			}

			characters = append(characters, singleCharacter)
		}

		c.JSON(http.StatusOK,
			responses.CharacterResponse{
				Status: http.StatusOK,
				Message: "success",
				Data: map[string]interface{}{"data": characters},
			},
		)
	}
}

// PUT
func EditCharacter() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
		characterId := c.Param("characterId")

		var character models.Character
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(characterId)

		// validate request body
		if err := c.BindJSON(&character); err != nil {
			c.JSON(http.StatusBadRequest, responses.CharacterResponse{
				Status: http.StatusBadRequest,
				Message: "error",
				Data: map[string]interface{}{"data": err.Error()},
			})
			return
		}

		// use validator library to validate required fields
		if validationErr := validate.Struct(&character); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.CharacterResponse{
				Status: http.StatusBadRequest,
				Message: "error",
				Data: map[string]interface{}{"data": validationErr.Error()},
			})
			return
		}

		update := bson.M{
			"name": character.Name,
			"nicknames": character.Nicknames,
			"born": character.Born,
		}
		result, err := characterCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.CharacterResponse{
				Status: http.StatusInternalServerError,
				Message: "error",
				Data: map[string]interface{}{"data": err.Error()},
			})
			return
		}

		// get updated character details
		var updatedCharacter models.Character
		if result.MatchedCount == 1 {
			err := characterCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedCharacter)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.CharacterResponse{
					Status: http.StatusInternalServerError,
					Message: "error",
					Data: map[string]interface{}{"data": err.Error()},
				})
				return
			}
		}

		c.JSON(http.StatusOK, responses.CharacterResponse{
			Status: http.StatusOK,
			Message: "success",
			Data: map[string]interface{}{"data": updatedCharacter},
		})
	}
}

// DELETE
func DeleteCharacter() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		characterId := c.Param("characterId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(characterId)

		result, err := characterCollection.DeleteOne(ctx, bson.M{"id": objId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.CharacterResponse{
				Status: http.StatusInternalServerError,
				Message: "error",
				Data: map[string]interface{}{"data": err.Error()},
			})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound, responses.CharacterResponse{
				Status: http.StatusNotFound,
				Message: "error",
				Data: map[string]interface{}{"data": "Character with specified ID not found!"},
			})
			return
		}

		c.JSON(http.StatusOK, responses.CharacterResponse{
			Status: http.StatusOK,
			Message: "success",
			Data: map[string]interface{}{"data": "Character successfully deleted!"},
		})
	}
}