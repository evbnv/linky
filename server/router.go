package server

import (
	"linky/handlers"
	"net/http"
)

func InitRoutes() {
	http.HandleFunc("/api/shorten", handlers.POST)
	http.HandleFunc("/", handlers.GET)

	http.ListenAndServe(":8080", nil)
}
