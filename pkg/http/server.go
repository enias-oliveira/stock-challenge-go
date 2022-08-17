package http

import (
	"stock-challenge-go/pkg/http/handler"
	"stock-challenge-go/pkg/http/middleware"

	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(accountHandler *handler.AccountHandler, stockHandler *handler.StockHandler, historyHandler *handler.HistoryHandler) *ServerHTTP {
	engine := gin.New()

	engine.Use(gin.Logger())

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	engine.POST("/register", accountHandler.Register)
	engine.POST("/login", accountHandler.Login)

	api := engine.Group("/api", middleware.AuthorizationMiddleware)

	api.GET("/profile", accountHandler.Profile)
	api.GET("/stock", stockHandler.GetStock)
	api.Use(middleware.RoleGuardMiddleware).GET("history", historyHandler.GetHistory)

	return &ServerHTTP{engine: engine}
}

func (s *ServerHTTP) Run(addr string) {
	s.engine.Run(addr)
}
