package services

import (
	"stock-challenge-go/pkg/domain"
	srvcInterface "stock-challenge-go/pkg/service/interface"

	"gorm.io/gorm"
)

type HistoryService struct {
	db *gorm.DB
}

func NewHistoryService(db *gorm.DB) srvcInterface.HistoryService {
	return &HistoryService{db}
}

func (hs *HistoryService) SaveStockQuoteRequest(sqRequest domain.StockQuoteRequest) (domain.StockQuoteRequest, error) {
	err := hs.db.Create(&sqRequest).Error

	return sqRequest, err
}
