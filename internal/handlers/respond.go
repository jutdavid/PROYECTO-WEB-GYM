package handlers

import (
	"atletismo-api/internal/service"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func ResponderJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data == nil {
		return
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

func ResponderError(w http.ResponseWriter, status int, mensaje string) {
	ResponderJSON(w, status, map[string]string{"error": mensaje})
}

/////////////////////////////////////////////////

func statusDeError(err error) int {
	switch {
	case errors.Is(err, service.ErrNoEncontrado):
		return http.StatusBadRequest
	case errors.Is(err, service.ErrCredencialesInvalidas):
		return http.StatusUnauthorized
	case errors.Is(err, service.ErrEmailEnUso):
		return http.StatusConflict
	case errors.Is(err, service.ErrNombreVacio):
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError

	}
}
