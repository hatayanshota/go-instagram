package public

import (
	"log"

	"github.com/joho/godotenv"
)

func SetEnv() {
	err := godotenv.Load("../envfiles/develop.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
