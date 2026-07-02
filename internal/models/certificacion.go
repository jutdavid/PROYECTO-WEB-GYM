package models

import "time"

// Certificacion representa un título, curso o aval técnico del entrenador.
// EntrenadorID referencia al Entrenador que posee la certificación.
type Certificacion struct {
	ID           int       `json:"id"`
	EntrenadorID int       `json:"entrenador_id"`
	Nombre       string    `json:"nombre"`
	Institucion  string    `json:"institucion"`
	FechaEmision time.Time `json:"fecha_emision"`
}
