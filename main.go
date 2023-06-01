package main

import (
	"log"

	"github.com/asevketicli/twittor/bd"
	"github.com/asevketicli/twittor/handlers"
)

func main() {

	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}

	handlers.Manejadores()
}
