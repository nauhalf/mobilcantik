package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func LoadEnvironment() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
		return err
	}

	env := os.Getenv("ENV")

	if strings.ToUpper(env) == "LOCAL" {
		env = ".local"
	} else {
		env = ".production"
	}

	err = godotenv.Load(".env" + env)
	if err != nil {
		log.Println("Error loading .env" + env + " file")
		return err
	}
	return nil
}
