package main

import (
	"linky/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/api/", handlers.HandleAddress)
	http.ListenAndServe(":8080", nil)
}
