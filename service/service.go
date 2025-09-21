package service

import (
	"crypto/sha256"
	"encoding/base64"
	"linky/store"
	"log"
	"strconv"
)

var MapCreated bool
var count int

func URLTransform(longURL string) string {
	var shortURL string

	h := sha256.New()

	for {
		shortURL = base64.URLEncoding.EncodeToString([]byte(h.Sum(nil)[:8]))
		if _, ok := store.URLs[shortURL]; !ok {
			store.URLs[shortURL] = longURL
			break
		}
		// there is collision
		log.Println("There is collision")
		longURL = longURL + ":" + strconv.Itoa(count)
		count++
		h.Write([]byte(longURL))
	}
	log.Println("Short URL created")
	return shortURL
}
