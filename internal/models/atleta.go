package models

type Atleta struct {
	ID                  int             `json:"id" gorm:"primaryKey"`
	Nombre              string          `json:"nombre"`
	MetodologiaObjetivo string          `json:"metodologia_objetivo"`
	Peso                float64         `json:"peso"`
	CoachID             int             `json:"coach_id"`
	Metricas            []MetricaFisica `json:"metricas" gorm:"foreignKey:AtletaID"`
	Lesiones            []Lesion        `json:"lesiones" gorm:"foreignKey:AtletaID"`
}