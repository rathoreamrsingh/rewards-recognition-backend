package model

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Recognition struct {
	Id            primitive.ObjectID `json:"id" bson:"_id"`
	Subject       string             `json:"subject" bson:"subject"`
	Body          string             `json:"body" bson:"body"`
	HashTags      []string           `json:"hashTags" bson:"hashTags"`
	TaggedUserIds []int              `json:"taggedUserIds" bson:"taggedUserIds"`
	IsPrivate     bool               `json:"isPrivate" bson:"isPrivate"`
	PointsGiven   int                `json:"pointsGiven" bson:"pointsGiven"`
	CreatedAt     time.Time          `json:"createdAt" bson:"createdAt"`
	CreatedBy     int                `json:"createdBy" bson:"createdBy"`
	UpdatedAt     time.Time          `json:"updatedAt" bson:"updatedAt"`
	UpdatedBy     int                `json:"updatedBy" bson:"updatedBy"`
	IsEdited      bool               `json:"edited" bson:"edited"`
	DeletedAt     time.Time          `json:"deletedAt" bson:"deletedAt"`
	DeletedBy     int                `json:"deletedBy" bson:"deletedBy"`
	IsDeleleted   bool               `json:"deleted" bson:"deleted"`
}

/**
* Get all recognitions
* @param c *gin.Context
* @param DB *mongo.Database
* @return List of recognitions
**/
func (r *Recognition) GetRecognitions(c *gin.Context, DB *mongo.Database) {
	c.Header("Content-Type", "application/json")
	// Get all recognitions from the database
	collection := DB.Collection("recognitions")
	filter := bson.M{"deleted": false}
	cursor, err := collection.Find(c.Request.Context(), filter)
	//cursor, err := collection.Find(c.Request.Context(), bson.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(c.Request.Context())
	var recognitions []Recognition
	if err := cursor.All(c.Request.Context(), &recognitions); err != nil {
		log.Printf("Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, recognitions)
}

/**
* Get recognition by id
* @param c *gin.Context
* @param DB *mongo.Database
* @return Recognition
**/
func (r *Recognition) GetRecognitionById(c *gin.Context, DB *mongo.Database) {
	c.Header("Content-Type", "application/json")
	recognitionId, err := getObjectIdFromParam(c, "recognitionId")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	collection := DB.Collection("recognitions")
	filter := bson.M{"_id": recognitionId, "deleted": false}
	var recognition Recognition
	err = collection.FindOne(c.Request.Context(), filter).Decode(&recognition)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, recognition)
}

/**
* Create recognition
* @param c *gin.Context
* @param DB *mongo.Database
* @return Recognition
**/
func (r *Recognition) CreateRecognition(c *gin.Context, DB *mongo.Database) {
	c.Header("Content-Type", "application/json")
	var recognition Recognition
	if err := c.ShouldBindJSON(&recognition); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	collection := DB.Collection("recognitions")
	// Generate a new ID for the recognition
	recognition.Id = primitive.NewObjectID()
	// Set the created and updated timestamps
	recognition.CreatedAt = time.Now()
	recognition.UpdatedAt = time.Now()
	recognition.DeletedAt = time.Time{}
	recognition.DeletedBy = 0
	recognition.IsDeleleted = false
	result, err := collection.InsertOne(c.Request.Context(), recognition)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

/**
* Update recognition
* @param c *gin.Context
* @param DB *mongo.Database
* @return Recognition
**/
func (r *Recognition) UpdateRecognition(c *gin.Context, DB *mongo.Database) {
	c.Header("Content-Type", "application/json")

	recognitionId, err := getObjectIdFromParam(c, "recognitionId")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var recognition Recognition
	if err := c.ShouldBindJSON(&recognition); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := DB.Collection("recognitions")

	existingRecognition, err := getRecognitionById(c, collection, recognitionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if isRecognitionDeleted(existingRecognition) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "recognition is deleted"})
		return
	}

	// Update fields
	recognition.UpdatedAt = time.Now()
	recognition.UpdatedBy = 0
	recognition.IsEdited = true

	update := bson.M{"$set": recognition}
	result, err := collection.UpdateOne(c.Request.Context(), bson.M{"_id": recognitionId}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

/**
* Delete recognition
* @param c *gin.Context
* @param DB *mongo.Database
* @return Recognition
**/
func (r *Recognition) DeleteRecognition(c *gin.Context, DB *mongo.Database) {
	c.Header("Content-Type", "application/json")
	recognitionId, err := getObjectIdFromParam(c, "recognitionId")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	collection := DB.Collection("recognitions")
	filter := bson.M{"_id": recognitionId, "deleted": false}
	// Check if the recognition exists and is not deleted
	var recognition Recognition
	err = collection.FindOne(c.Request.Context(), filter).Decode(&recognition)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "recognition not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Set the deleted flag and the deleted timestamp
	update := bson.M{
		"$set": bson.M{
			"deleted":   true,
			"deletedAt": time.Now(),
			"deletedBy": 0,
			"updatedAt": time.Now(),
			"updatedBy": 0,
			"edited":    true,
		},
	}
	// Perform the update
	result, err := collection.UpdateOne(c.Request.Context(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func getObjectIdFromParam(c *gin.Context, paramName string) (primitive.ObjectID, error) {
	id := c.Param(paramName)
	if id == "" {
		return primitive.NilObjectID, fmt.Errorf("%s is required", paramName)
	}
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("%s is not in valid format", paramName)
	}
	return objectId, nil
}

func getRecognitionById(c *gin.Context, collection *mongo.Collection, id primitive.ObjectID) (Recognition, error) {
	var recognition Recognition
	err := collection.FindOne(c.Request.Context(), bson.M{"_id": id}).Decode(&recognition)
	return recognition, err
}

func isRecognitionDeleted(rec Recognition) bool {
	return rec.IsDeleleted
}
