package storage

import (
	"time"

	"gorm.io/gorm"

	"atletismo-api/internal/models"
)

// AlmacenSQLite implementa la interfaz Almacen usando GORM sobre SQLite.
//
// Fíjense: los métodos tienen EXACTAMENTE las mismas firmas que los de Memoria.
// Por eso el Server y los handlers no se enteran de cuál de los dos reciben.
type AlmacenSQLite struct {
	db *gorm.DB
}

// NuevoAlmacenSQLite envuelve una conexión *gorm.DB ya abierta.
func NuevoAlmacenSQLite(db *gorm.DB) *AlmacenSQLite {
	return &AlmacenSQLite{db: db}
}

// =========================================================
// CICLO ENTRENAMIENTO
// =========================================================

func (a *AlmacenSQLite) ListarCiclos() []models.CicloEntrenamiento {
	var ciclos []models.CicloEntrenamiento
	a.db.Find(&ciclos)
	return ciclos
}

func (a *AlmacenSQLite) BuscarCicloPorID(id int) (models.CicloEntrenamiento, bool) {
	var p models.CicloEntrenamiento
	if err := a.db.First(&p, id).Error; err != nil {
		// Absorbemos el error de la DB y conservamos la firma comma-ok.
		return models.CicloEntrenamiento{}, false
	}
	return p, true
}

func (a *AlmacenSQLite) CrearCiclo(p models.CicloEntrenamiento) models.CicloEntrenamiento {
	a.db.Create(&p) // GORM rellena el ID autogenerado en &p
	return p
}

func (a *AlmacenSQLite) ActualizarCiclo(id int, datos models.CicloEntrenamiento) (models.CicloEntrenamiento, bool) {
	var existente models.CicloEntrenamiento
	if err := a.db.First(&existente, id).Error; err != nil {
		return models.CicloEntrenamiento{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarCiclo(id int) bool {
	res := a.db.Delete(&models.CicloEntrenamiento{}, id)
	return res.RowsAffected > 0
}

// =========================================================
// EVALUACION CICLO
// =========================================================

func (a *AlmacenSQLite) ListarEvaluaciones() []models.EvaluacionCiclo {
	var evaluaciones []models.EvaluacionCiclo
	a.db.Find(&evaluaciones)
	return evaluaciones
}

func (a *AlmacenSQLite) BuscarEvaluacionPorID(id int) (models.EvaluacionCiclo, bool) {
	var c models.EvaluacionCiclo
	if err := a.db.First(&c, id).Error; err != nil {
		return models.EvaluacionCiclo{}, false
	}
	return c, true
}

func (a *AlmacenSQLite) CrearEvaluacion(c models.EvaluacionCiclo) models.EvaluacionCiclo {
	a.db.Create(&c)
	return c
}

func (a *AlmacenSQLite) ActualizarEvaluacion(id int, datos models.EvaluacionCiclo) (models.EvaluacionCiclo, bool) {
	var existente models.EvaluacionCiclo
	if err := a.db.First(&existente, id).Error; err != nil {
		return models.EvaluacionCiclo{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarEvaluacion(id int) bool {
	res := a.db.Delete(&models.EvaluacionCiclo{}, id)
	return res.RowsAffected > 0
}

// =========================================================
// MICROCICLOS
// =========================================================

func (a *AlmacenSQLite) ListarMicrociclos() []models.Microciclo {
	var microciclos []models.Microciclo
	a.db.Find(&microciclos)
	return microciclos
}

func (a *AlmacenSQLite) BuscarMicrocicloPorID(id int) (models.Microciclo, bool) {
	var p models.Microciclo
	if err := a.db.First(&p, id).Error; err != nil {
		// Absorbemos el error de la DB y conservamos la firma comma-ok.
		return models.Microciclo{}, false
	}
	return p, true
}

func (a *AlmacenSQLite) CrearMicrociclo(p models.Microciclo) models.Microciclo {
	a.db.Create(&p) // GORM rellena el ID autogenerado en &p
	return p
}

func (a *AlmacenSQLite) ActualizarMicrociclo(id int, datos models.Microciclo) (models.Microciclo, bool) {
	var existente models.Microciclo
	if err := a.db.First(&existente, id).Error; err != nil {
		return models.Microciclo{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarMicrociclo(id int) bool {
	res := a.db.Delete(&models.Microciclo{}, id)
	return res.RowsAffected > 0
}

// =========================================================
// SEEDS
// =========================================================

// SembrarSiVacio inserta datos iniciales solo si aún no hay categorías.
// Así no duplicamos datos en cada arranque del servidor.
func (a *AlmacenSQLite) SembrarSiVacio() {
	var n int64
	a.db.Model(&models.CicloEntrenamiento{}).Count(&n)
	if n > 0 {
		return
	}

	ciclos := []models.CicloEntrenamiento{
		{ID: 1, AtletaID: 1, Estado: "activo"},
		{ID: 2, AtletaID: 2, Estado: "completado"},
		{ID: 3, AtletaID: 3, Estado: "activo"},
	}
	a.db.Create(&ciclos)

	evaluaciones := []models.EvaluacionCiclo{
		{ID: 1, CicloEntrenamientoID: 1, NivelFatiga: 3, Comentarios: "Buen asimilamiento de las cargas de velocidad", FechaEvaluacion: time.Now()},
		//{ID: 2, CicloEntrenamientoID: 1, NivelFatiga: 4, Comentarios: "Ligeramente elevado", FechaEvaluacion: time.Now()},
	}
	a.db.Create(&evaluaciones)

	microciclos := []models.Microciclo{
		{ID: 1, CicloEntrenamientoID: 1, NumeroSemana: 1, EnfoqueEspecifico: "Fuerza Maxima", FechaInicio: time.Now()},
	}
	a.db.Create(&microciclos)
}

// Chequeo en tiempo de compilación: AlmacenSQLite debe cumplir Almacen.
var _ Almacen = (*AlmacenSQLite)(nil)
