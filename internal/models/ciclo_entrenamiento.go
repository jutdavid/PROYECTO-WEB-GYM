package models

import "time"

// CicloEntrenamiento representa un ciclo de entrenamiento asignado a un atleta.
//
// AtletaID referencia el ID de un Atleta por número.
type CicloEntrenamiento struct {
	ID          uint      `json:"id"`
	AtletaID    uint      `json:"atleta_id"`
	Estado      string    `json:"estado"`
	FechaInicio time.Time `json:"fecha_inicio"`
}
