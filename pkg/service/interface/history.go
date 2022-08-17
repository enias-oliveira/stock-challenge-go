package interfaces

import (
	domain "stock-challenge-go/pkg/domain"
)

type MostRequestedStockResult struct {
	Stock          string `json:"stock"`
	TimesRequested int    `json:"times_requested"`
}

type HistoryService interface {
	SaveStockQuoteRequest(stockQuoteRequest domain.StockQuoteRequest) (domain.StockQuoteRequest, error)
	FindAccountStockHistory(userID int) ([]domain.StockQuoteRequest, error)
	FindMostRquestedStock() ([]MostRequestedStockResult, error)
}
