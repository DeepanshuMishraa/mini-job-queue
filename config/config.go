package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT           string
	DATABASE_URL   string
	REDIS_ADDRESS  string
	REDIS_PASSWORD string
}

func Load() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		return &Config{}, errors.New("Env Not Recognized")
	}

	DATABASE_URL := os.Getenv("DATABASE_URL")
	PORT := os.Getenv("PORT")
	REDIS_ADDR := os.Getenv("REDIS_ADDRESS")
	REDIS_PASSWORD := os.Getenv("REDIS_PASSWORD")

	if DATABASE_URL == "" || PORT == "" || REDIS_ADDR == "" || REDIS_PASSWORD == "" {
		return &Config{}, errors.New("env vars cannot be empty")
	}

	log.Println("All Env Vars Loaded")

	return &Config{
		DATABASE_URL:   DATABASE_URL,
		PORT:           PORT,
		REDIS_ADDRESS:  REDIS_ADDR,
		REDIS_PASSWORD: REDIS_PASSWORD,
	}, nil
}
