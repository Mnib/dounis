package auth

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Id int `json:"id"`
	jwt.RegisteredClaims
}

func getSecret() []byte {
	secret := os.Getenv("JWT_SECRET")

	if secret == "" {
		secret = "jwt_secret"
	}

	return []byte(secret)
}

func getClaims(token_str string) (*Claims, error) {
	fmt.Println(token_str)

	token, err := jwt.ParseWithClaims(token_str, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return getSecret(), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("Invalid token")
}

func getTokenForUser(claims Claims) (string, error) {
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(getSecret())
}

func JWTMiddleware(c *gin.Context) {
	auth_header := c.Request.Header.Get("Authorization")

	if auth_header == "" {
		c.Status(http.StatusUnauthorized)
		c.Abort()
		return
	}

	parts := strings.SplitN(auth_header, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.Status(http.StatusUnauthorized)
		c.Abort()
		return
	}

	claims, err := getClaims(parts[1])
	if err != nil {
		c.Status(http.StatusUnauthorized)
		c.Abort()
		return
	}

	c.Set("claims", claims)
	c.Next()
}
