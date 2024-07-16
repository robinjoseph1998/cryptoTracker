package main

import (
	"cryptoTracker/route"
	"cryptoTracker/src/controller"
	utils "cryptoTracker/utils/database"

	"github.com/gin-gonic/gin"
)

func main() {
	db := utils.ConnectDb()
	if db == nil {
		return
	}
	router := gin.Default()
	ctrl := controller.NewController("YOUR_API_KEY_HERE")
	route.SetUpRoutes(router, ctrl)
	router.Run(":8080")

}
