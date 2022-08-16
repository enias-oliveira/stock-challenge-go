package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := NewServerHTTP()
	server.Run(":8080")
}

type ServerHTTP struct {
	engine *gin.Engine
}

type RegisterRequest struct {
	Email string `json:"email"`
	Role  string `json:"role"`
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
		var rq RegisterRequest

		if err := c.ShouldBindJSON(&rq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"email":    rq.Email,
			"password": "password",
		})
	})

	return &ServerHTTP{engine: engine}
}

func (s *ServerHTTP) Run(addr string) {
	s.engine.Run(addr)
}
