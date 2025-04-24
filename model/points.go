package model

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Points struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	UserId           int                `json:"userId" bson:"userId"`
	GivablePoints    int                `json:"givablePoints" bson:"givablePoints"`
	RedeemablePoints int                `json:"redeemablePoints" bson:"redeemablePoints"`
}

func (p *Points) GetPointsForUser(c *gin.Context, DB *mongo.Database) {
	log.Println("Getting points")
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

	collection := DB.Collection("points")

	// Query points where userId matches
	filter := bson.M{"userId": userIdInt}

	cursor, err := collection.Find(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer cursor.Close(c.Request.Context())

	var points []Points
	if err := cursor.All(c.Request.Context(), &points); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, points)
}
