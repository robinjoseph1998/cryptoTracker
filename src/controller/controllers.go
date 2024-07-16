package controller

import (
	"cryptoTracker/src/models"
	"cryptoTracker/src/repository"
	utils "cryptoTracker/utils/coinmarketcap"
	"encoding/json"
	"log"
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

/***API to Search by name***/
func (ctrl *Controller) SearchCoinbyName(c *gin.Context) {
	var input models.NameSymbol
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Invalid input": err})
		return
	}
	if input.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "please enter the coin name"})
		return
	}
	crypto, err := ctrl.Repo.SearchByName(input.Name)
	if err != nil {
		if err == repository.ErrCryptoNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "No cryptocurrency found with this name"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"can't fetch the datas": err})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"coin data": crypto})
}

/***API to Search by Symbol***/
func (ctrl *Controller) SearchCoinbySymbol(c *gin.Context) {
	var input models.NameSymbol
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Invalid input": err})
		return
	}
	if input.Symbol == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "please enter the coin symbol"})
		return
	}
	crypto, err := ctrl.Repo.SearchBySymbol(input.Symbol)
	if err != nil {
		if err == repository.ErrCryptoNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "No cryptocurrency found with this symbol"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"can't fetch the datas": err})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"coin data": crypto})
}

/***API to Automaticaly Updating Coin Value Changes***/
func (ctrl *Controller) UpdateCoins() {
	data, err := ctrl.CMCClient.GetLatestListings()
	if err != nil {
		log.Printf("Error fetching latest listings: %v", err)
		return
	}

	var listings map[string]interface{}
	if err := json.Unmarshal([]byte(data), &listings); err != nil {
		log.Printf("Error unmarshalling data: %v", err)
		return
	}

	if listingsData, ok := listings["data"].([]interface{}); ok {
		for _, listing := range listingsData {
			listingMap := listing.(map[string]interface{})
			crypto := models.Cryptocurrency{
				Name:             listingMap["name"].(string),
				Symbol:           listingMap["symbol"].(string),
				CurrentPrice:     listingMap["quote"].(map[string]interface{})["USD"].(map[string]interface{})["price"].(float64),
				MarketCap:        listingMap["quote"].(map[string]interface{})["USD"].(map[string]interface{})["market_cap"].(float64),
				Volume24h:        listingMap["quote"].(map[string]interface{})["USD"].(map[string]interface{})["volume_24h"].(float64),
				PercentChange1h:  listingMap["quote"].(map[string]interface{})["USD"].(map[string]interface{})["percent_change_1h"].(float64),
				PercentChange24h: listingMap["quote"].(map[string]interface{})["USD"].(map[string]interface{})["percent_change_24h"].(float64),
				PercentChange7d:  listingMap["quote"].(map[string]interface{})["USD"].(map[string]interface{})["percent_change_7d"].(float64),
			}

			existingCrypto, err := ctrl.Repo.SearchBySymbol(crypto.Symbol)
			if err != nil {
				if err == repository.ErrCryptoNotFound {
					err = ctrl.Repo.SaveCryptocurrency(&crypto)
					if err != nil {
						log.Printf("Error saving new cryptocurrency: %v", err)
					} else {
						log.Printf("Saved new cryptocurrency: %s (%s)", crypto.Name, crypto.Symbol)
					}
				} else {
					log.Printf("Error finding cryptocurrency: %v", err)
				}
				continue
			}

			// Check for updates
			if existingCrypto.CurrentPrice != crypto.CurrentPrice ||
				existingCrypto.MarketCap != crypto.MarketCap ||
				existingCrypto.Volume24h != crypto.Volume24h ||
				existingCrypto.PercentChange1h != crypto.PercentChange1h ||
				existingCrypto.PercentChange24h != crypto.PercentChange24h ||
				existingCrypto.PercentChange7d != crypto.PercentChange7d {

				err = ctrl.Repo.SaveCryptocurrency(&crypto)
				if err != nil {
					log.Printf("Error updating cryptocurrency: %v", err)
				} else {
					log.Printf("Updated cryptocurrency: %s (%s)", crypto.Name, crypto.Symbol)
				}
			}
		}
		log.Printf("Cryptocurrency data updated successfully")
	} else {
		log.Printf("No data found in listings")
	}
}
