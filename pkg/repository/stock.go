package repository

import (
	"stock-challenge-go/pkg/client"
	"stock-challenge-go/pkg/domain"
	repoInterface "stock-challenge-go/pkg/repository/interface"
)

type stockRepository struct {
	stooqClient *client.StooqClient
}

func NewStockRepository(stooqClient *client.StooqClient) repoInterface.StockRepository {
	return &stockRepository{
		stooqClient: stooqClient,
	}
}

func (sr *stockRepository) GetStock(symbol string) (domain.Stock, error) {
	return sr.stooqClient.GetStock(client.StooqArgs{
		Symbol: symbol,
		// Challenge Defaults
		F: sr.stooqClient.ApiKey,
		E: "csv",
		H: "1",
	}.QueryParams())
}
