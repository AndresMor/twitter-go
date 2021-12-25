package main

import (
	"Documents/Go/twitter-go/bd"
	"Documents/Go/twitter-go/handlers"
	"log"
)

func main() {
	if bd.CheckConnection() == false {
		log.Fatal("Sin conexión")
	}
	handlers.Handlers()
}
