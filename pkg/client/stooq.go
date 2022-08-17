package client

import (
	"encoding/json"
	"fmt"
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

type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	ErrorCode  string `json:"errorCode"`
	Message    string `json:"message"`
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

	switch res.StatusCode {
	case 200:
		body, err := ioutil.ReadAll(res.Body)

		if err != nil {
			return domain.Stock{}, err
		}

		var stocks []domain.Stock
		err = gocsv.UnmarshalBytes(body, &stocks)

		return stocks[0], err

	case 400, 401, 403, 500:
		var errRes ErrorResponse
		if err := json.NewDecoder(res.Body).Decode(&errRes); err != nil {
			return domain.Stock{}, err
		}

		if errRes.StatusCode == 0 {
			errRes.StatusCode = res.StatusCode
		}
		return domain.Stock{}, &errRes

	default:
		return domain.Stock{}, fmt.Errorf("unexpected status code %d", res.StatusCode)
	}

}

func (args StooqArgs) QueryParams() url.Values {
	params := url.Values{}
	params.Add("s", args.Symbol)
	params.Add("f", args.F)
	params.Add("e", args.E)
	params.Add("h", args.H)
	return params
}

func (err *ErrorResponse) Error() string {
	if err.ErrorCode == "" {
		return fmt.Sprintf("%d API error: %s", err.StatusCode, err.Message)
	}
	return fmt.Sprintf("%d (%s) API error: %s", err.StatusCode, err.ErrorCode, err.Message)
}
