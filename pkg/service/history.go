package services

import (
	"stock-challenge-go/pkg/domain"

	repoInterface "stock-challenge-go/pkg/repository/interface"
	srvcInterface "stock-challenge-go/pkg/service/interface"
)

type HistoryService struct {
	repo repoInterface.HistoryRepository
}

func NewHistoryService(repo repoInterface.HistoryRepository) srvcInterface.HistoryService {
	return &HistoryService{repo}
}

func (hs *HistoryService) SaveStockQuoteRequest(sqRequest domain.StockQuoteRequest) (domain.StockQuoteRequest, error) {
	savedRequest, err := hs.repo.Save(sqRequest)

	return savedRequest, err
}
