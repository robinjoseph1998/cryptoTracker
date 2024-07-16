package main

import (
	"cryptoTracker/route"
	"cryptoTracker/src/controller"
	"cryptoTracker/src/repository"
	utils "cryptoTracker/utils/database"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

func main() {
	db := utils.ConnectDb()
	if db == nil {
		return
	}
	repo := repository.NewPsqlRepository(db)

	//APIkey of Coin Market Cap Price Ticker API
	apiKey := "843b432c-744a-4317-836f-55eef2e4c9ce"

	router := gin.Default()
	ctrl := controller.NewController(apiKey, repo)
	route.SetUpRoutes(router, ctrl)

	go func() {
		if err := router.Run(":8080"); err != nil {
			log.Fatalf("Failed to start Gin server: %v", err)
		}
	}()
	c := cron.New()
	err := c.AddFunc("*/5 * * * *", func() {
		// Function to fetch and update cryptocurrency data
		ctrl.UpdateCoins()

	})
	if err != nil {
		log.Fatalf("Error scheduling cron job: %v", err)
	}
	c.Start()
	select {}
}
