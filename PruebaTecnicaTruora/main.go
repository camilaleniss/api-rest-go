package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

const API_DOMAINS_URL = "https://api.ssllabs.com/api/v3/analyze?host="

var domains []Domain

/*
type mensaje struct {
	msg string
}

func (m mensaje) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.msg)
}
*/

func main() {

	router := chi.NewRouter()
	//router.Use(middleware.Logger)
	//router.Use(middleware.Recoverer)

	router.Get("/api/{id}", SearchDomain)
	/*
		router.Get("/", holaraiz)
		router.Get("/search", holabusqueda)
		router.Get("/history", holaconsulta)
		router.Handle("/prueba", msg)
	*/

	http.ListenAndServe(":8080", router)

}

func SearchDomain(w http.ResponseWriter, r *http.Request) {

}

/*
func holaraiz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hola Mundo Divino</h1>")
}

func holabusqueda(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hola Búsqueda</h1>")
}

func holaconsulta(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hola Consulta</h1>")
}
*/
