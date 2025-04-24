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
	// Define your routes here using a.Router
	service := services.Service{DB: database}
	r.Router.GET("/users", service.GetUsers)
	r.Router.GET("/points/:userId", service.GetPointsForTheUser)
	r.Router.GET("/user/:userId", service.GetUserByUserId)
	// ... other routes
}

func (r *Routers) Run(addr string) {
	log.Fatal(r.Router.Run(addr)) // Use Gin's Run method
}
