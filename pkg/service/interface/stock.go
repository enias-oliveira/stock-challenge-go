package interfaces

import (
	domain "stock-challenge-go/pkg/domain"
)

type StockService interface {
	GetStock(userId int, symbol string) (domain.StockQuoteRequest, error)
}
