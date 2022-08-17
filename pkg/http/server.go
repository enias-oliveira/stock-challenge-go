package http

import (
	"github.com/gin-gonic/gin"

	_ "stock-challenge-go/cmd/api/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"stock-challenge-go/pkg/http/handler"
	"stock-challenge-go/pkg/http/middleware"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(accountHandler *handler.AccountHandler, stockHandler *handler.StockHandler, historyHandler *handler.HistoryHandler) *ServerHTTP {
	engine := gin.New()

	engine.Use(gin.Logger())

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	engine.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	engine.POST("/register", accountHandler.Register)
	engine.POST("/login", accountHandler.Login)

	api := engine.Group("/api", middleware.AuthorizationMiddleware)

	api.GET("/profile", accountHandler.Profile)
	api.GET("/stock", stockHandler.GetStock)
	api.GET("/history", historyHandler.GetHistory)

	api.Use(middleware.RoleGuardMiddleware).GET("/stat", historyHandler.GetStat)

	return &ServerHTTP{engine: engine}
}

func (s *ServerHTTP) Run(addr string) {
	s.engine.Run(addr)
}
