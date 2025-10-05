package handlers

import (
	"encoding/json"
	"linky/models"
	"linky/service"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Handler struct {
	service *service.Service
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) GetLongURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path == "/" {
		h.serveIndexHandler(w, r)
		return
	}

	shortPath := strings.TrimPrefix(r.URL.Path, "/")
	longPath, err := h.service.GetLongURL(shortPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	log.Println("Long URL got")
	http.Redirect(w, r, longPath, http.StatusPermanentRedirect)
}

func (h *Handler) PostShortURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.ShortenRequest
	var resp models.ShortenResponse
	var shortURL string

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// URL validation
	if req.URL == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "URL cannot be empty"})
		return
	}

	parsedURL, err := url.Parse(req.URL)
	if err != nil || (parsedURL.Scheme != "http" && parsedURL.Scheme != "https") {
		if parsedURL, err = url.Parse("https://" + req.URL); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{Error: "Invalid URL. Please check the address for typos."})
			return
		}
	}

	// URL transformation
	shortURL = h.service.URLTransform(parsedURL.String())

	resp = models.ShortenResponse{ShortURL: shortURL}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) serveIndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/index.html")
}
