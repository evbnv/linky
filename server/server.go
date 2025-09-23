package server

import (
	"linky/database"
	"linky/envs"
	"log"
)

func InitServer() {
	// Load envs
	errEnvs := envs.LoadEnvs()
	if errEnvs != nil {
		// Вывод сообщения об ошибке
		log.Fatal("Ошибка инициализации ENV: ", errEnvs)
	} else {
		log.Println("Инициализация ENV прошла успешно")
	}
	// Init database
	errDatabase := database.InitDatabase()
	if errDatabase != nil {
		log.Fatal("Ошибка подключения к базе данных: ", errDatabase)
	} else {
		log.Println("Успешное подключение к базе данных")
	}
}

func StartServer() {
	InitRoutes()
}
