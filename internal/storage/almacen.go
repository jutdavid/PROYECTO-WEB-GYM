package storage

import "atletismo-api/internal/models"

type CicloEntrenamientoRepository interface {
	ListarCiclos() []models.CicloEntrenamiento
	BuscarCicloPorID(id int) (models.CicloEntrenamiento, bool)
	CrearCiclo(ci models.CicloEntrenamiento) models.CicloEntrenamiento
	ActualizarCiclo(id int, datos models.CicloEntrenamiento) (models.CicloEntrenamiento, bool)
	BorrarCiclo(id int) bool
}

type EvaluacionCicloRepository interface {
	ListarEvaluaciones() []models.EvaluacionCiclo
	BuscarEvaluacionPorID(id int) (models.EvaluacionCiclo, bool)
	CrearEvaluacion(eva models.EvaluacionCiclo) models.EvaluacionCiclo
	ActualizarEvaluacion(id int, datos models.EvaluacionCiclo) (models.EvaluacionCiclo, bool)
	BorrarEvaluacion(id int) bool
}

type MicrocicloRepository interface {
	ListarMicrociclos() []models.Microciclo
	BuscarMicrocicloPorID(id int) (models.Microciclo, bool)
	CrearMicrociclo(micro models.Microciclo) models.Microciclo
	ActualizarMicrociclo(id int, datos models.Microciclo) (models.Microciclo, bool)
	BorrarMicrociclo(id int) bool
}

type UserRepository interface {
	CrearUsuario(u models.Usuario) (models.Usuario, error)
	BuscarUsuarioPorEmail(email string) (models.Usuario, bool)
}

type Almacen interface {
	// CicloEntrenamiento
	CicloEntrenamientoRepository
	// EvaluacionCiclo
	EvaluacionCicloRepository
	// Microciclo
	MicrocicloRepository
}

// Chequeo en tiempo de compilación: si Memoria dejara de cumplir Almacen,
// el proyecto NO compila. Red de seguridad opcional.
var _ Almacen = (*Memoria)(nil)
