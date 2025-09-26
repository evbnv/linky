package handlers

import (
	"encoding/json"
	"linky/models"
	"linky/service"
	"log"
	"net/http"
	"strings"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) GetLongURLHandler(w http.ResponseWriter, r *http.Request) {
	shortPath := strings.TrimPrefix(r.URL.Path, "/")
	longPath, err := h.service.GetLongURL(shortPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	log.Println("Long URL got")
	http.Redirect(w, r, longPath, http.StatusPermanentRedirect)
}

func (h *Handler) MakeShortURLHandler(w http.ResponseWriter, r *http.Request) {
	var req models.ShortenRequest
	var resp models.ShortenResponse
	var shortURL string

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	longURL := req.URL
	shortURL = h.service.URLTransform(longURL)

	resp = models.ShortenResponse{ShortURL: shortURL}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
