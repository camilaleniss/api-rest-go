package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const API_DOMAINS_URL = "https://api.ssllabs.com/api/v3/analyze?host="

const DOMAIN_NAME_EXAMPLE = "truora.com"

type server struct {
	Host     string `json:"host"`
	Protocol string `json:"protocol"`
}

func main() {
	//router := chi.NewRouter()

	downloadJSON()
}

func downloadJSON() {
	spaceClient := http.Client{
		Timeout: time.Second * 2,
	}

	url := makeURL()

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)

	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	domain1 := Domain{}

	jsonErr := json.Unmarshal(body, &domain1)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(domain1.Host)

}

func makeURL() string {
	return API_DOMAINS_URL + DOMAIN_NAME_EXAMPLE
}

//var domains []Domain

/*
type mensaje struct {
	msg string
}

func (m mensaje) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.msg)
}
*/

/*
func main() {

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/api/{id}", SearchDomain)

		router.Get("/", holaraiz)
		router.Get("/search", holabusqueda)
		router.Get("/history", holaconsulta)
		router.Handle("/prueba", msg)


	http.ListenAndServe(":8080", router)

}
*/

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
