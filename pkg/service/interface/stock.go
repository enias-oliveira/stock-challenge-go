package interfaces

import (
	domain "stock-challenge-go/pkg/domain"
)

type StockService interface {
	GetStock(symbol string) (domain.StockQuoteRequest, error)
}
