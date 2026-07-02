package models

// Lesion mantiene un registro de las condiciones físicas limitantes del atleta.
// AtletaID referencia al Atleta afectado.
type Lesion struct {
	ID          int    `json:"id"`
	AtletaID    int    `json:"atleta_id"`
	Descripcion string `json:"descripcion"`
	Gravedad    string `json:"gravedad"`
	Estado      string `json:"estado"`
}
