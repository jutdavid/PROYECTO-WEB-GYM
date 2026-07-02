package models

import "time"

// Microciclo representa una semana o bloque corto dentro de un ciclo mayor.
type Microciclo struct {
	ID                   int       `json:"id"`
	CicloEntrenamientoID int       `json:"ciclo_entrenamiento_id"`
	NumeroSemana         int       `json:"numero_semana"`
	EnfoqueEspecifico    string    `json:"enfoque_especifico"`
	FechaInicio          time.Time `json:"fecha_inicio"`
}
