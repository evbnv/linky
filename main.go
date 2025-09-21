package main

import (
	"linky/handlers"
	"linky/store"
	"net/http"
)

func main() {
	http.HandleFunc("/api/", handlers.HandleAddress)
	store.InitStore()
	http.ListenAndServe(":8080", nil)
}
