package storage

import (
	"atletismo-api/internal/models"
	"time"
)

func (m *Memoria) SeedEvaluaciones() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.evaluaciones = []models.EvaluacionCiclo{
		{ID: 1, CicloEntrenamientoID: 1, NivelFatiga: 3, Comentarios: "Buen asimilamiento de las cargas de velocidad", FechaEvaluacion: time.Now()},
		//{ID: 2, CicloEntrenamientoID: 1, NivelFatiga: 4, Comentarios: "Ligeramente elevado", FechaEvaluacion: time.Now()},
	}
	m.nextEvaluacionID = 2
}

func (m *Memoria) ListarEvaluaciones() []models.EvaluacionCiclo {
	m.mu.RLock()
	defer m.mu.RUnlock()

	copia := make([]models.EvaluacionCiclo, len(m.evaluaciones))
	copy(copia, m.evaluaciones)
	return copia
}

func (m *Memoria) BuscarEvaluacionPorID(id int) (models.EvaluacionCiclo, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, ev := range m.evaluaciones {
		if ev.ID == id {
			return ev, true
		}
	}
	return models.EvaluacionCiclo{}, false
}

func (m *Memoria) CrearEvaluacion(eval models.EvaluacionCiclo) models.EvaluacionCiclo {
	m.mu.Lock()
	defer m.mu.Unlock()

	eval.ID = m.nextEvaluacionID
	m.nextEvaluacionID++
	m.evaluaciones = append(m.evaluaciones, eval)
	return eval
}

func (m *Memoria) ActualizarEvaluacion(id int, datos models.EvaluacionCiclo) (models.EvaluacionCiclo, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, ev := range m.evaluaciones {
		if ev.ID == id {
			datos.ID = id
			m.evaluaciones[i] = datos
			return datos, true
		}
	}
	return models.EvaluacionCiclo{}, false
}

func (m *Memoria) BorrarEvaluacion(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, ev := range m.evaluaciones {
		if ev.ID == id {
			m.evaluaciones = append(m.evaluaciones[:i], m.evaluaciones[i+1:]...)
			return true
		}
	}
	return false
}
