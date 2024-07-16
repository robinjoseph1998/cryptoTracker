package main

import (
	"cryptoTracker/route"
	"cryptoTracker/src/controller"
	"cryptoTracker/src/repository"
	utils "cryptoTracker/utils/database"

	"github.com/gin-gonic/gin"
)

func main() {
	db := utils.ConnectDb()
	if db == nil {
		return
	}
	repo := repository.NewPsqlRepository(db)
	apiKey := "843b432c-744a-4317-836f-55eef2e4c9ce"
	router := gin.Default()
	ctrl := controller.NewController(apiKey, repo)
	route.SetUpRoutes(router, ctrl)
	router.Run(":8080")

}
