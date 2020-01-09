package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/camilaleniss/api-rest-go/connection"
	"github.com/camilaleniss/api-rest-go/model"
	repository "github.com/camilaleniss/api-rest-go/repository"
	domain "github.com/camilaleniss/api-rest-go/repository/domain"
	"github.com/go-chi/chi"
)

//This API gives us information about the domain, its servers and the ssl grade of each server
const API_DOMAINS_URL = "https://api.ssllabs.com/api/v3/analyze?host="

func NewDomainHandler(db *connection.DB) *DomainHnd {
	return &DomainHnd{
		repo: domain.NewSQLDomainRepo(db.Cock),
	}
}

type DomainHnd struct {
	repo repository.DomainRepo
}

//First view of the Api
func (d *DomainHnd) InitApi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>This is the first view</h1>")
}

func (d *DomainHnd) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "domainId")

	//ESTO DEBERÍA BUSCARLO EN EL API
	payload := getJsonDomain(d, id)

	/*
		if payload == nil {
			respondWithError(w, http.StatusNoContent, "Content not found")
		}
	*/
	respondwithJSON(w, http.StatusOK, payload)
}

//Search all the domains
func (d *DomainHnd) Fetch(w http.ResponseWriter, r *http.Request) {
	payload := d.repo.SearchDomains()

	for i := 0; i < len(payload); i++ {
		fmt.Println(payload[i].Host)
	}

	respondwithJSON(w, http.StatusOK, payload)
}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}

func getJsonDomain(d *DomainHnd, host string) model.Domain {

	var ssl_grade_bd, ssl_grade_obtained, ssl_previous_grade, logo, title string
	var servers_api []model.ServerApi

	domain_api := downloadJSONApi(host)

	is_Down := len(domain_api.Erros) != 0

	domain_bd, _ := searchDomainDB(d, host)

	exists := domain_bd.Host != ""

	if is_Down {
		servers_api = []model.ServerApi{}
	} else {
		servers_api = domain_api.Endpoints
	}

	if exists {
		ssl_grade_bd = domain_bd.Ssl_grade
		last_search := domain_bd.Last_search
		if is_Down {
			ssl_grade_obtained = ssl_grade_bd
			ssl_previous_grade = domain_bd.Ssl_previous_grade
		} else {
			ssl_grade_obtained, ssl_previous_grade = compareOneHourBefore(servers_api, ssl_grade_bd, last_search)
		}

	} else {
		ssl_grade_obtained = model.GenerateSSLGrade(servers_api)
		ssl_previous_grade = ssl_grade_obtained
		d.repo.CreateDomain(host, ssl_grade_obtained, ssl_previous_grade)
	}

	domain := createDomain(servers_api, ssl_grade_obtained, ssl_previous_grade, logo, title, is_Down)

	return domain
}

func createDomain(serversapi []model.ServerApi, ssl_grade string, ssl_previous_grade string, logo string, title string, is_down bool) model.Domain {

	var domain model.Domain
	var servers []model.Server

	if is_down != true {
		servers = getServersFromServersApi(serversapi)
	}

	domain.Servers = servers
	domain.ServersChanged = ssl_grade != ssl_previous_grade
	domain.SslGrade = ssl_grade
	domain.PreviousSslGrade = ssl_previous_grade
	domain.Logo = logo
	domain.Title = title
	domain.IsDown = is_down

	return domain
}

func getServersFromServersApi(serversapi []model.ServerApi) []model.Server {
	var servers []model.Server
	for i := 0; i < len(serversapi); i++ {
		ip := serversapi[i].IpAddress
		grade := serversapi[i].Grade
		owner, country := makeWhoIs(serversapi[i])
		server := model.Server{
			Address:  ip,
			SslGrade: grade,
			Owner:    owner,
			Country:  country,
		}
		servers = append(servers, server)
	}
	return servers
}

func downloadJSONApi(domainurl string) model.DomainApi {

	spaceClient := http.Client{
		Timeout: time.Second * 2,
	}

	url := makeURL(domainurl)

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

	//domain1 contains the info of the domain obtained from the api
	domain1 := model.DomainApi{}

	jsonErr := json.Unmarshal(body, &domain1)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return domain1
}

func makeWhoIs(server model.ServerApi) (string, string) {
	owner, country := model.WhoisServerAttributes(server)
	return owner, country
}

func makeURL(domain string) string {
	return API_DOMAINS_URL + domain
}

func searchDomainDB(d *DomainHnd, host string) (connection.DomainBD, error) {
	payload, err := d.repo.SearchDomain(host)
	d.repo.UpdateDomainVisit(host)
	if err != nil {
		return connection.DomainBD{}, err
	}
	return payload, err
}

func compareOneHourBefore(servers_api []model.ServerApi, ssl_grade_bd string, last_search time.Time) (string, string) {
	today := time.Now()
	comparator := last_search.Add(1 * time.Hour)

	if today.Before(comparator) {

		return ssl_grade_bd, ssl_grade_bd
	} else {
		return model.GenerateSSLGrade(servers_api), ssl_grade_bd
	}
}
