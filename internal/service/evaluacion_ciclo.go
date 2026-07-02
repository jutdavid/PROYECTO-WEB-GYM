package service

import (
	"atletismo-api/internal/models"
	"atletismo-api/internal/storage"
	"strings"
)

type EvaluacionCicloService struct {
	repo storage.EvaluacionCicloRepository
}

func NewEvaluacionCicloService(repo storage.EvaluacionCicloRepository) *EvaluacionCicloService {
	return &EvaluacionCicloService{repo: repo}
}

func (s *EvaluacionCicloService) Listar() []models.EvaluacionCiclo {
	return s.repo.ListarEvaluaciones()
}

func (s *EvaluacionCicloService) Obtener(id int) (models.EvaluacionCiclo, error) {
	eva, ok := s.repo.BuscarEvaluacionPorID(id)
	if !ok {
		return models.EvaluacionCiclo{}, ErrNoEncontrado
	}
	return eva, nil
}

func (s *EvaluacionCicloService) Crear(eva models.EvaluacionCiclo) (models.EvaluacionCiclo, error) {
	if err := validacionEvaluacionCiclo(eva); err != nil {
		return models.EvaluacionCiclo{}, err
	}
	return s.repo.CrearEvaluacion(eva), nil
}

func (s *EvaluacionCicloService) Actualizar(id int, eva models.EvaluacionCiclo) (models.EvaluacionCiclo, error) {
	if err := validacionEvaluacionCiclo(eva); err != nil {
		return models.EvaluacionCiclo{}, err
	}
	actualizado, ok := s.repo.ActualizarEvaluacion(id, eva)
	if !ok {
		return models.EvaluacionCiclo{}, ErrNoEncontrado
	}
	return actualizado, nil
}

func (s *EvaluacionCicloService) Borrar(id int) error {
	if !s.repo.BorrarEvaluacion(id) {
		return ErrNoEncontrado
	}
	return nil
}

func validacionEvaluacionCiclo(eva models.EvaluacionCiclo) error {
	if strings.TrimSpace(eva.Comentarios) == "" {
		return ErrComentarioVacio
	}
	return nil
}
