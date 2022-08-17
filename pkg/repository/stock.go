package repository

import (
	"io/ioutil"
	"net/http"
	"stock-challenge-go/pkg/domain"
	repoInterface "stock-challenge-go/pkg/repository/interface"

	"github.com/gocarina/gocsv"
)

type stockRepository struct {
}

func NewStockRepository() repoInterface.StockRepository {
	return &stockRepository{}
}

func (s *stockRepository) GetStock(symbol string) (domain.Stock, error) {
	requestURL := "https://stooq.com/q/l/?s=" + symbol + "&f=sd2t2ohlcvn&h&e=csv"

	res, err := http.Get(requestURL)

	if err != nil {
		return domain.Stock{}, err
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return domain.Stock{}, err
	}

	var stocks []domain.Stock

	err = gocsv.UnmarshalBytes(body, &stocks)

	return stocks[0], err
}
