package main

import (
	"cryptoTracker/route"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	route.SetUpRoutes(router)
	router.Run(":8080")

}
