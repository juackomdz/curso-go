package jwt

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var Clave = []byte("rpIVZInJ81Be9xGMoi75")

func Generar_token(correo string, username string, id string) (string, error) {

	//echojwt.JWT(Clave)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nombre": username,
		"email":  correo,
		"id":     id,
		"iat":    time.Now().Unix(),
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(Clave)

	return tokenString, err
}

func MiddlewareJWT() echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			h := c.Request().Header.Get("Authorization")

			if len(h) == 0 {
				return c.JSON(400, echo.Map{"mensaje": "falta autorizacion"})
			}

			tokenString := strings.Replace(h, "Bearer", "", 1)

			to := strings.TrimSpace(tokenString)
			token, err := jwt.Parse(to, func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("error inesperado")
				}
				return Clave, nil
			})

			if err != nil {
				return c.JSON(400, echo.Map{"error": err.Error()})
				//log.Println("error en token")
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok && !token.Valid {

				return c.JSON(400, "error en token claims")
			}

			c.Set("claims", claims)

			return next(c)
		}
	}
}
