package handlers

import (
	dto "clase_5/DTO"
	"clase_5/database"
	util "clase_5/jwt"
	"clase_5/models"
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Registro(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	var registro dto.UserDTO
	if err := json.NewDecoder(req.Body).Decode(&registro); err != nil {
		outp := map[string]any{
			"estado": http.StatusBadRequest,
			"data":   "error",
		}
		json.NewEncoder(res).Encode(outp)
		return
	}

	usuario := models.User{}
	if database.Conectar().
		Where("email=?", registro.Email).
		Limit(1).
		Find(&usuario).RowsAffected > 0 {
		outp := map[string]interface{}{
			"estado": http.StatusBadRequest,
			"data":   "El correo ya esta usado",
		}
		json.NewEncoder(res).Encode(outp)
		return
	} else {
		cost := 8
		byte, _ := bcrypt.GenerateFromPassword([]byte(registro.Password), cost)
		datos := models.User{Username: registro.Username, Password: string(byte), Email: registro.Email, PerfilId: registro.PerfilId, Fecha: time.Now()}
		database.Conectar().Create(&datos)
		outp := map[string]interface{}{
			"estado": http.StatusOK,
			"data":   "El usuario se creo con exito",
		}
		json.NewEncoder(res).Encode(outp)
		return
	}
}

func Login(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var login dto.LoginDTO

	if err := json.NewDecoder(req.Body).Decode(&login); err != nil {
		outp := map[string]any{
			"estado": http.StatusBadRequest,
			"data":   "error",
		}
		json.NewEncoder(res).Encode(outp)
		return
	}

	usuario := models.User{}
	if database.Conectar().
		Where("username=?", login.Username).
		Limit(1).
		Find(&usuario).RowsAffected > 0 {

		passBytes := []byte(login.Password)
		passBd := []byte(usuario.Password)

		err := bcrypt.CompareHashAndPassword(passBd, passBytes)
		if err != nil {
			outp := map[string]interface{}{
				"estado": http.StatusBadRequest,
				"data":   "Credenciales incorrectas",
			}
			json.NewEncoder(res).Encode(outp)
			return
		} else {

			jwtKey, err := util.GenerarToken(usuario)
			if err != nil {
				outp := map[string]interface{}{
					"estado": http.StatusBadRequest,
					"data":   "error en el token: " + err.Error(),
				}
				json.NewEncoder(res).Encode(outp)
				return
			}

			response := dto.ResponseDTO{
				User:  usuario.Username,
				Token: jwtKey,
			}
			json.NewEncoder(res).Encode(response)
		}
	} else {
		outp := map[string]interface{}{
			"estado": http.StatusBadRequest,
			"data":   "El usuario no existe",
		}

		json.NewEncoder(res).Encode(outp)
		return
	}
}

func Protegido(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	//res.WriteHeader(http.StatusUnauthorized)
	outp := map[string]interface{}{
		"estado": http.StatusUnauthorized,
		"data":   "protegido",
	}

	json.NewEncoder(res).Encode(outp)
}
