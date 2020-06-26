package routes

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/skogsfrae/synthesis/models"
)

// Save full url godoc
// @Summary Save full url and get a short url associated to it on Synthesis
// @tags Api
// @Body {string} full url
// @Success 200 {string} full url
// @Failure 404 url associated to short url not found
// @Router /api/url/{shortUrl} [put]
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

// Get full url godoc
// @Summary Get full url associated to a short url saved on Synthesis
// @tags Api
// @Param shortUrl path string true "Short url"
// @Success 200 {string} full url
// @Failure 404 url associated to short url not found
// @Router /api/url/{shortUrl} [get]
func getFullUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url, err := models.FindUrl(vars["shortUrl"])
	if err != nil {
		http.NotFound(w, r)
	} else {
		w.Write([]byte(url.FullUrl))
	}
}

// Url visit count godoc
// @Summary Get the number of visits a short url redirected from Synthesis
// @tags Api
// @Param shortUrl path string true "Short url"
// @Success 200 {int} number of visits
// @Failure 404 url associated to short url not found
// @Router /api/url/{shortUrl}/visit-count [get]
func getUrlVisitCount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url, err := models.FindUrl(vars["shortUrl"])
	if err != nil {
		http.NotFound(w, r)
	} else {
		w.Write([]byte(strconv.FormatInt(url.VisitCount, 10)))
	}
}

// Visit url godoc
// @Summary Visit short url connected website
// @Param shortUrl path string true "Short url"
// @Success 300 redirects to full url
// @Failure 404 url associated to short url not found or expired
// @Router /{shortUrl} [get]
func visitUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url, err := models.VisitUrl(vars["shortUrl"])
	if err != nil {
		http.NotFound(w, r)
	} else {
		http.Redirect(w, r, url, 301)
	}
}

// Remove url godoc
// @Summary Remove url from Synthesis
// @tags Api
// @Param shortUrl path string true "Short url"
// @Success 200 {string} ok url removed
// @Failure 404 url associated to short url not found
// @Router /api/url/{shortUrl} [delete]
func deleteUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, err := models.DeleteUrl(vars["shortUrl"])
	if err != nil {
		http.NotFound(w, r)
	} else {
		w.Write([]byte("ok"))
	}
}
