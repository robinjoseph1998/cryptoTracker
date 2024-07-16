package repository

import (
	"cryptoTracker/src/models"

	"gorm.io/gorm"
)

type psqlRepo struct {
	db *gorm.DB
}

func NewPsqlRepository(db *gorm.DB) PsqlRepository {
	return &psqlRepo{db: db}
}

func (p *psqlRepo) SaveCryptocurrency(crypto *models.Cryptocurrency) error {
	err := p.db.Save(crypto)
	if err.Error != nil {
		return err.Error
	}
	return nil
}
