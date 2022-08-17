package services

import (
	"stock-challenge-go/pkg/domain"

	repoInterface "stock-challenge-go/pkg/repository/interface"
	srvcInterface "stock-challenge-go/pkg/service/interface"

	"gorm.io/gorm"
)

type HistoryService struct {
	repo repoInterface.HistoryRepository
	db   *gorm.DB
}

func NewHistoryService(repo repoInterface.HistoryRepository, db *gorm.DB) srvcInterface.HistoryService {
	return &HistoryService{repo, db}
}

func (hs *HistoryService) SaveStockQuoteRequest(sqRequest domain.StockQuoteRequest) (domain.StockQuoteRequest, error) {
	savedRequest, err := hs.repo.Save(sqRequest)

	return savedRequest, err
}

func (hs *HistoryService) FindAccountStockHistory(userID int) ([]domain.StockQuoteRequest, error) {
	return hs.repo.FindByUserID(userID)
}

func (hs *HistoryService) FindMostRquestedStock() ([]domain.MostRequestedStockResult, error) {
	var mostRequestedStockResults []domain.MostRequestedStockResult

	err := hs.db.Model(&domain.StockQuoteRequest{}).Select("symbol as stock, count(*) as times_requested").Group("stock").Order("times_requested desc").Limit(5).Find(&mostRequestedStockResults).Error

	return mostRequestedStockResults, err
}
