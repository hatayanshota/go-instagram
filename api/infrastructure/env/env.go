package env

import (
	"log"

	"github.com/joho/godotenv"
)

func SetEnv() {
	func SetEnv() {
	err := godotenv.Load("../../../envfile/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
}