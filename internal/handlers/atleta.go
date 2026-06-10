package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"

	"atletismo-api/internal/models"
	"atletismo-api/internal/storage"
)

type Server struct {
	Storage *storage.Memoria
}

func NewServer(s *storage.Memoria) *Server {
	return &Server{Storage: s}
}

func (s *Server) ListarAtletas(w http.ResponseWriter, _ *http.Request) {
	atletas := s.Storage.ListarAtletas()
	ResponderJSON(w, http.StatusOK, atletas)
}

func (s *Server) ObtenerAtleta(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		ResponderError(w, http.StatusBadRequest, "id debe ser un numero entero positivo")
		return
	}
	atleta, encontrado := s.Storage.BuscarAtletaPorID(uint(id))
	if !encontrado {
		ResponderError(w, http.StatusNotFound, "atleta no encontrado")
		return
	}
	ResponderJSON(w, http.StatusOK, atleta)
}

func (s *Server) CrearAtleta(w http.ResponseWriter, r *http.Request) {
	var nuevo models.Atleta
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		ResponderError(w, http.StatusBadRequest, "JSON invalido: "+err.Error())
		return
	}
	if strings.TrimSpace(nuevo.Nombre) == "" {
		ResponderError(w, http.StatusBadRequest, "el campo nombre es obligatorio")
		return
	}
	if nuevo.Peso < 0 {
		ResponderError(w, http.StatusBadRequest, "el peso no puede ser negativo")
		return
	}
	creado := s.Storage.CrearAtleta(nuevo)
	ResponderJSON(w, http.StatusCreated, creado)
}

func (s *Server) ActualizarAtleta(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		ResponderError(w, http.StatusBadRequest, "id debe ser un numero entero positivo")
		return
	}
	var datos models.Atleta
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		ResponderError(w, http.StatusBadRequest, "JSON invalido: "+err.Error())
		return
	}
	if strings.TrimSpace(datos.Nombre) == "" {
		ResponderError(w, http.StatusBadRequest, "el campo nombre es obligatorio")
		return
	}
	if datos.Peso < 0 {
		ResponderError(w, http.StatusBadRequest, "el peso no puede ser negativo")
		return
	}
	actualizado, encontrado := s.Storage.ActualizarAtleta(uint(id), datos)
	if !encontrado {
		ResponderError(w, http.StatusNotFound, "atleta no encontrado")
		return
	}
	ResponderJSON(w, http.StatusOK, actualizado)
}
