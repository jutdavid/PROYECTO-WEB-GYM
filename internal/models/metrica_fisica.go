package models

import "time"

type MetricaFisica struct {
	ID              int       `json:"id" gorm:"primaryKey"`
	AtletaID        int       `json:"atleta_id"`
	FechaMedicion   time.Time `json:"fecha_medicion"`
	PorcentajeGrasa float64   `json:"porcentaje_grasa"`
	MasaMuscular    float64   `json:"masa_muscular"`
	Atleta          *Atleta   `json:"atleta,omitempty" gorm:"foreignKey:AtletaID"`
}