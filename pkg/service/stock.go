package services

import (
	"io/ioutil"
	"net/http"
	"stock-challenge-go/pkg/domain"

	"github.com/gocarina/gocsv"
	srvcInterface "stock-challenge-go/pkg/service/interface"
)

type StockService struct {
}

func NewStockService() srvcInterface.StockService {
	return &StockService{}
}

func (ss *StockService) GetStock(quote string) (domain.Stock, error) {
	requestURL := "https://stooq.com/q/l/?s=" + quote + "&f=sd2t2ohlcvn&h&e=csv"

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
