package jwt

import (
	"clase_5/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var Clave = []byte("rpIVZInJ81Be9xGMoi75")

func GenerarToken(usuario models.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nombre": usuario.Username,
		"email":  usuario.Email,
		"id":     usuario.Id,
		"iat":    time.Now().Unix(),
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(Clave)

	return tokenString, err
}
