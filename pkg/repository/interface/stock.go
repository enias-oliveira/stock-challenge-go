package interfaces

import (
	"stock-challenge-go/pkg/domain"
)

type StockRepository interface {
	GetStock(symbol string) (domain.Stock, error)
}
