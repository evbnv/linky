package main

import (
	"linky/database"
	"linky/handlers"
	"log"
	"net/http"
)

func init() {
	// Init database
	errDatabase := database.InitDatabase()
	if errDatabase != nil {
		log.Fatal("Ошибка подключения к базе данных: ", errDatabase)
	} else {
		log.Println("Успешное подключение к базе данных")
	}
}

func main() {
	http.HandleFunc("/api/", handlers.HandleAddress)

	http.ListenAndServe(":8080", nil)
}
