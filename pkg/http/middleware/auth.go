package middleware

import (
	"fmt"
	"net/http"
	"stock-challenge-go/pkg/http/handler"
	"strings"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

func AuthorizationMiddleware(c *gin.Context) {
	s := c.Request.Header.Get("Authorization")

	accessToken := strings.TrimPrefix(s, "Bearer ")

	parsedToken, err := validateToken(accessToken)

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set("user", parsedToken)
}

func validateToken(accessToken string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(accessToken, &handler.AccountClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte("secret"), nil
	})

	return token, err
}
