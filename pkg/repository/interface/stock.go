package interfaces

import (
	"stock-challenge-go/pkg/domain"
)

type StockRepository interface {
	GetStock(queryParams string) (domain.Stock, error)
}
