package handler

import (
	srvcInterface "stock-challenge-go/pkg/service/interface"
	"strconv"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

type HistoryHandler struct {
	historyService srvcInterface.HistoryService
}

type StockResponse struct {
	Symbol string  `json:"symbol"`
	Date   string  `json:"date"`
	Name   string  `json:"name"`
	Open   float32 `json:"open"`
	High   float32 `json:"high"`
	Low    float32 `json:"low"`
	Close  float32 `json:"close"`
}

func NewHistoryHandler(historyService srvcInterface.HistoryService) *HistoryHandler {
	return &HistoryHandler{historyService: historyService}
}

func (hh *HistoryHandler) GetHistory(c *gin.Context) {
	user, userExists := c.Get("user")

	if !userExists {
		c.JSON(500, gin.H{
			"message": "error",
		})
	}

	claims := user.(*jwt.Token).Claims.(*AccountClaims)

	id, scErr := strconv.Atoi(claims.Subject)

	if scErr != nil {
		c.JSON(500, gin.H{
			"message": "error",
		})
	}

	stockHistory, err := hh.historyService.FindAccountStockHistory(id)

	var stockHistoryResponse []StockResponse

	for _, stock := range stockHistory {
		stockHistoryResponse = append(stockHistoryResponse, StockResponse{
			Symbol: stock.Symbol,
			Date:   stock.CreatedAt.UTC().String(),
			Name:   stock.Name,
			Open:   stock.Open,
			High:   stock.High,
			Low:    stock.Low,
			Close:  stock.Close,
		})
	}

	if err != nil {
		c.JSON(500, gin.H{
			"message": "error",
		})
	}

	c.JSON(200, stockHistoryResponse)
}
