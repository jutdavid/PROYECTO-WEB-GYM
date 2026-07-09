package storage

import "atletismo-api/internal/models"

type AtletaRepository interface {
	ListarAtletas() []models.Atleta
	BuscarAtletaPorID(id int) (models.Atleta, bool)
	CrearAtleta(a models.Atleta) models.Atleta
	ActualizarAtleta(id int, datos models.Atleta) (models.Atleta, bool)
	BorrarAtleta(id int) bool
}

type MetricaFisicaRepository interface {
	ListarMetricas() []models.MetricaFisica
	BuscarMetricaPorID(id int) (models.MetricaFisica, bool)
	CrearMetrica(m models.MetricaFisica) models.MetricaFisica
	ActualizarMetrica(id int, datos models.MetricaFisica) (models.MetricaFisica, bool)
	BorrarMetrica(id int) bool
}

type LesionRepository interface {
	ListarLesiones() []models.Lesion
	BuscarLesionPorID(id int) (models.Lesion, bool)
	CrearLesion(l models.Lesion) models.Lesion
	ActualizarLesion(id int, datos models.Lesion) (models.Lesion, bool)
	BorrarLesion(id int) bool
}

type AlmacenAtleta interface {
	AtletaRepository
	MetricaFisicaRepository
	LesionRepository
}

var _ AlmacenAtleta = (*AlmacenSQLite)(nil)