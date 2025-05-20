package jwt

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var Clave = []byte("rpIVZInJ81Be9xGMoi75")

func Generar_token(username string, id string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nombre": username,
		"id":     id,
		"iat":    time.Now().Unix(),
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(Clave)

	return tokenString, err
}

func MiddlewareJWT() gin.HandlerFunc {

	return func(c *gin.Context) {
		h := c.GetHeader("Authorization")

		if strings.TrimSpace(h) == "" {
			c.JSON(401, gin.H{"mensaje": "falta autorizacion"})
			c.Abort()
			return
		}

		tokenString := strings.Replace(h, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("error inesperado")
			}
			return Clave, nil
		})

		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		_, ok := token.Claims.(jwt.MapClaims)
		if !ok && !token.Valid {

			c.JSON(400, "error en token claims")
			c.Abort()
			return
		}

		c.Next()

	}

}
