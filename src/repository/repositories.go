package repository

import (
	"cryptoTracker/src/models"
	"errors"

	"gorm.io/gorm"
)

type psqlRepo struct {
	db *gorm.DB
}

func NewPsqlRepository(db *gorm.DB) PsqlRepository {
	return &psqlRepo{db: db}
}

/***Inserting coin datas to database***/
func (p *psqlRepo) SaveCryptocurrency(crypto *models.Cryptocurrency) error {
	var existingCrypto models.Cryptocurrency
	if err := p.db.Where("symbol = ?", crypto.Symbol).First(&existingCrypto).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return p.db.Create(crypto).Error
		}
		return err
	}
	existingCrypto.Name = crypto.Name
	existingCrypto.CurrentPrice = crypto.CurrentPrice
	existingCrypto.MarketCap = crypto.MarketCap
	existingCrypto.Volume24h = crypto.Volume24h
	existingCrypto.PercentChange1h = crypto.PercentChange1h
	existingCrypto.PercentChange24h = crypto.PercentChange24h
	existingCrypto.PercentChange7d = crypto.PercentChange7d

	return p.db.Save(&existingCrypto).Error
}

var ErrCryptoNotFound = errors.New("cryptocurrency not found") //custom error to crypto not found

/***Searching coin datas in database by name***/
func (p *psqlRepo) SearchByName(name string) (*models.Cryptocurrency, error) {
	var crypto models.Cryptocurrency
	err := p.db.Where("name=?", name).First(&crypto)
	if err.Error != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			return &crypto, ErrCryptoNotFound
		}
		return &crypto, err.Error
	}
	return &crypto, nil
}

/***Searching coin datas in database by symbol***/
func (p *psqlRepo) SearchBySymbol(symbol string) (*models.Cryptocurrency, error) {
	var crypto models.Cryptocurrency
	err := p.db.Where("symbol=?", symbol).First(&crypto)
	if err.Error != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			return &crypto, ErrCryptoNotFound
		}
		return &crypto, err.Error
	}
	return &crypto, nil
}
