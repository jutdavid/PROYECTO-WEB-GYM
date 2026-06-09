package handlers

import (
	"encoding/json"
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
