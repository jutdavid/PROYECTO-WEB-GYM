package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"

	"atletismo-api/internal/models"
)

func (s *Server) ListarEntrenadores(w http.ResponseWriter, _ *http.Request) {
	entrenadores := s.Storage.ListarEntrenadores()
	ResponderJSON(w, http.StatusOK, entrenadores)
}

func (s *Server) ObtenerEntrenador(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		ResponderError(w, http.StatusBadRequest, "id debe ser un número entero positivo")
		return
	}

	entrenador, encontrado := s.Storage.BuscarEntrenadorPorID(uint(id))
	if !encontrado {
		ResponderError(w, http.StatusNotFound, "entrenador no encontrado")
		return
	}

	ResponderJSON(w, http.StatusOK, entrenador)
}

func (s *Server) CrearEntrenador(w http.ResponseWriter, r *http.Request) {
	var nuevo models.Entrenador
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		ResponderError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if strings.TrimSpace(nuevo.Nombre) == "" {
		ResponderError(w, http.StatusBadRequest, "el campo nombre es obligatorio")
		return
	}
	if nuevo.CapacidadMaxima <= 0 {
		ResponderError(w, http.StatusBadRequest, "capacidad_maxima debe ser mayor a 0")
		return
	}
	if nuevo.CargaActual < 0 {
		ResponderError(w, http.StatusBadRequest, "carga_actual no puede ser negativa")
		return
	}

	creado := s.Storage.CrearEntrenador(nuevo)
	ResponderJSON(w, http.StatusCreated, creado)
}

func (s *Server) ActualizarEntrenador(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		ResponderError(w, http.StatusBadRequest, "id debe ser un número entero positivo")
		return
	}

	var datos models.Entrenador
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		ResponderError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if strings.TrimSpace(datos.Nombre) == "" {
		ResponderError(w, http.StatusBadRequest, "el campo nombre es obligatorio")
		return
	}
	if datos.CapacidadMaxima <= 0 {
		ResponderError(w, http.StatusBadRequest, "capacidad_maxima debe ser mayor a 0")
		return
	}
	if datos.CargaActual < 0 {
		ResponderError(w, http.StatusBadRequest, "carga_actual no puede ser negativa")
		return
	}

	actualizado, encontrado := s.Storage.ActualizarEntrenador(uint(id), datos)
	if !encontrado {
		ResponderError(w, http.StatusNotFound, "entrenador no encontrado")
		return
	}

	ResponderJSON(w, http.StatusOK, actualizado)
}

func (s *Server) BorrarEntrenador(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		ResponderError(w, http.StatusBadRequest, "id debe ser un número entero positivo")
		return
	}

	if !s.Storage.BorrarEntrenador(uint(id)) {
		ResponderError(w, http.StatusNotFound, "entrenador no encontrado")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
