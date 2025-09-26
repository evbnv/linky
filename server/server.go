package server

import (
	"linky/database"
	"linky/envs"
	"linky/service"
	"log"
)

func RunServer() {
	// Load envs
	errEnvs := envs.LoadEnvs()
	if errEnvs != nil {
		// Вывод сообщения об ошибке
		log.Fatal("Ошибка инициализации ENV: ", errEnvs)
	} else {
		log.Println("Инициализация ENV прошла успешно")
	}
	// Init database
	dbClient, errDatabase := database.InitDatabase()
	if errDatabase != nil {
		log.Fatal("Ошибка подключения к базе данных: ", errDatabase)
	} else {
		log.Println("Успешное подключение к базе данных")
	}
	// Init store, service, routes
	myStore := database.NewPostgresStore(dbClient)
	myService := service.NewService(myStore)
	InitRoutes(myService)
}
