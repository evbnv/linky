package handlers

import (
	"encoding/json"
	"linky/service"
	"net/http"
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
		var shortURL string

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		longURL := req.URL
		shortURL = service.URLTransform(longURL)

		resp = ShortenResponse{ShortURL: shortURL}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(resp)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
