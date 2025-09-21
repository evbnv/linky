package store

import "log"

var URLs map[string]string

func InitStore() {
	URLs = make(map[string]string)
	log.Println("Store is initilized")
}
