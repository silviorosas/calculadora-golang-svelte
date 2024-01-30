package main

import (
	"encoding/json"
	"log"
	"net/http"

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

				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(response)
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
