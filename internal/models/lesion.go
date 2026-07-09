package models

type Lesion struct {
	ID          int     `json:"id" gorm:"primaryKey"`
	AtletaID    int     `json:"atleta_id"`
	Descripcion string  `json:"descripcion"`
	Gravedad    string  `json:"gravedad"`
	Estado      string  `json:"estado"`
	Atleta      *Atleta `json:"atleta,omitempty" gorm:"foreignKey:AtletaID"`
}