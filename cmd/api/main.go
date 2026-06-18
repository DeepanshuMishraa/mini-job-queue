package main

import (
	"fmt"
	"log"

	"github.com/DeepanshuMishraa/mini-job-queue/config"
	"github.com/DeepanshuMishraa/mini-job-queue/db"
	"github.com/DeepanshuMishraa/mini-job-queue/utils"
	"github.com/gofiber/fiber/v3"
)

func main() {
	cfg, err := config.Load()
	app := fiber.New()

	if err != nil {
		log.Fatal("Failed to load env vars")
	}

	_, err = db.ConnectDB(cfg.DATABASE_URL)
	_ = utils.Connect(cfg.REDIS_ADDRESS, cfg.REDIS_PASSWORD)

	if err != nil {
		log.Fatal("Failed to connect to the database with error: ", err)
	}

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	fmt.Println("[API-SERVER] running on ", cfg.PORT)

	log.Fatal(app.Listen(":" + cfg.PORT))
}
