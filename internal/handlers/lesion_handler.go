package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"atletismo-api/internal/models"
)

// ListarLesiones atiende GET /api/v1/lesiones
func (s *Server) ListarLesiones(w http.ResponseWriter, r *http.Request) {
	lesiones := s.Storage.ListarLesiones()
	ResponderJSON(w, http.StatusOK, lesiones)
}

// ObtenerLesion atiende GET /api/v1/lesiones/{id}
func (s *Server) ObtenerLesion(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id <= 0 {
		ResponderError(w, http.StatusBadRequest, "id debe ser un numero entero positivo")
		return
	}
	lesion, encontrado := s.Storage.BuscarLesionPorID(id)
	if !encontrado {
		ResponderError(w, http.StatusNotFound, "lesion no encontrada")
		return
	}
	ResponderJSON(w, http.StatusOK, lesion)
}

// CrearLesion atiende POST /api/v1/lesiones
func (s *Server) CrearLesion(w http.ResponseWriter, r *http.Request) {
	var nueva models.Lesion
	if err := json.NewDecoder(r.Body).Decode(&nueva); err != nil {
		ResponderError(w, http.StatusBadRequest, "JSON invalido: "+err.Error())
		return
	}
	if nueva.AtletaID <= 0 {
		ResponderError(w, http.StatusBadRequest, "el campo atleta_id es obligatorio")
		return
	}
	if strings.TrimSpace(nueva.Descripcion) == "" || strings.TrimSpace(nueva.Estado) == "" {
		ResponderError(w, http.StatusBadRequest, "descripcion y estado son campos obligatorios")
		return
	}

	creado := s.Storage.CrearLesion(nueva)
	ResponderJSON(w, http.StatusCreated, creado)
}

// BorrarLesion atiende DELETE /api/v1/lesiones/{id}
func (s *Server) BorrarLesion(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id <= 0 {
		ResponderError(w, http.StatusBadRequest, "id debe ser un numero entero positivo")
		return
	}
	if !s.Storage.BorrarLesion(id) {
		ResponderError(w, http.StatusNotFound, "lesion no encontrada")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
