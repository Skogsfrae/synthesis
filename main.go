package main

import (
	"log"
	"net/http"

	. "github.com/skogsfrae/synthesis/configuration"
	"github.com/skogsfrae/synthesis/models"
	"github.com/skogsfrae/synthesis/routes"
)

func main() {
	// TODO: implement command line parameters
	LoadConfiguration("./configuration.yaml")
	models.InitDB()
	addr := Config.Host + ":" + Config.Port

	r := routes.InitRoutes()
	s := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	log.Println("Server listening on:", addr)
	log.Fatal(s.ListenAndServe())
}
