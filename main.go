package main

import (
	"linky/server"
	"log/slog"
	"os"
)

func main() {
	// logger init
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	server.RunServer()
}
