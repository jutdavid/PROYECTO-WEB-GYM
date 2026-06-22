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

// ListarCertificaciones atiende GET /api/v1/certificaciones
func (s *Server) ListarCertificaciones(w http.ResponseWriter, r *http.Request) {
	certificaciones := s.Storage.ListarCertificaciones()
	ResponderJSON(w, http.StatusOK, certificaciones)
}

// ObtenerCertificacion atiende GET /api/v1/certificaciones/{id}
func (s *Server) ObtenerCertificacion(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id <= 0 {
		ResponderError(w, http.StatusBadRequest, "id debe ser un numero entero positivo")
		return
	}
	cert, encontrado := s.Storage.BuscarCertificacionPorID(id)
	if !encontrado {
		ResponderError(w, http.StatusNotFound, "certificacion no encontrada")
		return
	}
	ResponderJSON(w, http.StatusOK, cert)
}

// CrearCertificacion atiende POST /api/v1/certificaciones
func (s *Server) CrearCertificacion(w http.ResponseWriter, r *http.Request) {
	var nueva models.Certificacion
	if err := json.NewDecoder(r.Body).Decode(&nueva); err != nil {
		ResponderError(w, http.StatusBadRequest, "JSON invalido: "+err.Error())
		return
	}
	if nueva.EntrenadorID <= 0 {
		ResponderError(w, http.StatusBadRequest, "el campo entrenador_id es obligatorio")
		return
	}
	if strings.TrimSpace(nueva.Nombre) == "" {
		ResponderError(w, http.StatusBadRequest, "el campo nombre es obligatorio")
		return
	}
	if nueva.FechaEmision.IsZero() {
		nueva.FechaEmision = time.Now()
	}

	creado := s.Storage.CrearCertificacion(nueva)
	ResponderJSON(w, http.StatusCreated, creado)
}

// BorrarCertificacion atiende DELETE /api/v1/certificaciones/{id}
func (s *Server) BorrarCertificacion(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id <= 0 {
		ResponderError(w, http.StatusBadRequest, "id debe ser un numero entero positivo")
		return
	}
	if !s.Storage.BorrarCertificacion(id) {
		ResponderError(w, http.StatusNotFound, "certificacion no encontrada")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
