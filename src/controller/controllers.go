package controller

import (
	utils "cryptoTracker/utils/coinmarketcap"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	CMCClient *utils.Client
}

func NewController(apiKey string) *Controller {
	return &Controller{CMCClient: utils.NewClient(apiKey)}
}

func (ctrl *Controller) LatestListings(c *gin.Context) {
	data, err := ctrl.CMCClient.GetLatestListings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error in fetching latest listings": err.Error()})
		return
	}
	c.Data(http.StatusOK, "application/json", []byte(data))
}
