package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/rs/cors"
)

type OperationRequest struct {
	Number1  float64 `json:"number1"`
	Number2  float64 `json:"number2"`
	Operator string  `json:"operator"`
}

type OperationResponse struct {
	Result float64 `json:"result"`
	Error  string  `json:"error,omitempty"`
}

type OperacionGuardada struct {
	Operation string    `json:"operation"`
	Result    float64   `json:"result"`
	Date      time.Time `json:"date"`
}

var historialOperaciones []OperacionGuardada // Mover la variable fuera de la función main para que sea global

func main() {
	// Configuración de CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "POST"},
	})

	// Configura el manejador HTTP
	handler := corsHandler.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/calculate":
			if r.Method == http.MethodPost {
				decoder := json.NewDecoder(r.Body)
				var request OperationRequest
				err := decoder.Decode(&request)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}

				var response OperationResponse
				switch request.Operator {
				case "+":
					response.Result = request.Number1 + request.Number2
				case "-":
					response.Result = request.Number1 - request.Number2
				case "*":
					response.Result = request.Number1 * request.Number2
				case "/":
					if request.Number2 == 0 {
						response.Error = "No se puede div por 0"
					} else {
						response.Result = request.Number1 / request.Number2
					}
				default:
					response.Error = "Operador inválido: " + request.Operator
				}

				// Guardar la operación en el historial
				if response.Error == "" {
					operation := fmt.Sprintf("%.2f %s %.2f = %.2f", request.Number1, request.Operator, request.Number2, response.Result)
					record := OperacionGuardada{Operation: operation, Result: response.Result, Date: time.Now()}
					historialOperaciones = append(historialOperaciones, record)

					// Imprimir operación guardada
					fmt.Printf("Operación: %s - Resultado: %.2f - Fecha: %s\n", operation, response.Result, time.Now().Format("2006-01-02 15:04:05"))
				}

				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(response)
			} else {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
		case "/history":
			if r.Method == http.MethodGet {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(historialOperaciones)
			} else {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
		default:
			http.NotFound(w, r)
		}
	}))

	// Inicia el servidor HTTP con CORS habilitado
	log.Fatal(http.ListenAndServe(":9000", handler))
}
