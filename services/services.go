package services

import (
	"backend/rewards-recognition-api/model"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	DB *mongo.Database
}

func (s *Service) GetUsers(c *gin.Context) {
	log.Print("Getting users")
	c.Header("Content-Type", "application/json")

	u := model.User{}

	u.GetUsers(c, s.DB)
}
