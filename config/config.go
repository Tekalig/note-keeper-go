package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	HasuraURL         string
	HasuraAdminSecret string
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	HasuraURL = os.Getenv("HASURA_URL")
	log.Println("HASURA_URL: ", HasuraURL)
	if HasuraURL == "" {
		log.Fatal("HASURA_URL environment variable is required")
	}

	HasuraAdminSecret = os.Getenv("HASURA_ADMIN_SECRET")
	if HasuraAdminSecret == "" {
		log.Fatal("HASURA_ADMIN_SECRET environment variable is required")
	}
}
