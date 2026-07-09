package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"atletismo-api/internal/models"
)

func parseID(r *http.Request) (int, bool) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return 0, false
	}
	return id, true
}

// ════════════════════════════ ATLETAS ═════════════════════════════════════

func (s *Server) ListarAtletas(w http.ResponseWriter, r *http.Request) {
	ResponderJSON(w, http.StatusOK, s.Atleta.ListarAtletas())
}

func (s *Server) ObtenerAtleta(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		ResponderError(w, http.StatusBadRequest, "id debe ser un número entero positivo")
		return
	}
	a, err := s.Atleta.ObtenerAtleta(id)
	if err != nil {
		ResponderError(w, statusDeError(err), err.Error())
		return
	}
	ResponderJSON(w, http.StatusOK, a)
}

func (s *Server) CrearAtleta(w http.ResponseWriter, r *http.Request) {
	var body models.Atleta
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ResponderError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	a, err := s.Atleta.CrearAtleta(body)
	if err != nil {
		ResponderError(w, statusDeError(err), err.Error())
		return
	}
	ResponderJSON(w, http.StatusCreated, a)
}

func (s *Server) ActualizarAtleta(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		ResponderError(w, http.StatusBadRequest, "id debe ser un número entero positivo")
		return
	}
	var body models.Atleta
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ResponderError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	a, err := s.Atleta.ActualizarAtleta(id, body)
	if err != nil {
		ResponderError(w, statusDeError(err), err.Error())
		return
	}
	ResponderJSON(w, http.StatusOK, a)
}

func (s *Server) BorrarAtleta(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		ResponderError(w, http.StatusBadRequest, "id debe ser un número entero positivo")
		return
	}
	if err := s.Atleta.BorrarAtleta(id); err != nil {
		ResponderError(w, statusDeError(err), err.Error())
		return
	}
	ResponderJSON(w, http.StatusNoContent, nil)
}

// ════════════════════════════ MÉTRICAS FÍSICAS ═════════════════════════════

func (s *Server) ListarMetricas(w http.ResponseWriter, r *http.Request) {
	ResponderJSON(w, http.StatusOK, s.Atleta.ListarMetricas())
}

func (s *Server) ObtenerMetrica(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		ResponderError(w, http.StatusBadRequest, "id debe ser un número entero positivo")
		return
	}
	m, err := s.Atleta.ObtenerMetrica(id)
	if err != nil {
		ResponderError(w, statusDeError(err), err.Error())
		return
	}
	ResponderJSON(w, http.StatusOK, m)
}

func (s *Server) CrearMetrica(w http.ResponseWriter, r *http.Request) {
	var body models.MetricaFisica
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ResponderError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	m, err := s.Atleta.CrearMetrica(body)
	if err != nil {
		ResponderError(w, statusDeError(err), err.Error())
		return
	}
	ResponderJSON(w, http.StatusCreated, m)
}

func (s *Server) ActualizarMetrica(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		ResponderError(w, http.StatusBadRequest, "id debe ser un número entero positivo")
		return
	}
	var body models.MetricaFisica
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ResponderError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	m, err := s.Atleta.ActualizarMetrica(id, body)
	if err != nil {
		ResponderError(w, statusDeError(err), err.Error())
		return
	}
	ResponderJSON(w, http.StatusOK, m)
}

func (s *Server) BorrarMetrica(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		ResponderError(w, http.StatusBadRequest, "id debe ser un número entero positivo")
		return
	}
	if err := s.Atleta.BorrarMetrica(id); err != nil {
		ResponderError(w, statusDeError(err), err.Error())
		return
	}
	ResponderJSON(w, http.StatusNoContent, nil)
}

// ════════════════════════════ LESIONES ═════════════════════════════════════

func (s *Server) ListarLesiones(w http.ResponseWriter, r *http.Request) {
	ResponderJSON(w, http.StatusOK, s.Atleta.ListarLesiones())
}

func (s *Server) ObtenerLesion(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		ResponderError(w, http.StatusBadRequest, "id debe ser un número entero positivo")
		return
	}
	l, err := s.Atleta.ObtenerLesion(id)
	if err != nil {
		ResponderError(w, statusDeError(err), err.Error())
		return
	}
	ResponderJSON(w, http.StatusOK, l)
}

func (s *Server) CrearLesion(w http.ResponseWriter, r *http.Request) {
	var body models.Lesion
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ResponderError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	l, err := s.Atleta.CrearLesion(body)
	if err != nil {
		ResponderError(w, statusDeError(err), err.Error())
		return
	}
	ResponderJSON(w, http.StatusCreated, l)
}

func (s *Server) ActualizarLesion(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		ResponderError(w, http.StatusBadRequest, "id debe ser un número entero positivo")
		return
	}
	var body models.Lesion
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ResponderError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	l, err := s.Atleta.ActualizarLesion(id, body)
	if err != nil {
		ResponderError(w, statusDeError(err), err.Error())
		return
	}
	ResponderJSON(w, http.StatusOK, l)
}

func (s *Server) BorrarLesion(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		ResponderError(w, http.StatusBadRequest, "id debe ser un número entero positivo")
		return
	}
	if err := s.Atleta.BorrarLesion(id); err != nil {
		ResponderError(w, statusDeError(err), err.Error())
		return
	}
	ResponderJSON(w, http.StatusNoContent, nil)
}