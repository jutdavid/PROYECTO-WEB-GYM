package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	//"strings"
	//"time"

	"github.com/go-chi/chi/v5"

	"atletismo-api/internal/models"
)

// ListarMicrociclos atiende GET /api/v1/microciclos
func (s *Server) ListarMicrociclos(w http.ResponseWriter, r *http.Request) {
	microciclos := s.Microciclo.Listar()
	ResponderJSON(w, http.StatusOK, microciclos)
}

// ObtenerMicrociclo atiende GET /api/v1/microciclos/{id}
func (s *Server) ObtenerMicrociclo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id <= 0 {
		ResponderError(w, http.StatusBadRequest, "id debe ser un numero entero positivo")
		return
	}
	microciclo, err := s.Microciclo.Obtener(id)
	if err != nil {
		ResponderError(w, http.StatusNotFound, err.Error())
		return
	}
	ResponderJSON(w, http.StatusOK, microciclo)
}

// CrearMicrociclo atiende POST /api/v1/microciclos
func (s *Server) CrearMicrociclo(w http.ResponseWriter, r *http.Request) {
	var nuevo models.Microciclo
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		ResponderError(w, http.StatusBadRequest, "JSON invalido: "+err.Error())
		return
	}

	creado, err := s.Microciclo.Crear(nuevo)
	if err != nil {
		ResponderError(w, http.StatusBadRequest, err.Error())
		return
	}

	ResponderJSON(w, http.StatusCreated, creado)
}

// BorrarMicrociclo atiende DELETE /api/v1/microciclos/{id}
func (s *Server) BorrarMicrociclo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id <= 0 {
		ResponderError(w, http.StatusBadRequest, "id debe ser un numero entero positivo")
		return
	}
	if err := s.Microciclo.Borrar(id); err != nil {
		ResponderError(w, http.StatusNotFound, err.Error())
		return
	}

	ResponderJSON(w, http.StatusNoContent, nil)
}
