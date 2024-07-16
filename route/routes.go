package route

import (
	"cryptoTracker/src/controller"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(router *gin.Engine, ctrl *controller.Controller) {

	router.GET("/api/v1/cryptocurrencies", ctrl.LatestListings)

}
