package database

import (
	"backend/rewards-recognition-api/database/config"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	DB *mongo.Database
}

func (d *Database) Initialize() {
	//cfg, err := config.GetConfig()
	cfg, err := config.GetConfig()

	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
		return
	}

	dbURI := fmt.Sprintf("%s://%s:%s@%s/?retryWrites=true&w=majority",
		cfg.DB.Protocol,
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Host)

	clientOptions := options.Client().ApplyURI(dbURI)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("Successfully connected to MongoDB")
	d.DB = client.Database(cfg.DB.Appname)
}
