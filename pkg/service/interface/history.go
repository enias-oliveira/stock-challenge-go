package interfaces

import domain "stock-challenge-go/pkg/domain"

type HistoryService interface {
	SaveStockQuoteRequest(stockQuoteRequest domain.StockQuoteRequest) (domain.StockQuoteRequest, error)
}
