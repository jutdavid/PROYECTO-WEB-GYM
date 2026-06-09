package models

type Entrenador struct {
	ID              uint    `json:"id"`
	Nombre          string  `json:"nombre"`
	Especialidad    string  `json:"especialidad"`
	CapacidadMaxima int     `json:"capacidad_maxima"`
	CargaActual     float64 `json:"carga_actual"`
}
