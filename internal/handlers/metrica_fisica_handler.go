package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"atletismo-api/internal/models"
)

// ListarMetricas atiende GET /api/v1/metricas
func (s *Server) ListarMetricas(w http.ResponseWriter, r *http.Request) {
	metricas := s.Storage.ListarMetricas()
	ResponderJSON(w, http.StatusOK, metricas)
}

// ObtenerMetrica atiende GET /api/v1/metricas/{id}
func (s *Server) ObtenerMetrica(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id <= 0 {
		ResponderError(w, http.StatusBadRequest, "id debe ser un numero entero positivo")
		return
	}
	metrica, encontrado := s.Storage.BuscarMetricaPorID(id)
	if !encontrado {
		ResponderError(w, http.StatusNotFound, "metrica no encontrada")
		return
	}
	ResponderJSON(w, http.StatusOK, metrica)
}

// CrearMetrica atiende POST /api/v1/metricas
func (s *Server) CrearMetrica(w http.ResponseWriter, r *http.Request) {
	var nueva models.MetricaFisica
	if err := json.NewDecoder(r.Body).Decode(&nueva); err != nil {
		ResponderError(w, http.StatusBadRequest, "JSON invalido: "+err.Error())
		return
	}
	if nueva.AtletaID <= 0 {
		ResponderError(w, http.StatusBadRequest, "el campo atleta_id es obligatorio")
		return
	}
	if nueva.PorcentajeGrasa < 0 || nueva.MasaMuscular < 0 {
		ResponderError(w, http.StatusBadRequest, "las metricas fisicas no pueden ser negativas")
		return
	}
	if nueva.FechaMedicion.IsZero() {
		nueva.FechaMedicion = time.Now()
	}

	creado := s.Storage.CrearMetrica(nueva)
	ResponderJSON(w, http.StatusCreated, creado)
}

// BorrarMetrica atiende DELETE /api/v1/metricas/{id}
func (s *Server) BorrarMetrica(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id <= 0 {
		ResponderError(w, http.StatusBadRequest, "id debe ser un numero entero positivo")
		return
	}
	if !s.Storage.BorrarMetrica(id) {
		ResponderError(w, http.StatusNotFound, "metrica no encontrada")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
