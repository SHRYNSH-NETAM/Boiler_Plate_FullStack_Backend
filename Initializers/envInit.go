package initializers

import (
	"log"

	"github.com/joho/godotenv"
)
func Initializers() {
  if err := godotenv.Load(); err != nil {
    log.Fatal("Error loading .env file")
  }
}