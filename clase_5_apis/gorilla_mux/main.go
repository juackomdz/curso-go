package main

import (
	"clase_5/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	//models.Migrar()
	mux := mux.NewRouter()

	prefix := "/api/v1"
	mux.HandleFunc(prefix+"/ejemplo", handlers.Ejemplo).Methods("GET")
	mux.HandleFunc(prefix+"/ejemplo/{id:[0-9]+}", handlers.Ejemplo_params).Methods("GET")
	mux.HandleFunc(prefix+"/ejemplo", handlers.Ejemplo_post).Methods("POST")
	mux.HandleFunc(prefix+"/ejemplo/{id}", handlers.Ejemplo_put).Methods("PUT")
	mux.HandleFunc(prefix+"/ejemplo/{id}", handlers.Ejemplo_delete).Methods("DELETE")
	mux.HandleFunc(prefix+"/query-string", handlers.Ejemplo_querystring).Methods("GET")

	prefix2 := "/api/v2"
	mux.HandleFunc(prefix2+"/ejemplo", handlers.Get_db).Methods("GET")
	mux.HandleFunc(prefix2+"/ejemplo/{id}", handlers.Get_db_id).Methods("GET")

	log.Fatal(http.ListenAndServe(":8084", mux))

}
