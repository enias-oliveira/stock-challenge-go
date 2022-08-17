package handler

import (
	"github.com/gin-gonic/gin"
	srvcInterface "stock-challenge-go/pkg/service/interface"
)

type StockHandler struct {
	stockService srvcInterface.StockService
}

func NewStockHandler(stockService srvcInterface.StockService) *StockHandler {
	return &StockHandler{
		stockService,
	}
}

func (h *StockHandler) GetStock(c *gin.Context) {
	symbol := c.Query("q")

	stock, err := h.stockService.GetStock(symbol)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"date":   stock.CreatedAt.UTC(),
		"name":   stock.Name,
		"symbol": stock.Symbol,
		"open":   stock.Open,
		"high":   stock.High,
		"low":    stock.Low,
		"close":  stock.Close,
	})
}
