package main

import (
	"log"
	"net/http"

	. "github.com/skogsfrae/synthesis/configuration"
	"github.com/skogsfrae/synthesis/docs"
	"github.com/skogsfrae/synthesis/models"
	"github.com/skogsfrae/synthesis/routes"
)

// @title Synthesis API
// @version 1.0
// @tag.name Api

// @contact.name API Support
// @contact.url https://github.com/Skogsfrae/synthesis
// @contact.email skogsfrae@gmail.com
func main() {
	// TODO: implement command line parameters
	LoadConfiguration("./configuration.yaml")
	models.InitDB()
	addr := Config.Host + ":" + Config.Port
	docs.SwaggerInfo.Host = addr

	r := routes.InitRoutes()
	s := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	log.Println("Server listening on:", addr)
	log.Fatal(s.ListenAndServe())
}
