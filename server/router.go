package server

import (
	"linky/handlers"
	"linky/service"
	"net/http"
)

func InitRoutes(s *service.Service) {
	h := handlers.NewHandler(s)

	mux := http.NewServeMux()

	mux.HandleFunc("/api/shorten", h.PostShortURLHandler)
	mux.HandleFunc("/", h.GetLongURLHandler)

	http.ListenAndServe(":8080", mux)
}
