package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	dto "clase_5/DTO"
	db "clase_5/database"
	"clase_5/models"

	"github.com/gorilla/mux"
)

type Respuesta struct {
	Estado  string
	Mensaje string
}

func Ejemplo(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	outp, _ := json.Marshal(Respuesta{"ok", "hola mundo"})
	fmt.Fprintln(res, string(outp))
}

func Ejemplo_params(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	fmt.Fprintln(res, "hola con parametros: "+id)
}

func Ejemplo_post(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var cate dto.CategoriaDTO
	err := json.NewDecoder(req.Body).Decode(&cate)
	if err != nil {
		outp := map[string]string{
			"estado":  "error",
			"mensaje": "error al procesar",
		}
		json.NewEncoder(res).Encode(outp)
		return
	}
	outp := map[string]string{
		"estado":        "ok",
		"mensaje":       "hola mundo 3",
		"nombre":        cate.Nombre,
		"Authorization": req.Header.Get("Authorization"),
		"Test":          req.Header.Get("test"),
	}
	json.NewEncoder(res).Encode(outp)
}

/*
func Ejemplo_post(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	//outp, _ := json.Marshal(Respuesta{"ok", "hola mundo 2"})
	outp := map[string]string{
		"estado":  "ok",
		"mensaje": "hola mundo 2",
	}
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(outp)
	//fmt.Fprintln(res, string(outp))
}
*/
/*
	func Ejemplo_post(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "hola con post")
	}
*/
func Ejemplo_put(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	fmt.Fprintln(res, "hola con put | id modificado: "+id)
}

func Ejemplo_delete(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	fmt.Fprintln(res, "hola con delete | id eliminado: "+id)
}

func Ejemplo_querystring(res http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	fmt.Fprintln(res, "hola con query string | id: "+id)
}

//--------------------------------handlers con db-------------------------------//

func Get_db(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	usuarios := models.Usuarios{}
	db.Conectar().Find(&usuarios)

	json.NewEncoder(res).Encode(usuarios)
}

func Get_db_id(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	user := models.Usuario{}
	if err := db.Conectar().First(&user, id); err.Error != nil {
		res.WriteHeader(http.StatusNotFound)
		outp := map[string]any{
			"estado": http.StatusNotFound,
			"data":   "error",
		}

		json.NewEncoder(res).Encode(outp)
		return
	} else {

		res.WriteHeader(http.StatusOK)
		/*
			outp := map[string]any{
				"estado": http.StatusAccepted,
				"data":   user,
			}

			json.NewEncoder(res).Encode(outp)
		*/
		json.NewEncoder(res).Encode(user)
	}

}

func Post_db(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	var user dto.UsuarioDTO
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		outp := map[string]any{
			"estado": http.StatusBadRequest,
			"data":   "error",
		}
		json.NewEncoder(res).Encode(outp)
		return
	}

	datos := models.Usuario{Nombre: user.Nombre, Apellido: user.Apellido, Email: user.Email}
	db.Conectar().Create(&datos)

	res.WriteHeader(http.StatusCreated)

	outp := map[string]any{
		"estado": http.StatusCreated,
		"data":   datos,
	}
	json.NewEncoder(res).Encode(outp)
}

func Put_db(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(req)

	idU, _ := strconv.Atoi(vars["id"])
	var user dto.UsuarioDTO
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		outp := map[string]any{
			"estado": http.StatusBadRequest,
			"data":   "error",
		}
		json.NewEncoder(res).Encode(outp)
		return
	}

	usuario := models.Usuario{}
	if err := db.Conectar().First(&usuario, idU); err.Error != nil {
		outp := map[string]any{
			"estado": http.StatusBadRequest,
			"data":   "error",
		}
		json.NewEncoder(res).Encode(outp)
		return
	} else {
		usuario.Nombre = user.Nombre
		usuario.Apellido = user.Apellido
		usuario.Email = user.Email

		db.Conectar().Save(&usuario)
		outp := map[string]any{
			"estado": http.StatusOK,
			"data":   "Modificado con exito",
		}
		json.NewEncoder(res).Encode(outp)
		return
	}

}

func Delete_db(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(req)

	idU, _ := strconv.Atoi(vars["id"])

	var user dto.UsuarioDTO
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		outp := map[string]any{
			"estado": http.StatusBadRequest,
			"data":   "error",
		}
		json.NewEncoder(res).Encode(outp)
		return
	}

	usuario := models.Usuario{}
	if err := db.Conectar().First(&usuario, idU); err.Error != nil {
		outp := map[string]any{
			"estado": http.StatusNotFound,
			"data":   "error",
		}
		json.NewEncoder(res).Encode(outp)
		return
	} else {

		db.Conectar().Delete(&usuario)
		outp := map[string]any{
			"estado": http.StatusOK,
			"data":   "Eliminado con exito",
		}
		json.NewEncoder(res).Encode(outp)
		return
	}
}

func Post_db_c(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	var canal dto.CanalDTO
	if err := json.NewDecoder(req.Body).Decode(&canal); err != nil {
		outp := map[string]any{
			"estado": http.StatusBadRequest,
			"data":   "error",
		}
		json.NewEncoder(res).Encode(outp)
		return
	}

	datos := models.Canal{Nombre: canal.Nombre, UsuarioId: canal.UsuarioId}
	db.Conectar().Create(&datos)

	res.WriteHeader(http.StatusCreated)

	outp := map[string]any{
		"estado": http.StatusCreated,
		"data":   "Creado con exito",
	}
	json.NewEncoder(res).Encode(outp)

}

func Get_db_c(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	canales := models.Canales{}
	db.Conectar().Preload("Usuario").Find(&canales)

	json.NewEncoder(res).Encode(canales)
}

func Get_db_id_c(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(req)

	id, _ := strconv.Atoi(vars["id"])

	canales := models.Canales{}
	if err := db.Conectar().Preload("Usuario").First(&canales, id); err.Error != nil {
		res.WriteHeader(http.StatusNotFound)
		outp := map[string]interface{}{
			"estado": http.StatusNotFound,
			"data":   "No encontrdo",
		}
		json.NewEncoder(res).Encode(outp)
		return
	}

	json.NewEncoder(res).Encode(canales)
}

func Put_db_c(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(req)

	idU, _ := strconv.Atoi(vars["id"])
	var canalU dto.CanalDTO
	if err := json.NewDecoder(req.Body).Decode(&canalU); err != nil {
		outp := map[string]any{
			"estado": http.StatusBadRequest,
			"data":   "error",
		}
		json.NewEncoder(res).Encode(outp)
		return
	}

	canal := models.Canal{}
	if err := db.Conectar().First(&canal, idU); err.Error != nil {
		outp := map[string]any{
			"estado": http.StatusBadRequest,
			"data":   "error",
		}
		json.NewEncoder(res).Encode(outp)
		return
	} else {
		canal.Nombre = canalU.Nombre
		//canal.UsuarioId = canalU.UsuarioId

		if err := db.Conectar().Save(&canal); err.Error != nil {
			outp := map[string]any{
				"estado": http.StatusBadGateway,
				"data":   err.Error.Error(),
			}
			json.NewEncoder(res).Encode(outp)
			return
		}
		outp := map[string]any{
			"estado": http.StatusOK,
			"data":   "Modificado con exito",
		}
		json.NewEncoder(res).Encode(outp)
		return
	}

}
func Delete_db_c(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(req)

	idU, _ := strconv.Atoi(vars["id"])

	canalU := models.Canal{}
	if err := db.Conectar().First(&canalU, idU); err.Error != nil {
		outp := map[string]any{
			"estado": http.StatusNotFound,
			"data":   "error",
		}
		json.NewEncoder(res).Encode(outp)
		return
	} else {

		db.Conectar().Delete(&canalU)
		outp := map[string]any{
			"estado": http.StatusOK,
			"data":   "Eliminado con exito",
		}
		json.NewEncoder(res).Encode(outp)
		return
	}
}
