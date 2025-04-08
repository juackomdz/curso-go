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
	mux.HandleFunc(prefix2+"/ejemplo", handlers.Post_db).Methods("POST")
	mux.HandleFunc(prefix2+"/ejemplo/{id}", handlers.Put_db).Methods("PUT")
	mux.HandleFunc(prefix2+"/ejemplo/{id}", handlers.Delete_db).Methods("DELETE")

	mux.HandleFunc(prefix2+"/canal", handlers.Post_db_c).Methods("POST")
	mux.HandleFunc(prefix2+"/canal", handlers.Get_db_c).Methods("GET")
	mux.HandleFunc(prefix2+"/canal/{id}", handlers.Get_db_id_c).Methods("GET")
	mux.HandleFunc(prefix2+"/canal/{id}", handlers.Put_db_c).Methods("PUT")
	mux.HandleFunc(prefix2+"/canal/{id}", handlers.Delete_db_c).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8084", mux))

}
