package main

import (
	"fmt"
	"net/http"

	"github.com/camilaleniss/api-rest-go/connection"
	dh "github.com/camilaleniss/api-rest-go/handler/http"

	//"github.com/camilaleniss/api-rest-go/model"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

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
	r.Get("/{domainId}", pHandler.GetByID)
	r.Get("/domains", pHandler.Fetch)
	return r
}
