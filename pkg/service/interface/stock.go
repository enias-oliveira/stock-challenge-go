package interfaces

import (
	domain "stock-challenge-go/pkg/domain"
)

type StockService interface {
	GetStock(quote string) (domain.Stock, error)
}
