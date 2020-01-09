package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/camilaleniss/api-rest-go/connection"
	repository "github.com/camilaleniss/api-rest-go/repository"
	domain "github.com/camilaleniss/api-rest-go/repository/domain"
	"github.com/go-chi/chi"
)

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
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	//ESTO DEBERÍA BUSCARLO EN EL API
	payload, err := d.repo.SearchDomain(strconv.Itoa(id))

	fmt.Println(payload.Host)
	fmt.Println(payload.Last_search)
	fmt.Println(payload.Ssl_grade)
	fmt.Println(payload.Ssl_previous_grade)

	if err != nil {
		respondWithError(w, http.StatusNoContent, "Content not found")
	}

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
