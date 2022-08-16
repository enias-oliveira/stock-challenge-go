package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	server := NewServerHTTP()
	server.Run(":8080")
}

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

	return &ServerHTTP{engine: engine}
}

func (s *ServerHTTP) Run(addr string) {
	s.engine.Run(addr)
}
