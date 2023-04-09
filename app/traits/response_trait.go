package traits

import (
	"encoding/json"
	"net/http"
)

type ResponseTrait struct{}

type ErrorResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (t *ResponseTrait) RespondWithFailure(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := ErrorResponse{
		Status:  false,
		Message: message,
	}

	json.NewEncoder(w).Encode(response)
}

func (t *ResponseTrait) RespondWithSuccess(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := ErrorResponse{
		Status:  true,
		Message: message,
		Data:    data,
	}

	json.NewEncoder(w).Encode(response)
}
