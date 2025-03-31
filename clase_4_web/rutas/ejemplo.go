package rutas

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(res http.ResponseWriter, req *http.Request) {

	fmt.Fprintln(res, "Hola desde GoLang")
}

func Nosotros(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Pagina nosotros en go")
}

func Params(res http.ResponseWriter, req *http.Request) {
	//id:=req.p
	varss := mux.Vars(req)
	fmt.Fprintln(res, "id="+varss["id"]+" | slug="+varss["slug"])

}

func QueryStrings(res http.ResponseWriter, req *http.Request) {

	fmt.Println(req.URL)
	fmt.Println(req.URL.RawQuery)
	fmt.Println(req.URL.Query())
	fmt.Println(req.URL.Query().Get("id"))

	id := req.URL.Query().Get("id")
	slug := req.URL.Query().Get("slug")

	fmt.Fprintln(res, "id: "+id+"| slug: "+slug)
	fmt.Fprintln(res, "-------------------------")
	fmt.Fprintf(res, "id=%s | slug=%s", id, slug)

}
