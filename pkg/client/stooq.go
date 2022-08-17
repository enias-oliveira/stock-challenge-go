package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"stock-challenge-go/pkg/config"
	"time"

	"github.com/gocarina/gocsv"
)

type StooqStock struct {
	Symbol string  `csv:"Symbol"`
	Date   string  `csv:"Date"`
	Time   string  `csv:"Time"`
	Open   float32 `csv:"Open"`
	High   float32 `csv:"High"`
	Low    float32 `csv:"Low"`
	Close  float32 `csv:"Close"`
	Volume int     `csv:"Volume"`
	Name   string  `csv:"Name"`
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

var baseURL = url.URL{
	Scheme: "https",
	Host:   "stooq.com",
}

func NewStooqClient(cfg config.Config) *StooqClient {
	return &StooqClient{
		httpClient: &http.Client{
			Timeout: time.Minute,
		},
		ApiKey: cfg.StooqAPIKey,
	}
}

func (cl *StooqClient) GetStock(queryParams url.Values) (StooqStock, error) {
	endpt := baseURL.ResolveReference(&url.URL{Path: "/q/l/"})
	req, err := http.NewRequest("GET", endpt.String(), nil)

	if err != nil {
		return StooqStock{}, err
	}

	req.URL.RawQuery = queryParams.Encode()
	res, err := cl.httpClient.Do(req)

	if err != nil {
		return StooqStock{}, err
	}

	defer res.Body.Close()

	switch res.StatusCode {
	case 200:
		body, err := ioutil.ReadAll(res.Body)

		if err != nil {
			return StooqStock{}, err
		}

		var stocks []StooqStock
		err = gocsv.UnmarshalBytes(body, &stocks)

		return stocks[0], err

	case 400, 401, 403, 500:
		var errRes ErrorResponse
		if err := json.NewDecoder(res.Body).Decode(&errRes); err != nil {
			return StooqStock{}, err
		}

		if errRes.StatusCode == 0 {
			errRes.StatusCode = res.StatusCode
		}
		return StooqStock{}, &errRes

	default:
		return StooqStock{}, fmt.Errorf("unexpected status code %d", res.StatusCode)
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
