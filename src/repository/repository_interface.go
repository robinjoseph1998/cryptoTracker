package repository

import "cryptoTracker/src/models"

type PsqlRepository interface {
	SaveCryptocurrency(crypto *models.Cryptocurrency) error
}
