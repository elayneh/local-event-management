// constants.go
package controller

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	adminSecret string
	secretKey   []byte
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	adminSecret = os.Getenv("HASURA_GRAPHQL_ADMIN_SECRET")
	secretKey = []byte(os.Getenv("JWT_SECRET"))

	if adminSecret == "" {
		log.Println("Warning: HASURA_GRAPHQL_ADMIN_SECRET is not set")
	}
	if len(secretKey) == 0 {
		log.Println("Warning: JWT_SECRET is not set or empty")
	}

	println("Env variables: ", secretKey, "\n", adminSecret)
}
