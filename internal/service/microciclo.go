package service

import (
	"atletismo-api/internal/models"
	"atletismo-api/internal/storage"
	"strings"
)

type MicrocicloService struct {
	repo storage.MicrocicloRepository
}

func NewMicrocicloService(repo storage.MicrocicloRepository) *MicrocicloService {
	return &MicrocicloService{repo: repo}
}

func (s *MicrocicloService) Listar() []models.Microciclo {
	return s.repo.ListarMicrociclos()
}

func (s *MicrocicloService) Obtener(id int) (models.Microciclo, error) {
	micro, ok := s.repo.BuscarMicrocicloPorID(id)
	if !ok {
		return models.Microciclo{}, ErrNoEncontrado
	}
	return micro, nil
}

func (s *MicrocicloService) Crear(micro models.Microciclo) (models.Microciclo, error) {
	if err := validacionMicrociclo(micro); err != nil {
		return models.Microciclo{}, err
	}
	return s.repo.CrearMicrociclo(micro), nil
}

func (s *MicrocicloService) Actualizar(id int, micro models.Microciclo) (models.Microciclo, error) {
	if err := validacionMicrociclo(micro); err != nil {
		return models.Microciclo{}, err
	}
	actualizado, ok := s.repo.ActualizarMicrociclo(id, micro)
	if !ok {
		return models.Microciclo{}, ErrNoEncontrado
	}
	return actualizado, nil
}

func (s *MicrocicloService) Borrar(id int) error {
	if !s.repo.BorrarMicrociclo(id) {
		return ErrNoEncontrado
	}
	return nil
}

func validacionMicrociclo(micro models.Microciclo) error {
	if strings.TrimSpace(micro.EnfoqueEspecifico) == "" {
		return ErrEnfoqueEspecificoVacio
	}
	return nil
}
