package config

import (
	"FP2/models"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() models.Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}
	serverPort := os.Getenv("SERVICE_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	config := models.Config{
		ServerPort: serverPort,
		Database: models.Database{
			Host:     dbHost,
			Port:     dbPort,
			Username: dbUsername,
			Password: dbPassword,
			Name:     dbName,
		},
	}
	return config
}
