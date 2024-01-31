package modelos

import "time"

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
