package calculadorahandler

import (
	"encoding/json"
	"fmt"
	"main/modelos"
	"net/http"
	"time"
)

var historialOperaciones []modelos.OperacionGuardada

func HandleCalculate(w http.ResponseWriter, r *http.Request) {
	var request modelos.OperationRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var response modelos.OperationResponse

	switch request.Operator {
	case "+":
		response.Result = request.Number1 + request.Number2
	case "-":
		response.Result = request.Number1 - request.Number2
	case "*":
		response.Result = request.Number1 * request.Number2
	case "/":
		if request.Number2 == 0 {
			response.Error = "No se puede dividir por 0"
		} else {
			response.Result = request.Number1 / request.Number2
		}
	default:
		response.Error = "Operador inválido: " + request.Operator
	}

	if response.Error == "" {
		operation := fmt.Sprintf("%.2f %s %.2f = %.2f", request.Number1, request.Operator, request.Number2, response.Result)
		record := modelos.OperacionGuardada{Operation: operation, Result: response.Result, Date: time.Now()}
		historialOperaciones = append(historialOperaciones, record)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func HandleHistory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(historialOperaciones)
}
