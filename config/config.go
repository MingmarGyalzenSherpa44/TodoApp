package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ConfigStruct struct {
	DB_URI string
}

var Config ConfigStruct

func init() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Config = ConfigStruct{
		DB_URI: os.Getenv("MONGODB_URI"),
	}
}
