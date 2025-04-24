package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Protocol string
	Username string
	Password string
	Host     string //For now give both url and port
	Appname  string
}

type Config struct {
	Port   int
	AppEnv string
	DB     *DBConfig
}

// TODO: check why this is not loading .env in debug mode for visual studio
func GetConfig() (*Config, error) {
	err := godotenv.Load("./../.env")
	if err != nil {
		// Handle error if .env file doesn't load, but don't necessarily fatal exit
		// It's common to have environment variables set directly as well
		log.Println("Error loading .env file")
	}

	portStr := os.Getenv("PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		port = 8080 // Default port if not set or invalid
	}

	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "local" // Default environment
	}

	protocol := os.Getenv("PROTOCOL")
	if protocol == "" {
		protocol = "mongodb" // Default protocol
	}

	username := os.Getenv("USERNAME")
	if username == "" {
		username = "mongo" // Default username
	}

	password := os.Getenv("PASSWORD")
	if password == "" {
		password = "mongo" // Default password
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost:27017" // Default host
	}

	appName := os.Getenv("DB")
	if appName == "" {
		appName = "rewardsAndRecognition" // Default app name  
	}

	return &Config{
		Port:   port,
		AppEnv: appEnv,
		DB: &DBConfig{
			Protocol: protocol,
			Username: username,
			Password: password,
			Host:     host,
			Appname:  appName,
		},
	}, nil
}
