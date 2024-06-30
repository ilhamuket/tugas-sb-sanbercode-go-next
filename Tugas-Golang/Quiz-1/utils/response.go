package utils

import (
	"encoding/json"
	"net/http"
)

// ResponseFormatter formats a standard JSON response with status, message, and data.
func ResponseFormatter(w http.ResponseWriter, status int, message string, data interface{}) {
	response := map[string]interface{}{
		"status":  status,
		"message": message,
		"data":    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

// ResponseError sends an error response with a specified status code and message.
func ResponseError(w http.ResponseWriter, status int, message string) {
	ResponseFormatter(w, status, message, nil)
}
