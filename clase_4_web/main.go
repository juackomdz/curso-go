package main

import (
	"clase_4/rutas"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	mux := mux.NewRouter()

	mux.HandleFunc("/", rutas.Home)
	mux.HandleFunc("/nosotros", rutas.Nosotros)
	mux.HandleFunc("/parametros/{id}/{slug}", rutas.Params)
	mux.HandleFunc("/query-string", rutas.QueryStrings)

	serv := &http.Server{
		Addr:    "localhost:8082",
		Handler: mux,
	}

	log.Fatal(serv.ListenAndServe())
	//mux := http.NewServeMux()

	/*
		http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
			fmt.Fprintln(res, "Hola GO")
		})

		log.Fatal(http.ListenAndServe("localhost:8082", nil))

	*/
}
