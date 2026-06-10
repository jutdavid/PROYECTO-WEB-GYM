package handlers

import (
	"net/http"

	"atletismo-api/internal/storage"
)

// Server agrupa las dependencias compartidas por los handlers.
type Server struct {
	Storage *storage.Memoria
}

// NewServer construye un Server listo para usar.
func NewServer(s *storage.Memoria) *Server {
	return &Server{Storage: s}
}

// ListarAtletas atiende GET /api/v1/atletas.
func (s *Server) ListarAtletas(w http.ResponseWriter, _ *http.Request) {
	atletas := s.Storage.ListarAtletas()
	ResponderJSON(w, http.StatusOK, atletas)
}
