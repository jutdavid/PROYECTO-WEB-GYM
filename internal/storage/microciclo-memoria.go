package storage

import (
	"atletismo-api/internal/models"
	"time"
)

func (m *Memoria) SeedMicrociclos() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.microciclos = []models.Microciclo{
		{ID: 1, CicloEntrenamientoID: 1, NumeroSemana: 1, EnfoqueEspecifico: "Fuerza Maxima", FechaInicio: time.Now()},
	}
	m.nextMicrocicloID = 2
}

func (m *Memoria) ListarMicrociclos() []models.Microciclo {
	m.mu.RLock()
	defer m.mu.RUnlock()

	copia := make([]models.Microciclo, len(m.microciclos))
	copy(copia, m.microciclos)
	return copia
}

func (m *Memoria) BuscarMicrocicloPorID(id int) (models.Microciclo, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, mc := range m.microciclos {
		if mc.ID == id {
			return mc, true
		}
	}
	return models.Microciclo{}, false
}

func (m *Memoria) CrearMicrociclo(micro models.Microciclo) models.Microciclo {
	m.mu.Lock()
	defer m.mu.Unlock()

	micro.ID = m.nextMicrocicloID
	m.nextMicrocicloID++
	m.microciclos = append(m.microciclos, micro)
	return micro
}

func (m *Memoria) ActualizarMicrociclo(id int, datos models.Microciclo) (models.Microciclo, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, mc := range m.microciclos {
		if mc.ID == id {
			datos.ID = id
			m.microciclos[i] = datos
			return datos, true
		}
	}
	return models.Microciclo{}, false
}

func (m *Memoria) BorrarMicrociclo(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, mc := range m.microciclos {
		if mc.ID == id {
			m.microciclos = append(m.microciclos[:i], m.microciclos[i+1:]...)
			return true
		}
	}
	return false
}
