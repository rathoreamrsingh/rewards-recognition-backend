package model

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Points struct {
	Id               primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	UserId           int                `json:"userId" bson:"userId"`
	GivablePoints    int                `json:"givablePoints" bson:"givablePoints"`
	RedeemablePoints int                `json:"redeemablePoints" bson:"redeemablePoints"`
}

func (p *Points) GetPointsForUser(c *gin.Context, DB *mongo.Database) {
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

	var points []Points // Change points to a slice of Points
    if err := cursor.All(c.Request.Context(), &points); err != nil { // Pass a pointer to the slice
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Handle the case where no points are found.
    if len(points) == 0 {
        c.JSON(http.StatusOK, []Points{}) //return empty array
        return
    }

    c.JSON(http.StatusOK, points[0])
}
