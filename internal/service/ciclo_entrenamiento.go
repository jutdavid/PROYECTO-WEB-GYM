package service

import (
	"atletismo-api/internal/models"
	"atletismo-api/internal/storage"
	"strings"
)

type CicloEntrenamientoService struct {
	repo storage.CicloEntrenamientoRepository
}

func NewCicloEntrenamientoService(repo storage.CicloEntrenamientoRepository) *CicloEntrenamientoService {
	return &CicloEntrenamientoService{repo: repo}
}

func (s *CicloEntrenamientoService) Listar() []models.CicloEntrenamiento {
	return s.repo.ListarCiclos()
}

func (s *CicloEntrenamientoService) Obtener(id int) (models.CicloEntrenamiento, error) {
	ci, ok := s.repo.BuscarCicloPorID(id)
	if !ok {
		return models.CicloEntrenamiento{}, ErrNoEncontrado
	}
	return ci, nil
}

func (s *CicloEntrenamientoService) Crear(ci models.CicloEntrenamiento) (models.CicloEntrenamiento, error) {
	if err := validacionCicloEntrenamiento(ci); err != nil {
		return models.CicloEntrenamiento{}, err
	}
	return s.repo.CrearCiclo(ci), nil
}

func (s *CicloEntrenamientoService) Actualizar(id int, ci models.CicloEntrenamiento) (models.CicloEntrenamiento, error) {
	if err := validacionCicloEntrenamiento(ci); err != nil {
		return models.CicloEntrenamiento{}, err
	}
	actualizado, ok := s.repo.ActualizarCiclo(id, ci)
	if !ok {
		return models.CicloEntrenamiento{}, ErrNoEncontrado
	}
	return actualizado, nil
}

func (s *CicloEntrenamientoService) Borrar(id int) error {
	if !s.repo.BorrarCiclo(id) {
		return ErrNoEncontrado
	}
	return nil
}

func validacionCicloEntrenamiento(ci models.CicloEntrenamiento) error {
	if strings.TrimSpace(ci.Estado) == "" {
		return ErrEstadoVacio
	}
	return nil
}
