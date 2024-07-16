package models

/***Coins Table Schema***/
type Cryptocurrency struct {
	ID               uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name             string  `json:"name" gorm:"type:varchar(100);not null"`
	Symbol           string  `json:"symbol" gorm:"type:varchar(100);not null"`
	CurrentPrice     float64 `json:"current_price" gorm:"type:decimal(18,8);not null"`
	MarketCap        float64 `json:"market_cap" gorm:"type:decimal(18,2);not null"`
	Volume24h        float64 `json:"volume_24h" gorm:"type:decimal(18,2);not null"`
	PercentChange1h  float64 `json:"percent_change_1h" gorm:"type:decimal(5,2);not null"`
	PercentChange24h float64 `json:"percent_change_24h" gorm:"type:decimal(5,2);not null"`
	PercentChange7d  float64 `json:"percent_change_7d" gorm:"type:decimal(5,2);not null"`
}

/***JSON Request Model***/
type NameSymbol struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}
