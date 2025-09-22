package service

import (
	"crypto/sha256"
	"encoding/base64"
	"linky/database"
	"log"
	"strconv"
)

var count int

func URLTransform(longURL string) string {
	var shortURL string

	for {
		h := sha256.New()
		h.Write([]byte(longURL))
		shortURL = base64.URLEncoding.EncodeToString([]byte(h.Sum(nil)[:8]))

		query := "INSERT INTO urls (short_url, long_url) VALUES ($1, $2)"
		_, err := database.PostgresClient.Exec(query, shortURL, longURL)
		if err != nil {
			log.Println(err)
			// there is collision
			log.Println("There is collision")
			longURL = longURL + ":" + strconv.Itoa(count)
			count++
			continue
		}
		break
	}
	log.Println("Short URL created")
	return shortURL
}
