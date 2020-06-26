package routes

import (
	"fmt"
	"log"
	"strings"

	"github.com/gorilla/mux"
	_ "github.com/skogsfrae/synthesis/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func InitRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/{shortUrl:[a-z,A-Z,0-9]+$}", visitUrl).Methods("GET")

	s := r.PathPrefix("/api/url").Subrouter()
	s.HandleFunc("", shortenUrl).Methods("PUT")
	s.HandleFunc("/{shortUrl:[a-z,A-Z,0-9]+$}", getFullUrl).Methods("GET")
	s.HandleFunc("/{shortUrl:[a-z,A-Z,0-9]+}/visit-count", getUrlVisitCount).Methods("GET")
	s.HandleFunc("/{shortUrl:[a-z,A-Z,0-9]+$}", deleteUrl).Methods("DELETE")

	swag := r.PathPrefix("/swagger").Subrouter()
	swag.HandleFunc("/{path}", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3000/swagger/doc.json"),
	))

	err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			fmt.Println("ROUTE:", pathTemplate)
		}
		pathRegexp, err := route.GetPathRegexp()
		if err == nil {
			fmt.Println("Path regexp:", pathRegexp)
		}
		queriesTemplates, err := route.GetQueriesTemplates()
		if err == nil {
			fmt.Println("Queries templates:", strings.Join(queriesTemplates, ","))
		}
		queriesRegexps, err := route.GetQueriesRegexp()
		if err == nil {
			fmt.Println("Queries regexps:", strings.Join(queriesRegexps, ","))
		}
		methods, err := route.GetMethods()
		if err == nil {
			fmt.Println("Methods:", strings.Join(methods, ","))
		}
		fmt.Println()
		return nil
	})

	if err != nil {
		log.Fatalln(err)
	}

	return r
}
