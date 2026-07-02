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

// ListarCiclos atiende GET /api/v1/ciclos.
func (s *Server) ListarCiclos(w http.ResponseWriter, _ *http.Request) {
	ciclos := s.CicloEntrenamiento.Listar()
	ResponderJSON(w, http.StatusOK, ciclos)
}

// ObtenerCiclo atiende GET /api/v1/ciclos/{id}.
func (s *Server) ObtenerCiclo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		ResponderError(w, http.StatusBadRequest, "id debe ser un número entero positivo")
		return
	}

	ciclo, err := s.CicloEntrenamiento.Obtener(id)
	if err != nil {
		ResponderError(w, http.StatusNotFound, err.Error())
		return
	}

	ResponderJSON(w, http.StatusOK, ciclo)
}

// CrearCiclo atiende POST /api/v1/ciclos.
func (s *Server) CrearCiclo(w http.ResponseWriter, r *http.Request) {
	var nuevo models.CicloEntrenamiento
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		ResponderError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	creado, err := s.CicloEntrenamiento.Crear(nuevo)
	if err != nil {
		ResponderError(w, http.StatusBadRequest, err.Error())
		return
	}

	ResponderJSON(w, http.StatusCreated, creado)
}

// ActualizarCiclo atiende PUT /api/v1/ciclos/{id}.
func (s *Server) ActualizarCiclo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		ResponderError(w, http.StatusBadRequest, "id debe ser un número entero positivo")
		return
	}

	var datos models.CicloEntrenamiento
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		ResponderError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	actualizado, err := s.CicloEntrenamiento.Actualizar(id, datos)
	if err != nil {
		ResponderError(w, http.StatusBadRequest, err.Error())
		return
	}

	ResponderJSON(w, http.StatusOK, actualizado)
}

// BorrarCiclo atiende DELETE /api/v1/ciclos/{id}.
func (s *Server) BorrarCiclo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		ResponderError(w, http.StatusBadRequest, "id debe ser un número entero positivo")
		return
	}

	if err := s.CicloEntrenamiento.Borrar(id); err != nil {
		ResponderError(w, statusDeError(err), err.Error())
		return
	}

	ResponderJSON(w, http.StatusNoContent, nil)
}
