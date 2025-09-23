package main

import (
	"linky/server"
)

func init() {
	server.InitServer()
}

func main() {
	server.StartServer()
}
