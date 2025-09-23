package handlers

import (
	"encoding/json"
	"linky/service"
	"net/http"
	"strings"
)

type ShortenRequest struct {
	URL string `json:"url"`
}
type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

func GET(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getLongURLHandler(w, r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func POST(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		makeShortURLHandler(w, r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func getLongURLHandler(w http.ResponseWriter, r *http.Request) {
	shortPath := strings.TrimPrefix(r.URL.Path, "/")
	longPath, err := service.GetLongURL(shortPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	http.Redirect(w, r, longPath, http.StatusPermanentRedirect)
}

func makeShortURLHandler(w http.ResponseWriter, r *http.Request) {
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
}
