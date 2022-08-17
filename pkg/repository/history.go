package repository

import (
	"stock-challenge-go/pkg/domain"

	repoInterface "stock-challenge-go/pkg/repository/interface"

	"gorm.io/gorm"
)

type historyRepository struct {
	db *gorm.DB
}

func NewHistoryRepository(db *gorm.DB) repoInterface.HistoryRepository {
	return &historyRepository{db}
}

func (hr *historyRepository) Save(sqReq domain.StockQuoteRequest) (domain.StockQuoteRequest, error) {
	err := hr.db.Create(&sqReq).Error

	return sqReq, err
}
