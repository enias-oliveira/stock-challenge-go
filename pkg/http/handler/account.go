package handler

import (
	"fmt"
	"net/http"
	"stock-challenge-go/pkg/config"
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
	cfg            config.Config
}

type AccountClaims struct {
	Role  string `json:"role"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func NewAccountHandler(accountService srvcInterface.AccountService, cfg config.Config) *AccountHandler {
	return &AccountHandler{accountService, cfg}
}

func (ah *AccountHandler) Register(ctx *gin.Context) {
	var account domain.Account

	if err := ctx.ShouldBindJSON(&account); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	account, arErr := ah.accountService.Register(account)

	if arErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": arErr.Error(),
		})

		return
	}

	ctx.JSON(200, gin.H{
		"email":    account.Email,
		"password": account.Password,
	})
}

func (ah *AccountHandler) Login(ctx *gin.Context) {
	var account domain.Account

	if bjErr := ctx.ShouldBindJSON(&account); bjErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": bjErr.Error(),
		})

		return
	}

	validAcc, vaErr := ah.accountService.ValidateAccount(account)

	if vaErr != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
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

	tokenString, tsErr := token.SignedString([]byte(ah.cfg.JWTSecret))

	if tsErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": tsErr.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token": tokenString,
	})
}

func (ah *AccountHandler) Profile(ctx *gin.Context) {
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

	ctx.JSON(200, gin.H{
		"id":    id,
		"email": claims.Email,
		"role":  claims.Role,
	})
}
