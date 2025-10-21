package server

import (
	"linky/handlers"
	"linky/service"
	"log/slog"
	"net/http"
)

func InitRoutes(s *service.Service) {
	h := handlers.NewHandler(s)

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("web"))
	mux.Handle("/css/", fileServer)
	mux.Handle("/js/", fileServer)

	mux.HandleFunc("/api/shorten", h.PostShortURLHandler)
	mux.HandleFunc("/", h.GetLongURLHandler)

	slog.Info("Server listening on :8080", "event", "server_started")
	http.ListenAndServe(":8080", mux)
}
