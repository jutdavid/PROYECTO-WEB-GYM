package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

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
