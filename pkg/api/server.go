package http

import (
	"stock-challenge-go/pkg/domain"

	"github.com/gin-gonic/gin"

	services "stock-challenge-go/pkg/services"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP() *ServerHTTP {
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
			c.JSON(400, gin.H{
				"error": err.Error(),
			})

			return
		}

		account, raErr := services.NewAccountService().Register(c.Request.Context(), account)

		if raErr != nil {
			c.JSON(400, gin.H{
				"error": raErr.Error(),
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
