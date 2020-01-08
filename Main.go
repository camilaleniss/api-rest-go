package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/camilaleniss/api-rest-go/model"
)

const API_DOMAINS_URL = "https://api.ssllabs.com/api/v3/analyze?host="

const DOMAIN_NAME_EXAMPLE = "truora.com"

func main() {
	//router := chi.NewRouter()
	downloadJSON()

	/*
		db := connection.GetConnection()

		if _, err := db.Exec(
			`INSERT INTO domain (host, ssl_grade, ssl_previous_grade)
			VALUES ('truora.com', 'A', 'A');`); err != nil {
			log.Fatal(err)
		}
	*/
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

	if len(domain1.Erros) != 0 {
		fmt.Println("Server is Down")
	} else {
		fmt.Println(domain1.Endpoints)
	}

	fmt.Println("---Who Is Info---")
	makeWhoIs(domain1)
}

func makeWhoIs(domain model.DomainApi) {
	for i := 0; i < len(domain.Endpoints); i++ {
		fmt.Println("----------------SERVER NUMERO --------------")

		fmt.Println(model.WhoisServerAttributes(domain.Endpoints[i]))
	}
}

func makeURL(domain string) string {
	return API_DOMAINS_URL + domain
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
