package models

import "time"

// CicloEntrenamiento representa un ciclo de entrenamiento asignado a un atleta.
//
// AtletaID referencia el ID de un Atleta por número.
type CicloEntrenamiento struct {
	ID          int       `json:"id"`
	AtletaID    int       `json:"atleta_id"`
	Estado      string    `json:"estado"`
	FechaInicio time.Time `json:"fecha_inicio"`
}
