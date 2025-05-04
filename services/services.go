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

func (s *Service) GetUserByUserId(c *gin.Context) {
	log.Print("Getting users")
	c.Header("Content-Type", "application/json")

	u := model.User{}

	u.GetUserByUserId(c, s.DB)
}

func (s *Service) GetPointsForTheUser(c *gin.Context) {
	u := model.Points{}

	u.GetPointsForUser(c, s.DB)
}

func (s *Service) GetRecognitions(c *gin.Context) {
	log.Print("Getting recognitions")
	c.Header("Content-Type", "application/json")

	r := model.Recognition{}

	r.GetRecognitions(c, s.DB)
}

func (s *Service) GetRecognitionById(c *gin.Context) {
	log.Print("Getting recognition by id")
	c.Header("Content-Type", "application/json")

	r := model.Recognition{}

	r.GetRecognitionById(c, s.DB)
}

func (s *Service) CreateRecognition(c *gin.Context) {
	log.Print("Creating recognition")
	c.Header("Content-Type", "application/json")

	r := model.Recognition{}

	r.CreateRecognition(c, s.DB)
}
func (s *Service) UpdateRecognition(c *gin.Context) {
	log.Print("Updating recognition")
	c.Header("Content-Type", "application/json")

	r := model.Recognition{}

	r.UpdateRecognition(c, s.DB)
}
func (s *Service) DeleteRecognition(c *gin.Context) {
	log.Print("Deleting recognition")
	c.Header("Content-Type", "application/json")

	r := model.Recognition{}

	r.DeleteRecognition(c, s.DB)
}
