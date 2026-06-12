// Package models define las entidades del dominio atletismo.
package models

// Atleta representa un atleta registrado en el sistema.
//
// CoachID referencia el ID de un Entrenador por numero.
// Decision arquitectonica: usamos ID en lugar de struct anidado
// para facilitar la transicion a una base de datos relacional (GORM).
type Atleta struct {
	ID                  uint    `json:"id"`
	Nombre              string  `json:"nombre"`
	MetodologiaObjetivo string  `json:"metodologia_objetivo"`
	Peso                float64 `json:"peso"`
	CoachID             uint    `json:"coach_id"`
}
