package handlers

import (
	"encoding/json"
	"fmt"
	"linky/store"
	"net/http"
	"strconv"
)

type ShortenRequest struct {
	URL string `json:"url"`
}

func HandleAddress(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":

	case "POST":
		var req ShortenRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if !store.MapCreated {
			store.URLs = make(map[string]string)
			fmt.Println("Created!")
		}
		store.MapCreated = true
		longUrl := req.URL

		for {
			shortUrl := store.URLTransform(longUrl)

			if _, ok := store.URLs[shortUrl]; !ok {
				store.URLs[shortUrl] = longUrl
				break
			}

			longUrl = longUrl + ":" + strconv.Itoa(store.Count)
			store.Count++
		}

		fmt.Println(store.URLs)

		w.WriteHeader(http.StatusCreated)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
