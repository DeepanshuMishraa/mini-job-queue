package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT         string
	DATABASE_URL string
}

func Load() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		return &Config{}, errors.New("Env Not Recognized")
	}

	DATABASE_URL := os.Getenv("DATABASE_URL")
	PORT := os.Getenv("PORT")

	if DATABASE_URL == "" || PORT == "" {
		return &Config{}, errors.New("env vars cannot be empty")
	}

	log.Println("All Env Vars Loaded")

	return &Config{
		DATABASE_URL: DATABASE_URL,
		PORT:         PORT,
	}, nil
}
