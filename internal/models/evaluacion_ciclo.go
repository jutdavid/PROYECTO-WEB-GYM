package models

import "time"

// EvaluacionCiclo es el feedback técnico al final de ciertas fases.
type EvaluacionCiclo struct {
	ID                   int       `json:"id"`
	CicloEntrenamientoID int       `json:"ciclo_entrenamiento_id"`
	NivelFatiga          int       `json:"nivel_fatiga"`
	Comentarios          string    `json:"comentarios"`
	FechaEvaluacion      time.Time `json:"fecha_evaluacion"`
}
