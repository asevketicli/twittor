package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/asevketicli/twittor/middleW"
	"github.com/asevketicli/twittor/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* Manejadores seteo mi puerto, el Handler y pongo a escuchar al Serve */
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middleW.ChequeoBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
