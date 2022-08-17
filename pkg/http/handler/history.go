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

	if err != nil {
		c.JSON(500, gin.H{
			"message": "error",
		})
	}

	c.JSON(200, stockHistory)
}
