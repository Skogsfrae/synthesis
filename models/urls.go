package models

import (
	"time"

	. "github.com/skogsfrae/synthesis/configuration"
	"github.com/skogsfrae/synthesis/utils"
)

type Url struct {
	ID             int64
	FullUrl        string
	ShortUrl       string
	ExpirationTime time.Time
}

func SaveUrl(urlString string) (*Url, error) {
	url := &Url{
		FullUrl:        urlString,
		ExpirationTime: time.Now().Add(time.Duration(Config.KeyTimer) * time.Second),
	}

	_, err := db.Model(url).Insert()
	url.ShortUrl = utils.ShortenUrl(url.ID)
	_, err = db.Model(url).
		Set("short_url = ?", url.ShortUrl).
		Where("id = ?", url.ID).
		Update()

	return url, err
}

func FindUrl(shortnedUrl string) (*Url, error) {
	url := new(Url)
	err := db.Model(url).Where("short_url = ?", shortnedUrl).
		Where("expiration_time > ?", time.Now()).
		Select()
	return url, err
}

func DeleteUrl(shortenedUrl string) (bool, error) {
	url := new(Url)
	res, err := db.Model(url).Where("short_url = ?", shortenedUrl).Delete()
	if res.RowsAffected() > 0 {
		return true, err
	} else {
		return false, err
	}
}
