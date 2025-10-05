package server

import (
	"linky/handlers"
	"linky/service"
	"log"
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

	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
