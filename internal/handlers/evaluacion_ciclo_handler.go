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

// ListarEvaluaciones atiende GET /api/v1/evaluaciones
func (s *Server) ListarEvaluaciones(w http.ResponseWriter, r *http.Request) {
	evaluaciones := s.EvaluacionCiclo.Listar()
	ResponderJSON(w, http.StatusOK, evaluaciones)
}

// ObtenerEvaluacion atiende GET /api/v1/evaluaciones/{id}
func (s *Server) ObtenerEvaluacion(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id <= 0 {
		ResponderError(w, http.StatusBadRequest, "id debe ser un numero entero positivo")
		return
	}
	eval, err := s.EvaluacionCiclo.Obtener(id)
	if err != nil {
		ResponderError(w, http.StatusNotFound, err.Error())
		return
	}
	ResponderJSON(w, http.StatusOK, eval)
}

// CrearEvaluacion atiende POST /api/v1/evaluaciones
func (s *Server) CrearEvaluacion(w http.ResponseWriter, r *http.Request) {
	var nueva models.EvaluacionCiclo
	if err := json.NewDecoder(r.Body).Decode(&nueva); err != nil {
		ResponderError(w, http.StatusBadRequest, "JSON invalido: "+err.Error())
		return
	}

	creado, err := s.EvaluacionCiclo.Crear(nueva)
	if err != nil {
		ResponderError(w, http.StatusBadRequest, err.Error())
		return
	}

	ResponderJSON(w, http.StatusCreated, creado)
}

// BorrarEvaluacion atiende DELETE /api/v1/evaluaciones/{id}
func (s *Server) BorrarEvaluacion(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id <= 0 {
		ResponderError(w, http.StatusBadRequest, "id debe ser un numero entero positivo")
		return
	}
	if err := s.EvaluacionCiclo.Borrar(id); err != nil {
		ResponderError(w, statusDeError(err), err.Error())
		return
	}
	ResponderJSON(w, http.StatusNoContent, nil)
}
