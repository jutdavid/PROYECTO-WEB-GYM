package models

import "time"

// MetricaFisica representa el progreso corporal del atleta en el tiempo.
// AtletaID referencia al Atleta evaluado.
type MetricaFisica struct {
	ID              int       `json:"id"`
	AtletaID        int       `json:"atleta_id"`
	PorcentajeGrasa float64   `json:"porcentaje_grasa"`
	MasaMuscular    float64   `json:"masa_muscular"`
	FechaMedicion   time.Time `json:"fecha_medicion"`
}
