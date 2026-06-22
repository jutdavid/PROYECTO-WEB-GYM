package storage

import (
	"sync"

	"atletismo-api/internal/models"
)

// Memoria es un almacén unificado del sistema de atletismo.
type Memoria struct {
	entrenadores     []models.Entrenador
	nextEntrenadorID int
	atletas          []models.Atleta
	nextAtletaID     int
	ciclos           []models.CicloEntrenamiento
	nextCicloID      int
	metricasFisicas  []models.MetricaFisica
	nextMetricaID    int
	lesiones         []models.Lesion
	nextLesionID     int
	certificaciones  []models.Certificacion
	nextCertID       int
	horarios         []models.HorarioAtencion
	nextHorarioID    int
	microciclos      []models.Microciclo
	nextMicrocicloID int
	evaluaciones     []models.EvaluacionCiclo
	nextEvaluacionID int

	mu sync.RWMutex
}

// NuevaMemoria inicializa el almacén central con slices vacíos.
func NuevaMemoria() *Memoria {
	return &Memoria{
		entrenadores:     []models.Entrenador{},
		nextEntrenadorID: 1,
		atletas:          []models.Atleta{},
		nextAtletaID:     1,
		ciclos:           []models.CicloEntrenamiento{},
		nextCicloID:      1,
		metricasFisicas:  []models.MetricaFisica{},
		nextMetricaID:    1,
		lesiones:         []models.Lesion{},
		nextLesionID:     1,
		certificaciones:  []models.Certificacion{},
		nextCertID:       1,
		horarios:         []models.HorarioAtencion{},
		nextHorarioID:    1,
		microciclos:      []models.Microciclo{},
		nextMicrocicloID: 1,
		evaluaciones:     []models.EvaluacionCiclo{},
		nextEvaluacionID: 1,
	}
}
