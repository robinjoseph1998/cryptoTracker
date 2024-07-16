package repository

import (
	"cryptoTracker/src/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type psqlRepo struct {
	db *gorm.DB
}

func NewPsqlRepository(db *gorm.DB) PsqlRepository {
	return &psqlRepo{db: db}
}

/***Inserting coin data to database***/
func (p *psqlRepo) SaveCryptocurrency(crypto *models.Cryptocurrency) error {
	err := p.db.Save(crypto)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

var ErrCryptoNotFound = errors.New("cryptocurrency not found") //custom error to crypto not found

/***Searching coin data in database by name***/
func (p *psqlRepo) SearchByName(name string) (models.Cryptocurrency, error) {
	fmt.Println("name", name)
	var crypto models.Cryptocurrency
	err := p.db.Where("name=?", name).First(&crypto)
	if err.Error != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			return crypto, ErrCryptoNotFound
		}
		return crypto, err.Error
	}
	return crypto, nil
}

/***Searching coin data in database by symbol***/
func (p *psqlRepo) SearchBySymbol(symbol string) (models.Cryptocurrency, error) {
	var crypto models.Cryptocurrency
	err := p.db.Where("symbol=?", symbol).First(&crypto)
	if err.Error != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			return crypto, ErrCryptoNotFound
		}
		return crypto, err.Error
	}
	return crypto, nil
}
