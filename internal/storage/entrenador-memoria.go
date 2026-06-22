package storage

import "atletismo-api/internal/models"

func (m *Memoria) SeedEntrenadores() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.entrenadores = []models.Entrenador{
		{ID: 1, Nombre: "Carlos Mendoza", Especialidad: "Velocidad", CapacidadMaxima: 10, CargaActual: 2.5},
		{ID: 2, Nombre: "Ana Ríos", Especialidad: "Resistencia", CapacidadMaxima: 8, CargaActual: 3.0},
		{ID: 3, Nombre: "Luis Paredes", Especialidad: "Fuerza", CapacidadMaxima: 6, CargaActual: 1.5},
	}
	m.nextEntrenadorID = 4
}

func (m *Memoria) ListarEntrenadores() []models.Entrenador {
	m.mu.RLock()
	defer m.mu.RUnlock()

	copia := make([]models.Entrenador, len(m.entrenadores))
	copy(copia, m.entrenadores)
	return copia
}

func (m *Memoria) BuscarEntrenadorPorID(id int) (models.Entrenador, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, e := range m.entrenadores {
		if e.ID == id {
			return e, true
		}
	}
	return models.Entrenador{}, false
}

func (m *Memoria) CrearEntrenador(e models.Entrenador) models.Entrenador {
	m.mu.Lock()
	defer m.mu.Unlock()

	e.ID = m.nextEntrenadorID
	m.nextEntrenadorID++
	m.entrenadores = append(m.entrenadores, e)
	return e
}

func (m *Memoria) ActualizarEntrenador(id int, datos models.Entrenador) (models.Entrenador, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, e := range m.entrenadores {
		if e.ID == id {
			datos.ID = id
			m.entrenadores[i] = datos
			return datos, true
		}
	}
	return models.Entrenador{}, false
}

func (m *Memoria) BorrarEntrenador(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, e := range m.entrenadores {
		if e.ID == id {
			m.entrenadores = append(m.entrenadores[:i], m.entrenadores[i+1:]...)
			return true
		}
	}
	return false
}
