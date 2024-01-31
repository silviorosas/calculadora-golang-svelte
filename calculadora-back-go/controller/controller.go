package controller

import (
	"main/calculadorahandler"

	"github.com/gorilla/mux"
)

func RegisterHandlers(r *mux.Router) {
	// Manejadores para la calculadora
	r.HandleFunc("/calculate", calculadorahandler.HandleCalculate).Methods("POST")

	r.HandleFunc("/history", calculadorahandler.HandleHistory).Methods("GET")
}
