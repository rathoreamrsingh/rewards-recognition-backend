package model

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID           int    `json:"_id" bson:"_id"`
	FirstName    string `json:"first_name" bson:"first_name"`
	LastName     string `json:"last_name" bson:"last_name"`
	EmailAddress string `json:"email_address" bson:"email_address"`
	AvatarURL    string `json:"avatar_url" bson:"avatar_url"`
}

func (u *User) GetUsers(c *gin.Context, DB *mongo.Database) {
	log.Print("Getting users")
	c.Header("Content-Type", "application/json")

	collection := DB.Collection("users")

	cursor, err := collection.Find(c.Request.Context(), bson.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer cursor.Close(c.Request.Context())

	var users []User
	if err := cursor.All(c.Request.Context(), &users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}
