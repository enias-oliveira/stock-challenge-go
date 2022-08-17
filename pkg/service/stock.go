package services

import (
	"stock-challenge-go/pkg/domain"

	repoInterface "stock-challenge-go/pkg/repository/interface"
	srvcInterface "stock-challenge-go/pkg/service/interface"
)

type StockService struct {
	stockRepo      repoInterface.StockRepository
	HistoryService srvcInterface.HistoryService
}

func NewStockService(repo repoInterface.StockRepository, hsService srvcInterface.HistoryService) srvcInterface.StockService {
	return &StockService{
		stockRepo:      repo,
		HistoryService: hsService,
	}
}

func (ss *StockService) GetStock(symbol string) (domain.StockQuoteRequest, error) {
	stock, err := ss.stockRepo.GetStock(symbol)

	if err != nil {
		return domain.StockQuoteRequest{}, err
	}

	sqReq, err := ss.HistoryService.SaveStockQuoteRequest(domain.StockQuoteRequest{
		UserID: 1,
		Name:   stock.Name,
		Symbol: stock.Symbol,
		Open:   stock.Open,
		Close:  stock.Close,
		High:   stock.High,
		Low:    stock.Low,
	})

	if err != nil {
		return domain.StockQuoteRequest{}, err
	}

	return sqReq, err
}
