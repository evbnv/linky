package server

import (
	"linky/handlers"
	"linky/service"
	"net/http"
)

func InitRoutes(s *service.Service) {
	h := handlers.NewHandler(s)

	http.HandleFunc("/api/shorten", h.MakeShortURLHandler)
	http.HandleFunc("/", h.GetLongURLHandler)

	http.ListenAndServe(":8080", nil)
}
