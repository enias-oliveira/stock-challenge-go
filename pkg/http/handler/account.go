package handler

import (
	"fmt"
	"net/http"
	"stock-challenge-go/pkg/domain"
	srvcInterface "stock-challenge-go/pkg/service/interface"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

// TODO: add payload validation

type AccountHandler struct {
	accountService srvcInterface.AccountService
}

type AccountClaims struct {
	Role  string `json:"role"`
	Email string `json:"email"`
	jwt.RegisteredClaims
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

	account, arErr := ah.accountService.Register(account)

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

	claims := AccountClaims{
		Role:  validAcc.Role,
		Email: validAcc.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			Subject:   fmt.Sprintf("%d", validAcc.ID),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// TODO: move secret to env variable
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

func (ah *AccountHandler) Profile(c *gin.Context) {
	user, userExists := c.Get("user")

	if !userExists {
		c.JSON(500, gin.H{
			"message": "error",
		})
		return
	}

	claims := user.(*jwt.Token).Claims.(*AccountClaims)

	id, err := strconv.Atoi(claims.Subject)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "error",
		})
		return
	}

	c.JSON(200, gin.H{
		"id":    id,
		"email": claims.Email,
		"role":  claims.Role,
	})
}
