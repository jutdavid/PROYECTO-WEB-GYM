package models

type Atleta struct {
	ID                  uint    `json:"id"`
	Nombre              string  `json:"nombre"`
	MetodologiaObjetivo string  `json:"metodologia_objetivo"`
	Peso                float64 `json:"peso"`
}
