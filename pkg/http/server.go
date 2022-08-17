package http

import (
	"io/ioutil"
	"net/http"
	"stock-challenge-go/pkg/http/handler"
	"stock-challenge-go/pkg/http/middleware"

	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(accountHandler *handler.AccountHandler) *ServerHTTP {
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

	api.GET("/stock", func(c *gin.Context) {
		quote := c.Query("q")

		requestURL := "https://stooq.com/q/l/?s=" + quote + "&f=sd2t2ohlcvn&h&e=csv"

		res, err := http.Get(requestURL)

		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		body, err := ioutil.ReadAll(res.Body)

		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"quote": quote,
			"data":  string(body),
		})
	})

	return &ServerHTTP{engine: engine}
}

func (s *ServerHTTP) Run(addr string) {
	s.engine.Run(addr)
}
