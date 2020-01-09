package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/camilaleniss/api-rest-go/connection"
	dh "github.com/camilaleniss/api-rest-go/handler/http"
	"github.com/camilaleniss/api-rest-go/model"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const API_DOMAINS_URL = "https://api.ssllabs.com/api/v3/analyze?host="

const DOMAIN_NAME_EXAMPLE = "truora.com"

func main() {
	connection, err := connection.GetConnection()
	if err != nil {
		fmt.Println(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	dHandler := dh.NewDomainHandler(connection)
	r.Route("/", func(rt chi.Router) {
		rt.Mount("/api", domainRouter(dHandler))
	})

	fmt.Println("Server listen at :8080")
	http.ListenAndServe(":8080", r)

}

func domainRouter(pHandler *dh.DomainHnd) http.Handler {
	r := chi.NewRouter()
	r.Get("/", pHandler.InitApi)
	r.Get("/{id}", pHandler.GetByID)
	r.Get("/domains", pHandler.Fetch)
	return r
}

func downloadJSON() {

	spaceClient := http.Client{
		Timeout: time.Second * 2,
	}

	url := makeURL(DOMAIN_NAME_EXAMPLE)

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

	domain1 := model.DomainApi{}

	jsonErr := json.Unmarshal(body, &domain1)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	//AQUI DEBERÍA COMENZAR A ARMAR EL JSON

	if len(domain1.Erros) != 0 {
		fmt.Println("Server is Down")
	} else {
		fmt.Println(domain1.Endpoints)
	}

}

func makeWhoIs(server model.ServerApi) {
	owner, country := model.WhoisServerAttributes(server)
	fmt.Println(owner)
	fmt.Println(country)
}

func makeURL(domain string) string {
	return API_DOMAINS_URL + domain
}

func CreateDomain(domain model.DomainApi) bool {

	is_Down := len(domain.Erros) != 0

	/*
		BUSCAR SI EXISTE EN LA BD

		SI ESTA CAIDO Y NO EXISTE:
			AÑADIRLO CON EL SS_GRADE Y ESO CON EL DEFAULT
		SI ESTA CAIDO Y EXISTE:
			TOMAR SUS DATOS DE SSL_GRADE

		REVISAR BIEN ESTA OPCIÓN

		SI NO ESTÁ CAIDO Y NO EXISTE
			AÑADIRLO COMO NUEVO

		SI NO ESTA CAIDO Y EXISTE
			COMPARAR SSL_GRADE
	*/

	if is_Down {

	} else {
		ssl_grade := model.GenerateSSLGrade(domain.Endpoints)
		fmt.Println(ssl_grade)

	}

	return true
}

func DomainIsDown(domain model.DomainApi) {

}

func DomainIsUp(domain model.DomainApi) {

}

func PruebaSSL_Grade() {
	server1 := model.ServerApi{"1.2.3.4.5", "B+"}
	server2 := model.ServerApi{"1.2.3.4.6", "A"}
	server3 := model.ServerApi{"1.2.3.4.4", "B"}

	server := []model.ServerApi{server1, server2, server3}

	respuesta := model.GenerateSSLGrade(server)

	fmt.Println(respuesta)
}

/*
func getJSON2(url string, result interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("cannot fetch URL %q: %v", url, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected http GET status: %s", resp.Status)
	}
	// We could check the resulting content type
	// here if desired.
	err := json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return fmt.Errorf("cannot decode JSON: %v", err)
	}
	return nil
}

*/

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
