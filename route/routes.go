package route

import (
	"cryptoTracker/src/controller"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(router *gin.Engine, ctrl *controller.Controller) {

	router.GET("/latestlistings", ctrl.LatestListings)
	router.POST("/savelistings", ctrl.SaveListings)
	router.GET("/searchbyname", ctrl.SearchCoinbyName)
	router.GET("/searchbysymbol", ctrl.SearchCoinbySymbol)

}
