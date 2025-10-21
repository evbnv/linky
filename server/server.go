package server

import (
	"linky/database"
	"linky/envs"
	"linky/service"
	"log"
	"log/slog"
)

func RunServer() {
	// Load envs
	errEnvs := envs.LoadEnvs()
	if errEnvs != nil {
		log.Fatal("Error loading ENVs", "event", "envs_load_error", "error", errEnvs)
	} else {
		slog.Info("ENVs successfully loaded", "event", "envs_loaded")
	}
	// Init database
	dbClient, errDatabase := database.InitDatabase()
	if errDatabase != nil {
		log.Fatal("DB connection failed", errDatabase, "event", "db_connection_fail", "error", errDatabase)
	} else {
		slog.Info("DB connection success", "event", "db_connection_success")
	}
	// Init store, service, routes
	myStore := database.NewPostgresStore(dbClient)
	myService := service.NewService(myStore)
	InitRoutes(myService)
}
