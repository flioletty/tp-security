package handler

import (
	"context"
	"encoding/json"
	"net/http"
)

type ClientResponseDto struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}

func NewResponseDto(ctx context.Context, w http.ResponseWriter, statusCode int, message string, payload interface{}) {
	response := ClientResponseDto{
		Status:  statusCode,
		Message: message,
		Payload: payload,
	}
	sendData(ctx, w, response, statusCode, message)
}

func sendData(ctx context.Context, w http.ResponseWriter, response ClientResponseDto, statusCode int, message string) {
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to parse into JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(responseJSON)
}
