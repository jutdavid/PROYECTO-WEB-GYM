package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"

	"atletismo-api/internal/models"
)

// ListarCiclos atiende GET /api/v1/ciclos.
func (s *Server) ListarCiclos(w http.ResponseWriter, _ *http.Request) {
	ciclos := s.Storage.ListarCiclos()
	ResponderJSON(w, http.StatusOK, ciclos)
}

// ObtenerCiclo atiende GET /api/v1/ciclos/{id}.
func (s *Server) ObtenerCiclo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		ResponderError(w, http.StatusBadRequest, "id debe ser un número entero positivo")
		return
	}

	ciclo, encontrado := s.Storage.BuscarCicloPorID(uint(id))
	if !encontrado {
		ResponderError(w, http.StatusNotFound, "ciclo de entrenamiento no encontrado")
		return
	}

	ResponderJSON(w, http.StatusOK, ciclo)
}