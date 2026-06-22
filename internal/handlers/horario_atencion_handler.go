package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"

	"atletismo-api/internal/models"
)

// ListarHorarios atiende GET /api/v1/horarios
func (s *Server) ListarHorarios(w http.ResponseWriter, r *http.Request) {
	horarios := s.Storage.ListarHorarios()
	ResponderJSON(w, http.StatusOK, horarios)
}

// ObtenerHorario atiende GET /api/v1/horarios/{id}
func (s *Server) ObtenerHorario(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id <= 0 {
		ResponderError(w, http.StatusBadRequest, "id debe ser un numero entero positivo")
		return
	}
	horario, encontrado := s.Storage.BuscarHorarioPorID(id)
	if !encontrado {
		ResponderError(w, http.StatusNotFound, "horario no encontrado")
		return
	}
	ResponderJSON(w, http.StatusOK, horario)
}

// CrearHorario atiende POST /api/v1/horarios
func (s *Server) CrearHorario(w http.ResponseWriter, r *http.Request) {
	var nuevo models.HorarioAtencion
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		ResponderError(w, http.StatusBadRequest, "JSON invalido: "+err.Error())
		return
	}
	if nuevo.EntrenadorID <= 0 {
		ResponderError(w, http.StatusBadRequest, "el campo entrenador_id es obligatorio")
		return
	}
	if strings.TrimSpace(nuevo.DiaSemana) == "" || strings.TrimSpace(nuevo.HoraInicio) == "" || strings.TrimSpace(nuevo.HoraFin) == "" {
		ResponderError(w, http.StatusBadRequest, "dia_semana, hora_inicio y hora_fin son obligatorios")
		return
	}

	creado := s.Storage.CrearHorario(nuevo)
	ResponderJSON(w, http.StatusCreated, creado)
}

// BorrarHorario atiende DELETE /api/v1/horarios/{id}
func (s *Server) BorrarHorario(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id <= 0 {
		ResponderError(w, http.StatusBadRequest, "id debe ser un numero entero positivo")
		return
	}
	if !s.Storage.BorrarHorario(id) {
		ResponderError(w, http.StatusNotFound, "horario no encontrado")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
