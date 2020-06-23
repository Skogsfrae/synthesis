package models

import (
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/skogsfrae/synthesis/configuration"
)

var db *pg.DB

func InitDB() {
	opts, err := pg.ParseURL(configuration.Config.PostgresUri)
	if err != nil {
		log.Fatalf("Invalid postgres uri")
	}

	db = pg.Connect(opts)
	_, err = db.Exec("SELECT 1")
	if err != nil {
		log.Fatalln("Could not connect to Postgres", err)
	}

	log.Println("Connected to Postgres")
	db.CreateTable((*Url)(nil), &orm.CreateTableOptions{})
}
