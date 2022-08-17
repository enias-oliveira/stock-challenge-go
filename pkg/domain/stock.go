package domain

import "gorm.io/gorm"

type Stock struct {
	Symbol string  `csv:"Symbol"`
	Date   string  `csv:"Date"`
	Time   string  `csv:"Time"`
	Open   float32 `csv:"Open"`
	High   float32 `csv:"High"`
	Low    float32 `csv:"Low"`
	Close  float32 `csv:"Close"`
	Volume int     `csv:"Volume"`
	Name   string  `csv:"Name"`
}

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
