package routes

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/{shortUrl}", fetchUrl).Methods("GET")
	r.HandleFunc("/", shortenUrl).Methods("POST")
	return r
}
