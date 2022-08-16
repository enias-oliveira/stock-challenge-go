package handler

import (
	"net/http"
	"stock-challenge-go/pkg/domain"
	srvcInterface "stock-challenge-go/pkg/service/interface"

	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	accountService srvcInterface.AccountService
}

func NewAccountHandler(accountService srvcInterface.AccountService) *AccountHandler {
	return &AccountHandler{accountService: accountService}
}

func (ah *AccountHandler) Register(c *gin.Context) {
	var account domain.Account

	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	account, arErr := ah.accountService.Register(c.Request.Context(), account)

	if arErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": arErr.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"email":    account.Email,
		"password": account.Password,
	})
}
