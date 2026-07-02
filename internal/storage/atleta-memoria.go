package storage

import (
	"atletismo-api/internal/models"
)

func (m *Memoria) SeedAtletas() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.atletas = []models.Atleta{
		{ID: 1, Nombre: "Pedro Gómez", MetodologiaObjetivo: "Sprint 100m", Peso: 72.5, CoachID: 1},
		{ID: 2, Nombre: "María Torres", MetodologiaObjetivo: "Maratón", Peso: 58.0, CoachID: 2},
		{ID: 3, Nombre: "José Ruiz", MetodologiaObjetivo: "Salto de longitud", Peso: 80.0, CoachID: 1},
		{ID: 4, Nombre: "Laura Vega", MetodologiaObjetivo: "Lanzamiento de bala", Peso: 90.0, CoachID: 3},
		{ID: 5, Nombre: "Andrés León", MetodologiaObjetivo: "5000m", Peso: 65.0, CoachID: 2},
	}
	m.nextAtletaID = 6
}

func (m *Memoria) ListarAtletas() []models.Atleta {
	m.mu.RLock()
	defer m.mu.RUnlock()

	copia := make([]models.Atleta, len(m.atletas))
	copy(copia, m.atletas)
	return copia
}

func (m *Memoria) BuscarAtletaPorID(id int) (models.Atleta, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, a := range m.atletas {
		if a.ID == id {
			return a, true
		}
	}
	return models.Atleta{}, false
}

func (m *Memoria) CrearAtleta(a models.Atleta) models.Atleta {
	m.mu.Lock()
	defer m.mu.Unlock()

	a.ID = m.nextAtletaID
	m.nextAtletaID++
	m.atletas = append(m.atletas, a)
	return a
}

func (m *Memoria) ActualizarAtleta(id int, datos models.Atleta) (models.Atleta, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, a := range m.atletas {
		if a.ID == id {
			datos.ID = id
			m.atletas[i] = datos
			return datos, true
		}
	}
	return models.Atleta{}, false
}

func (m *Memoria) BorrarAtleta(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, a := range m.atletas {
		if a.ID == id {
			m.atletas = append(m.atletas[:i], m.atletas[i+1:]...)
			return true
		}
	}
	return false
}
