package middleware

import (
	"clase_5/database"
	util "clase_5/jwt"
	"clase_5/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func ValidarJWT(h http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		header := r.Header.Get("Authorization")

		if len(header) == 0 {
			outp := map[string]interface{}{
				"error":  "Token de acesso no encontrado",
				"estado": http.StatusUnauthorized,
			}
			json.NewEncoder(w).Encode(outp)
			return
		}
		splitBearer := strings.Split(header, " ")
		if len(splitBearer) != 2 {
			outp := map[string]interface{}{
				"error":  "Token de acesso inválido 1",
				"estado": http.StatusUnauthorized,
			}
			json.NewEncoder(w).Encode(outp)
			return
		}
		splitToken := strings.Split(splitBearer[1], ".")
		if len(splitToken) != 3 {
			outp := map[string]interface{}{
				"error":  "Token de acesso inválido 2",
				"estado": http.StatusUnauthorized,
			}
			json.NewEncoder(w).Encode(outp)
			return
		}

		tk := strings.TrimSpace(splitBearer[1])
		tokenss, errT := jwt.Parse(tk, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("error inesperado: ")
			}
			return util.Clave, nil
		})

		if errT != nil {
			outp := map[string]interface{}{
				"error":  "Token de acesso inválido 3: " + errT.Error(),
				"estado": http.StatusUnauthorized,
			}
			json.NewEncoder(w).Encode(outp)
			return
		}

		if claims, ok := tokenss.Claims.(jwt.MapClaims); ok && tokenss.Valid {
			usuario := models.User{}
			if err := database.Conectar().
				Where("username=?", claims["nombre"]).First(&usuario); err.Error != nil {
				outp := map[string]interface{}{
					"error":  "Token de acesso inválido",
					"estado": http.StatusUnauthorized,
				}
				json.NewEncoder(w).Encode(outp)
				return

			} else {
				h.ServeHTTP(w, r)
			}

		}
	}
}
