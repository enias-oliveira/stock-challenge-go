package interfaces

import (
	"stock-challenge-go/pkg/client"
)

type StockRepository interface {
	GetStock(queryParams string) (client.StooqStock, error)
}
