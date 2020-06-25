package routes

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"

	"github.com/skogsfrae/synthesis/models"
)

func fetchUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url, err := models.FindUrl(vars["shortUrl"])
	if err != nil {
		http.NotFound(w, r)
	} else {
		http.Redirect(w, r, url.FullUrl, 301)
	}
}

func shortenUrl(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	// url validation
	parsedUrl, err := url.Parse(string(body))
	if err != nil {
		log.Fatalln(err)
	}

	// add schema if non present
	if parsedUrl.Scheme == "" {
		parsedUrl.Scheme = "http"
	}

	url, _ := models.SaveUrl(parsedUrl.String())
	w.Write([]byte(url.ShortUrl))
}

func getFullUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url, err := models.FindUrl(vars["shortUrl"])
	if err != nil {
		http.NotFound(w, r)
	} else {
		w.Write([]byte(url.FullUrl))
	}
}

func deleteUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, err := models.DeleteUrl(vars["shortUrl"])
	if err != nil {
		http.NotFound(w, r)
	} else {
		w.Write([]byte("ok"))
	}
}
