package handler

import (
	"net/http"
	"stock-challenge-go/pkg/domain"
	srvcInterface "stock-challenge-go/pkg/service/interface"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
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

func (ah *AccountHandler) Login(c *gin.Context) {
	var account domain.Account

	if bjErr := c.ShouldBindJSON(&account); bjErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": bjErr.Error(),
		})

		return
	}

	validAcc, vaErr := ah.accountService.ValidateAccount(c.Request.Context(), account)

	if vaErr != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims := jwt.MapClaims{
		"email": validAcc.Email,
		"role":  validAcc.Role,
		"id":    validAcc.ID,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, tsErr := token.SignedString([]byte("secret"))

	if tsErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": tsErr.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": tokenString,
	})
}
