package domain

import "gorm.io/gorm"

type StockQuoteRequest struct {
	gorm.Model
	UserID int     `gorm:"type:int"`
	Name   string  `gorm:"type:varchar(255)"`
	Symbol string  `gorm:"type:varchar(25)"`
	Open   float32 `gorm:"type:float"`
	High   float32 `gorm:"type:float"`
	Low    float32 `gorm:"type:float"`
	Close  float32 `gorm:"type:float"`
}
