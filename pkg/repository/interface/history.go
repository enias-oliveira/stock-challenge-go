package interfaces

import (
	"stock-challenge-go/pkg/domain"
	srvcInterface "stock-challenge-go/pkg/service/interface"
)

type HistoryRepository interface {
	Save(domain.StockQuoteRequest) (domain.StockQuoteRequest, error)
	FindByUserID(int) ([]domain.StockQuoteRequest, error)
	GetMostRequestedStocks() ([]srvcInterface.MostRequestedStockResult, error)
}
