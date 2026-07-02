package storage

import (
	"context"
	"database/sql"

	"atletismo-api/internal/models"
	"atletismo-api/internal/storage/sqlcdb"
)

// AlmacenSQLC implementa la interfaz Almacen usando código generado por sqlc
// (SQL escrito a mano + tipado generado) sobre database/sql.
//
// Es el TERCER backend de la cafetería, hermano de Memoria y AlmacenSQLite.
// El Server y los handlers no se enteran de cuál reciben: todos cumplen Almacen.
//
// Diferencias con sqlc que el adaptador tiene que resolver:
//  1. Los métodos generados piden context.Context  -> lo inyectamos acá dentro.
//  2. sqlc devuelve sus propios structs (int64)     -> los MAPEAMOS a models (int).
//  3. sqlc devuelve (T, error)                       -> lo absorbemos a (T, bool).
type AlmacenSQLC struct {
	q *sqlcdb.Queries
}

// NuevoAlmacenSQLC envuelve una conexión *sql.DB ya abierta.
func NuevoAlmacenSQLC(db *sql.DB) *AlmacenSQLC {
	return &AlmacenSQLC{q: sqlcdb.New(db)}
}

// =========================================================
// MAPEO sqlc <-> dominio (la "capa anticorrupción")
// =========================================================

func aCicloDominio(ci sqlcdb.CicloEntrenamiento) models.CicloEntrenamiento {
	return models.CicloEntrenamiento{
		ID:     int(ci.ID),
		Estado: ci.Estado,
		//FechaInicio: int(ci.FechaInicio),
		AtletaID: int(ci.AtletaID),
	}
}

func aEvaliacionDominio(eva sqlcdb.EvaluacionCiclo) models.EvaluacionCiclo {
	return models.EvaluacionCiclo{
		ID:                   int(eva.ID),
		CicloEntrenamientoID: int(eva.CicloEntrenamientoID),
		NivelFatiga:          int(eva.NivelFatiga),
		Comentarios:          eva.Comentarios,
		//FechaEvaluacion:      int(eva.FechaEvaluacion),
	}
}

func aMicrocicloDominio(micro sqlcdb.Microciclo) models.Microciclo {
	return models.Microciclo{
		ID:                   int(micro.ID),
		CicloEntrenamientoID: int(micro.CicloEntrenamientoID),
		NumeroSemana:         int(micro.NumeroSemana),
		EnfoqueEspecifico:    micro.EnfoqueEspecifico,
		//FechaInicio:          int(p.FechaInicio),
	}
}

// =========================================================
// CICLO ENTRENAMIENTO
// =========================================================

func (a *AlmacenSQLC) ListarCiclos() []models.CicloEntrenamiento {
	filas, err := a.q.ListarCiclos(context.Background())
	if err != nil {
		return nil
	}
	out := make([]models.CicloEntrenamiento, 0, len(filas))
	for _, f := range filas {
		out = append(out, aCicloDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) BuscarCicloPorID(id int) (models.CicloEntrenamiento, bool) {
	f, err := a.q.BuscarCicloPorID(context.Background(), int64(id))
	if err != nil {
		// Absorbemos sql.ErrNoRows (y cualquier otro error) y conservamos la firma comma-ok.
		return models.CicloEntrenamiento{}, false
	}
	return aCicloDominio(f), true
}

func (a *AlmacenSQLC) CrearCiclo(ci models.CicloEntrenamiento) models.CicloEntrenamiento {
	f, err := a.q.CrearCicloEntrenamiento(context.Background(), sqlcdb.CrearCicloEntrenamientoParams{
		Estado: ci.Estado,
		//FechaInicio: int64(ci.FechaInicio),
		AtletaID: int64(ci.AtletaID),
	})
	if err != nil {
		// La interfaz no permite reportar el fallo de una creación (igual que Memoria
		// y AlmacenSQLite). Devolvemos el zero value. Ver nota en la guía docente.
		return models.CicloEntrenamiento{}
	}
	return aCicloDominio(f)
}

func (a *AlmacenSQLC) ActualizarCiclo(id int, datos models.CicloEntrenamiento) (models.CicloEntrenamiento, bool) {
	f, err := a.q.ActualizarCicloEntrenamiento(context.Background(), sqlcdb.ActualizarCicloEntrenamientoParams{
		Estado: datos.Estado,
		//FechaInicio: int64(datos.FechaInicio),
		AtletaID: int64(datos.AtletaID),
		ID:       int64(id),
	})
	if err != nil {
		return models.CicloEntrenamiento{}, false
	}
	return aCicloDominio(f), true
}

func (a *AlmacenSQLC) BorrarCiclo(id int) bool {
	filas, err := a.q.BorrarCiclo(context.Background(), int64(id))
	if err != nil {
		return false
	}
	return filas > 0
}

// =========================================================
// EVALUACION ENTRENAMIENTO
// =========================================================

func (a *AlmacenSQLC) ListarEvaluaciones() []models.EvaluacionCiclo {
	filas, err := a.q.ListarEvaluacionCiclo(context.Background())
	if err != nil {
		return nil
	}
	out := make([]models.EvaluacionCiclo, 0, len(filas))
	for _, f := range filas {
		out = append(out, aEvaliacionDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) BuscarEvaluacionPorID(id int) (models.EvaluacionCiclo, bool) {
	f, err := a.q.BuscarEvaluacionCicloPorID(context.Background(), int64(id))
	if err != nil {
		return models.EvaluacionCiclo{}, false
	}
	return aEvaliacionDominio(f), true
}

func (a *AlmacenSQLC) CrearEvaluacion(eva models.EvaluacionCiclo) models.EvaluacionCiclo {
	f, err := a.q.CrearEvaluacionCiclo(context.Background(), sqlcdb.CrearEvaluacionCicloParams{
		CicloEntrenamientoID: int64(eva.CicloEntrenamientoID),
		NivelFatiga:          int64(eva.NivelFatiga),
		Comentarios:          eva.Comentarios,
	})
	if err != nil {
		return models.EvaluacionCiclo{}
	}
	return aEvaliacionDominio(f)
}

func (a *AlmacenSQLC) ActualizarEvaluacion(id int, datos models.EvaluacionCiclo) (models.EvaluacionCiclo, bool) {
	f, err := a.q.ActualizarEvaluacionCiclo(context.Background(), sqlcdb.ActualizarEvaluacionCicloParams{
		CicloEntrenamientoID: int64(datos.CicloEntrenamientoID),
		NivelFatiga:          int64(datos.NivelFatiga),
		Comentarios:          datos.Comentarios,
		ID:                   int64(id),
	})
	if err != nil {
		return models.EvaluacionCiclo{}, false
	}
	return aEvaliacionDominio(f), true
}

func (a *AlmacenSQLC) BorrarEvaluacion(id int) bool {
	filas, err := a.q.BorrarEvaluacionCiclo(context.Background(), int64(id))
	if err != nil {
		return false
	}
	return filas > 0
}

// =========================================================
// MICROCICLO
// =========================================================

func (a *AlmacenSQLC) ListarMicrociclos() []models.Microciclo {
	filas, err := a.q.ListarMicrociclo(context.Background())
	if err != nil {
		return nil
	}
	out := make([]models.Microciclo, 0, len(filas))
	for _, f := range filas {
		out = append(out, aMicrocicloDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) BuscarMicrocicloPorID(id int) (models.Microciclo, bool) {
	f, err := a.q.BuscarMicrocicloPorID(context.Background(), int64(id))
	if err != nil {
		return models.Microciclo{}, false
	}
	return aMicrocicloDominio(f), true
}

func (a *AlmacenSQLC) CrearMicrociclo(micro models.Microciclo) models.Microciclo {
	f, err := a.q.CrearMicrociclo(context.Background(), sqlcdb.CrearMicrocicloParams{
		CicloEntrenamientoID: int64(micro.CicloEntrenamientoID),
		NumeroSemana:         int64(micro.NumeroSemana),
		EnfoqueEspecifico:    micro.EnfoqueEspecifico,
	})
	if err != nil {
		return models.Microciclo{}
	}
	return aMicrocicloDominio(f)
}

func (a *AlmacenSQLC) ActualizarMicrociclo(id int, datos models.Microciclo) (models.Microciclo, bool) {
	f, err := a.q.ActualizarMicrociclo(context.Background(), sqlcdb.ActualizarMicrocicloParams{
		CicloEntrenamientoID: int64(datos.CicloEntrenamientoID),
		NumeroSemana:         int64(datos.NumeroSemana),
		EnfoqueEspecifico:    datos.EnfoqueEspecifico,
		ID:                   int64(id),
	})
	if err != nil {
		return models.Microciclo{}, false
	}
	return aMicrocicloDominio(f), true
}

func (a *AlmacenSQLC) BorrarMicrociclo(id int) bool {
	filas, err := a.q.BorrarMicrociclo(context.Background(), int64(id))
	if err != nil {
		return false
	}
	return filas > 0
}

// Chequeo en tiempo de compilación: AlmacenSQLC debe cumplir Almacen.
var _ Almacen = (*AlmacenSQLC)(nil)
