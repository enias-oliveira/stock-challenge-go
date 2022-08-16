package http

import (
	"crypto/sha256"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sethvargo/go-password/password"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServerHTTP struct {
	engine *gin.Engine
}

type User struct {
	gorm.Model
	Email        string `gorm:"type:varchar(100);unique;uniqueIndex;"`
	PasswordHash []byte `gorm:"type:binary(32);"`
	Role         string `gorm:"type:enum('admin','user');default:'user'"`
}

type RegisterRequest struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}

func NewServerHTTP() *ServerHTTP {
	dsn := "root@/stock-challenge"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&User{})

	if err != nil {
		panic(err)
	}

	engine := gin.New()

	engine.Use(gin.Logger())

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	engine.POST("/register", func(c *gin.Context) {
		var rq RegisterRequest

		if bErr := c.ShouldBindJSON(&rq); bErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": bErr.Error()})
			return
		}

		newPassword, gpErr := password.Generate(32, 4, 4, false, false)
		hashedPassword := sha256.Sum256([]byte(newPassword))

		if gpErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gpErr.Error()})
			return
		}

		cuErr := db.Create(&User{
			Model:        gorm.Model{},
			Email:        rq.Email,
			PasswordHash: hashedPassword[:],
			Role:         rq.Role,
		}).Error

		if cuErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": cuErr.Error()})
			return
		}

		c.JSON(200, gin.H{
			"email":    rq.Email,
			"password": newPassword,
		})
	})

	return &ServerHTTP{engine: engine}
}

func (s *ServerHTTP) Run(addr string) {
	s.engine.Run(addr)
}
