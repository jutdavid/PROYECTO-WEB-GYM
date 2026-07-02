package storage

import (
	"atletismo-api/internal/models"
)

func (m *Memoria) SeedLesiones() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.lesiones = []models.Lesion{
		{ID: 1, AtletaID: 1, Descripcion: "Esguince de tobillo", Gravedad: "Media", Estado: "Recuperacion"},
	}
	m.nextLesionID = 2
}

func (m *Memoria) ListarLesiones() []models.Lesion {
	m.mu.RLock()
	defer m.mu.RUnlock()

	copia := make([]models.Lesion, len(m.lesiones))
	copy(copia, m.lesiones)
	return copia
}

func (m *Memoria) BuscarLesionPorID(id int) (models.Lesion, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, l := range m.lesiones {
		if l.ID == id {
			return l, true
		}
	}
	return models.Lesion{}, false
}

func (m *Memoria) CrearLesion(lesion models.Lesion) models.Lesion {
	m.mu.Lock()
	defer m.mu.Unlock()

	lesion.ID = m.nextLesionID
	m.nextLesionID++
	m.lesiones = append(m.lesiones, lesion)
	return lesion
}

func (m *Memoria) ActualizarLesion(id int, datos models.Lesion) (models.Lesion, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, l := range m.lesiones {
		if l.ID == id {
			datos.ID = id
			m.lesiones[i] = datos
			return datos, true
		}
	}
	return models.Lesion{}, false
}

func (m *Memoria) BorrarLesion(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, l := range m.lesiones {
		if l.ID == id {
			m.lesiones = append(m.lesiones[:i], m.lesiones[i+1:]...)
			return true
		}
	}
	return false
}
