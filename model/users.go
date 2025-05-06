package model

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Id           int    `json:"id" bson:"_id"`
	FirstName    string `json:"firstName" bson:"firstName"`
	LastName     string `json:"lastName" bson:"lastName"`
	EmailAddress string `json:"emailAddress" bson:"emailAddress"`
	AvatarURL    string `json:"avatarUrl" bson:"avatarUrl"`
}

func (u *User) GetUsers(c *gin.Context, DB *mongo.Database) {
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
	log.Printf("Users: %+v", users)
	c.JSON(http.StatusOK, users)
}

func (u *User) GetUserByUserId(c *gin.Context, DB *mongo.Database) {
	c.Header("Content-Type", "application/json")

	userID := c.Param("userId")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userId is required"})
		return
	}

	userIdInt, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userId is not in right format"})
		return
	}

	collection := DB.Collection("users")

	// Query points where userId matches
	filter := bson.M{"_id": userIdInt}

	cursor, err := collection.Find(c.Request.Context(), filter)
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

	// Handle the case where no points are found.
	if len(users) == 0 {
		c.JSON(http.StatusOK, []User{}) //return empty array
		return
	}

	c.JSON(http.StatusOK, users)
}
