package interfaces

import "stock-challenge-go/pkg/domain"

type HistoryRepository interface {
	Save(domain.StockQuoteRequest) (domain.StockQuoteRequest, error)
	FindByUserID(int) ([]domain.StockQuoteRequest, error)
	GetMostRequestedStocks() ([]domain.MostRequestedStockResult, error)
}
