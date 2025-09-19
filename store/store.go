package store

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

var URLs map[string]string
var MapCreated bool
var Count int

func URLTransform(longURL string) string {
	h := sha256.New()
	var shortURL string

	h.Write([]byte(longURL))
	shortURL = base64.URLEncoding.EncodeToString([]byte(h.Sum(nil)[:8]))

	fmt.Println("shortURL", shortURL)

	return shortURL
}
