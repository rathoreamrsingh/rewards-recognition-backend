package handler

import (
	"backend/rewards-recognition-api/services"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Routers struct {
	Router *gin.Engine
}

func (r *Routers) Initialize(addr string, database *mongo.Database) {
	r.Router = gin.Default()
	r.initializeRoutes(database)
	r.Run(addr)
}

func (r *Routers) initializeRoutes(database *mongo.Database) {
	service := services.Service{DB: database}
	/**
	* Applying versiong to the API
	**/
	apiV1 := r.Router.Group("/api/v1")

	// Apis for users
	apiV1.GET("/users", service.GetUsers)
	apiV1.GET("/user/:userId", service.GetUserByUserId)

	// Apis for points
	apiV1.GET("/points/:userId", service.GetPointsForTheUser)

	// Apis for recognitions
	apiV1.GET("/recognitions", service.GetRecognitions)
	apiV1.GET("/recognition/:recognitionId", service.GetRecognitionById)
	apiV1.POST("/recognition", service.CreateRecognition)
	apiV1.PUT("/recognition/:recognitionId", service.UpdateRecognition)
	apiV1.DELETE("/recognition/:recognitionId", service.DeleteRecognition)
}

func (r *Routers) Run(addr string) {
	log.Fatal(r.Router.Run(addr))
}
