package storage

import (
	"atletismo-api/internal/models"
)

func (m *Memoria) SeedCiclos() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.ciclos = []models.CicloEntrenamiento{
		{ID: 1, AtletaID: 1, Estado: "activo"},
		{ID: 2, AtletaID: 2, Estado: "completado"},
		{ID: 3, AtletaID: 3, Estado: "activo"},
	}
	m.nextCicloID = 4
}

func (m *Memoria) ListarCiclos() []models.CicloEntrenamiento {
	m.mu.RLock()
	defer m.mu.RUnlock()

	copia := make([]models.CicloEntrenamiento, len(m.ciclos))
	copy(copia, m.ciclos)
	return copia
}

func (m *Memoria) BuscarCicloPorID(id int) (models.CicloEntrenamiento, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, c := range m.ciclos {
		if c.ID == id {
			return c, true
		}
	}
	return models.CicloEntrenamiento{}, false
}

func (m *Memoria) CrearCiclo(c models.CicloEntrenamiento) models.CicloEntrenamiento {
	m.mu.Lock()
	defer m.mu.Unlock()

	c.ID = m.nextCicloID
	m.nextCicloID++
	m.ciclos = append(m.ciclos, c)
	return c
}

func (m *Memoria) ActualizarCiclo(id int, datos models.CicloEntrenamiento) (models.CicloEntrenamiento, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, c := range m.ciclos {
		if c.ID == id {
			datos.ID = id
			m.ciclos[i] = datos
			return datos, true
		}
	}
	return models.CicloEntrenamiento{}, false
}

func (m *Memoria) BorrarCiclo(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, c := range m.ciclos {
		if c.ID == id {
			m.ciclos = append(m.ciclos[:i], m.ciclos[i+1:]...)
			return true
		}
	}
	return false
}
