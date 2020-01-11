package main

import (
	"fmt"
	"net/http"

	"github.com/camilaleniss/api-rest-go/server/connection"
	dh "github.com/camilaleniss/api-rest-go/server/handler/http"

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

	fmt.Println("Server listen at :8082")
	http.ListenAndServe(":8082", r)

}

func domainRouter(pHandler *dh.DomainHnd) http.Handler {
	r := chi.NewRouter()
	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	// cors := cors.New(cors.Options{
	// 	AllowedOrigins:     []string{"*"},
	// 	AllowedMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowedHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	// 	ExposedHeaders:     []string{"Link"},
	// 	AllowCredentials:   true,
	// 	OptionsPassthrough: true,
	// 	MaxAge:             3599, // Maximum value not ignored by any of major browsers
	// })

	// r.Use(cors.Handler)

	r.Get("/{domainId}", pHandler.GetByID)
	r.Get("/domains", pHandler.Fetch)
	return r
}
