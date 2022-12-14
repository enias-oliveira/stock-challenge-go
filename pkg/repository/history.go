package repository

import (
	"stock-challenge-go/pkg/domain"

	repoInterface "stock-challenge-go/pkg/repository/interface"
	srvcInterface "stock-challenge-go/pkg/service/interface"

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

func (hr *historyRepository) FindByUserID(userID int) ([]domain.StockQuoteRequest, error) {
	var sqReq []domain.StockQuoteRequest
	err := hr.db.Where(&domain.StockQuoteRequest{UserID: userID}, "UserID").Find(&sqReq).Error

	return sqReq, err
}

func (hr *historyRepository) GetMostRequestedStocks() ([]srvcInterface.MostRequestedStockResult, error) {
	var mostRequestedStockResults []srvcInterface.MostRequestedStockResult
	err := hr.db.Model(&domain.StockQuoteRequest{}).Select("symbol as stock, count(*) as times_requested").Group("stock").Order("times_requested desc").Limit(5).Find(&mostRequestedStockResults).Error

	return mostRequestedStockResults, err
}
