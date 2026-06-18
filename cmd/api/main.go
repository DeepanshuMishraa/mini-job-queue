package main

import (
	"github.com/DeepanshuMishraa/mini-job-queue/config"
	"github.com/DeepanshuMishraa/mini-job-queue/db"
	"github.com/DeepanshuMishraa/mini-job-queue/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg, err := config.Load()
	router := gin.Default()
	router.SetTrustedProxies(nil)

	if err != nil {
		log.Fatal("Failed to load env vars")
	}

	_, err = db.ConnectDB(cfg.DATABASE_URL)
	if err != nil {
		log.Fatal("Failed to connect to the database with error: ", err)
	}

	_, err = utils.Connect(cfg.REDIS_URL)

	if err != nil {
		log.Fatal("Failed to connect to redis with error: ", err)
	}

	log.Println("[API] SERVER RUNNING ON PORT: ", cfg.PORT)

	router.Run(":" + cfg.PORT)
}
