package main

import (
	"log"
	"main/controller"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Configuración de enrutador Gorilla Mux
	r := mux.NewRouter()

	// Manejadores para la calculadora y el historial de operaciones
	controller.RegisterHandlers(r)

	// Configuración de CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "POST"},
	})

	// Configura el manejador HTTP con CORS habilitado
	handler := corsHandler.Handler(r)

	// Inicia el servidor HTTP con Gorilla Mux y CORS habilitado
	log.Fatal(http.ListenAndServe(":9000", handler))
}
