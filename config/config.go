package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GoogleClientID string
	JWTSecret      string
}

func Load() *Config {
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found")
    }
    log.Print("google client:", os.Getenv("GOOGLE_CLIENT_ID"), "\njwt secret:", os.Getenv("JWT_SECRET"))
	return &Config{
		GoogleClientID: os.Getenv("GOOGLE_CLIENT_ID"),
		JWTSecret:      os.Getenv("JWT_SECRET"),
	}
}