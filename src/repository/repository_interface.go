package repository

import "cryptoTracker/src/models"

type PsqlRepository interface {
	SaveCryptocurrency(crypto *models.Cryptocurrency) error
	SearchByName(name string) (*models.Cryptocurrency, error)
	SearchBySymbol(symbol string) (*models.Cryptocurrency, error)
}
