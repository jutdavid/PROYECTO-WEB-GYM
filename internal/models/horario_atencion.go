package models

// HorarioAtencion define las franjas en las que el entrenador está disponible.
type HorarioAtencion struct {
	ID           int    `json:"id"`
	EntrenadorID int    `json:"entrenador_id"`
	DiaSemana    string `json:"dia_semana"`
	HoraInicio   string `json:"hora_inicio"`
	HoraFin      string `json:"hora_fin"`
}
