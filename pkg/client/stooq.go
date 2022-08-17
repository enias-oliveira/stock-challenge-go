package client

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"stock-challenge-go/pkg/config"
	"stock-challenge-go/pkg/domain"
	"time"

	"github.com/gocarina/gocsv"
)

var baseURL = url.URL{
	Scheme: "https",
	Host:   "stooq.com",
}

type StooqClient struct {
	httpClient *http.Client
	ApiKey     string
}

type StooqArgs struct {
	// Stock symbol
	Symbol string

	// API key
	F string

	// Response Format (json, csv, xml)
	E string

	// Include header (1 or 0)
	H string
}

func NewStooqClient(cfg config.Config) *StooqClient {
	return &StooqClient{
		httpClient: &http.Client{
			Timeout: time.Minute,
		},
		ApiKey: cfg.StooqAPIKey,
	}
}

func (cl *StooqClient) GetStock(queryParams url.Values) (domain.Stock, error) {
	endpt := baseURL.ResolveReference(&url.URL{Path: "/q/l/"})
	req, err := http.NewRequest("GET", endpt.String(), nil)

	if err != nil {
		return domain.Stock{}, err
	}

	req.URL.RawQuery = queryParams.Encode()
	res, err := cl.httpClient.Do(req)

	if err != nil {
		return domain.Stock{}, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	var stocks []domain.Stock
	err = gocsv.UnmarshalBytes(body, &stocks)

	return stocks[0], err
}

func (args StooqArgs) QueryParams() url.Values {
	params := url.Values{}
	params.Add("s", args.Symbol)
	params.Add("f", args.F)
	params.Add("e", args.E)
	params.Add("h", args.H)
	return params
}
