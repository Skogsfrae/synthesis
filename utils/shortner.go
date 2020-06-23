package utils

var DICT string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
var SALT int64 = 39201928349320

func ShortenUrl(id int64) string {
	var shortUrl string
	var idx = id + SALT

	for idx > 0 {
		shortUrl += string(DICT[idx%62])
		idx /= 62
	}

	return shortUrl
}
