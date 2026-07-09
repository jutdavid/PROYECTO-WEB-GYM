package service

import (
	"strings"

	"atletismo-api/internal/models"
	"atletismo-api/internal/storage"
)

// AtletaService concentra las reglas de negocio del módulo de Atletas:
// la entidad principal (Atleta) y sus dos entidades relacionadas
// (MetricaFisica y Lesion).
type AtletaService struct {
	repo storage.AlmacenAtleta
}

func NewAtletaService(repo storage.AlmacenAtleta) *AtletaService {
	return &AtletaService{repo: repo}}

// ═══════════════════════════ ATLETAS ═════════════════════════════════════════

func (s *AtletaService) ListarAtletas() []models.Atleta {
	return s.repo.ListarAtletas()
}

func (s *AtletaService) ObtenerAtleta(id int) (models.Atleta, error) {
	a, ok := s.repo.BuscarAtletaPorID(id)
	if !ok {
		return models.Atleta{}, ErrAtletaNoEncontrado
	}
	return a, nil
}

func (s *AtletaService) CrearAtleta(a models.Atleta) (models.Atleta, error) {
	if err := validarAtleta(a); err != nil {
		return models.Atleta{}, err
	}
	return s.repo.CrearAtleta(a), nil
}

func (s *AtletaService) ActualizarAtleta(id int, datos models.Atleta) (models.Atleta, error) {
	if err := validarAtleta(datos); err != nil {
		return models.Atleta{}, err
	}
	a, ok := s.repo.ActualizarAtleta(id, datos)
	if !ok {
		return models.Atleta{}, ErrAtletaNoEncontrado
	}
	return a, nil
}

func (s *AtletaService) BorrarAtleta(id int) error {
	if !s.repo.BorrarAtleta(id) {
		return ErrAtletaNoEncontrado
	}
	return nil
}

func validarAtleta(a models.Atleta) error {
	if strings.TrimSpace(a.Nombre) == "" {
		return ErrNombreVacio
	}
	if a.Peso <= 0 {
		return ErrPesoInvalido
	}
	if a.CoachID == 0 {
		return ErrCoachIDInvalido
	}
	return nil
}

// ═══════════════════════════ MÉTRICAS FÍSICAS ════════════════════════════════

func (s *AtletaService) ListarMetricas() []models.MetricaFisica {
	return s.repo.ListarMetricas()
}

func (s *AtletaService) ObtenerMetrica(id int) (models.MetricaFisica, error) {
	m, ok := s.repo.BuscarMetricaPorID(id)
	if !ok {
		return models.MetricaFisica{}, ErrMetricaNoEncontrada
	}
	return m, nil
}

func (s *AtletaService) CrearMetrica(m models.MetricaFisica) (models.MetricaFisica, error) {
	if err := validarMetrica(m); err != nil {
		return models.MetricaFisica{}, err
	}
	return s.repo.CrearMetrica(m), nil
}

func (s *AtletaService) ActualizarMetrica(id int, datos models.MetricaFisica) (models.MetricaFisica, error) {
	if err := validarMetrica(datos); err != nil {
		return models.MetricaFisica{}, err
	}
	m, ok := s.repo.ActualizarMetrica(id, datos)
	if !ok {
		return models.MetricaFisica{}, ErrMetricaNoEncontrada
	}
	return m, nil
}

func (s *AtletaService) BorrarMetrica(id int) error {
	if !s.repo.BorrarMetrica(id) {
		return ErrMetricaNoEncontrada
	}
	return nil
}

func validarMetrica(m models.MetricaFisica) error {
	if m.AtletaID == 0 {
		return ErrAtletaIDInvalido
	}
	if m.PorcentajeGrasa <= 0 {
		return ErrPorcentajeGrasaInvalido
	}
	if m.MasaMuscular <= 0 {
		return ErrMasaMuscularInvalida
	}
	return nil
}

// ═══════════════════════════ LESIONES ════════════════════════════════════════

func (s *AtletaService) ListarLesiones() []models.Lesion {
	return s.repo.ListarLesiones()
}

func (s *AtletaService) ObtenerLesion(id int) (models.Lesion, error) {
	l, ok := s.repo.BuscarLesionPorID(id)
	if !ok {
		return models.Lesion{}, ErrLesionNoEncontrada
	}
	return l, nil
}

func (s *AtletaService) CrearLesion(l models.Lesion) (models.Lesion, error) {
	if err := validarLesion(l); err != nil {
		return models.Lesion{}, err
	}
	return s.repo.CrearLesion(l), nil
}

func (s *AtletaService) ActualizarLesion(id int, datos models.Lesion) (models.Lesion, error) {
	if err := validarLesion(datos); err != nil {
		return models.Lesion{}, err
	}
	l, ok := s.repo.ActualizarLesion(id, datos)
	if !ok {
		return models.Lesion{}, ErrLesionNoEncontrada
	}
	return l, nil
}

func (s *AtletaService) BorrarLesion(id int) error {
	if !s.repo.BorrarLesion(id) {
		return ErrLesionNoEncontrada
	}
	return nil
}

func validarLesion(l models.Lesion) error {
	if l.AtletaID == 0 {
		return ErrAtletaIDInvalido
	}
	if strings.TrimSpace(l.Descripcion) == "" {
		return ErrDescripcionVacia
	}
	if strings.TrimSpace(l.Gravedad) == "" {
		return ErrGravedadVacia
	}
	if strings.TrimSpace(l.Estado) == "" {
		return ErrEstadoLesionVacio
	}
	return nil
}