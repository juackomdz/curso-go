package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	id := vars["id"]
	user := models.Usuario{}
	db.Conectar().Find(&user, id)

	res.WriteHeader(http.StatusAccepted)

	outp := map[string]any{
		"estado": http.StatusAccepted,
		"data":   user,
	}
	json.NewEncoder(res).Encode(outp)

	//json.NewEncoder(res).Encode(user)

}
