package services

import (
	"stock-challenge-go/pkg/domain"

	repoInterface "stock-challenge-go/pkg/repository/interface"
	srvcInterface "stock-challenge-go/pkg/service/interface"
)

type StockService struct {
	stockRepo repoInterface.StockRepository
}

func NewStockService(repo repoInterface.StockRepository) srvcInterface.StockService {
	return &StockService{
		stockRepo: repo,
	}
}

func (ss *StockService) GetStock(symbol string) (domain.Stock, error) {
	return ss.stockRepo.GetStock(symbol)
}
