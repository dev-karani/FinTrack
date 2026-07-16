package httpx

import (
	"encoding/json"
	"log"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

func RespondWithError(w http.ResponseWriter, status int, msg string) {
	RespondWithJSON(w, status, errorResponse{Error: msg})
}

func RespondWithJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err :=json.NewEncoder(w).Encode(v); err !=nil {
		log.Printf("encode response: %v", err)
	}
}
