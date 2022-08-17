package handler

import (
	srvcInterface "stock-challenge-go/pkg/service/interface"
	"strconv"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

type StockHandler struct {
	stockService srvcInterface.StockService
}

func NewStockHandler(stockService srvcInterface.StockService) *StockHandler {
	return &StockHandler{
		stockService,
	}
}

func (h *StockHandler) GetStock(ctx *gin.Context) {
	symbol := ctx.Query("q")
	user, userExists := ctx.Get("user")

	if !userExists {
		ctx.JSON(500, gin.H{
			"message": "error",
		})
		return
	}

	claims := user.(*jwt.Token).Claims.(*AccountClaims)
	id, err := strconv.Atoi(claims.Subject)

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "error",
		})
		return
	}

	stock, err := h.stockService.GetStock(id, symbol)

	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"date":   stock.CreatedAt.UTC(),
		"name":   stock.Name,
		"symbol": stock.Symbol,
		"open":   stock.Open,
		"high":   stock.High,
		"low":    stock.Low,
		"close":  stock.Close,
	})
}
