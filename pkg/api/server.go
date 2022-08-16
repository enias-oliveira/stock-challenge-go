package http

import (
	"net/http"
	"stock-challenge-go/pkg/domain"

	servInterface "stock-challenge-go/pkg/service/interface"

	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(accountService servInterface.AccountService) *ServerHTTP {
	engine := gin.New()

	engine.Use(gin.Logger())

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	engine.POST("/register", func(c *gin.Context) {
		var account domain.Account

		if err := c.ShouldBindJSON(&account); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		account, arErr := accountService.Register(c.Request.Context(), account)

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
	})

	return &ServerHTTP{engine: engine}
}

func (s *ServerHTTP) Run(addr string) {
	s.engine.Run(addr)
}
