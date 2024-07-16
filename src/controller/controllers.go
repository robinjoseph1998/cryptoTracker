package controller

import (
	"cryptoTracker/src/models"
	"cryptoTracker/src/repository"
	utils "cryptoTracker/utils/coinmarketcap"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	CMCClient *utils.Client
	Repo      repository.PsqlRepository
}

func NewController(apiKey string, repo repository.PsqlRepository) *Controller {
	return &Controller{
		CMCClient: utils.NewClient(apiKey),
		Repo:      repo}
}

/***API to Check the latest listings***/
func (ctrl *Controller) LatestListings(c *gin.Context) {
	data, err := ctrl.CMCClient.GetLatestListings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error in fetching latest listings": err.Error()})
		return
	}
	c.Data(http.StatusOK, "application/json", []byte(data))
}

/***API to Save Listings***/
func (ctrl *Controller) SaveListings(c *gin.Context) {
	data, err := ctrl.CMCClient.GetLatestListings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error in fetching latest listings": err.Error()})
		return
	}
	var listings map[string]interface{}
	if err := json.Unmarshal([]byte(data), &listings); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing response: " + err.Error()})
		return
	}
	var crypto models.Cryptocurrency
	if listingsData, ok := listings["data"].([]interface{}); ok {
		for _, listing := range listingsData {
			listingMap := listing.(map[string]interface{})
			crypto = models.Cryptocurrency{
				Name:             listingMap["name"].(string),
				Symbol:           listingMap["symbol"].(string),
				CurrentPrice:     listingMap["quote"].(map[string]interface{})["USD"].(map[string]interface{})["price"].(float64),
				MarketCap:        listingMap["quote"].(map[string]interface{})["USD"].(map[string]interface{})["market_cap"].(float64),
				Volume24h:        listingMap["quote"].(map[string]interface{})["USD"].(map[string]interface{})["volume_24h"].(float64),
				PercentChange1h:  listingMap["quote"].(map[string]interface{})["USD"].(map[string]interface{})["percent_change_1h"].(float64),
				PercentChange24h: listingMap["quote"].(map[string]interface{})["USD"].(map[string]interface{})["percent_change_24h"].(float64),
				PercentChange7d:  listingMap["quote"].(map[string]interface{})["USD"].(map[string]interface{})["percent_change_7d"].(float64),
			}
		}
		if err := ctrl.Repo.SaveCryptocurrency(&crypto); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error in Saving": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"Status": "Data Saved Successfully"})
}
