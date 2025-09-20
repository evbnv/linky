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
type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

func HandleAddress(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":

	case "POST":
		var req ShortenRequest
		var resp ShortenResponse
		var shortUrl string
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
			shortUrl = store.URLTransform(longUrl)

			if _, ok := store.URLs[shortUrl]; !ok {
				store.URLs[shortUrl] = longUrl
				break
			}

			longUrl = longUrl + ":" + strconv.Itoa(store.Count)
			store.Count++
		}
		//
		fmt.Println(store.URLs)
		//

		resp = ShortenResponse{ShortURL: shortUrl}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(resp)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
