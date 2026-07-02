package storage

import (
	"atletismo-api/internal/models"
)

func (m *Memoria) SeedHorarios() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.horarios = []models.HorarioAtencion{
		{ID: 1, EntrenadorID: 1, DiaSemana: "Lunes", HoraInicio: "08:00", HoraFin: "12:00"},
	}
	m.nextHorarioID = 2
}

func (m *Memoria) ListarHorarios() []models.HorarioAtencion {
	m.mu.RLock()
	defer m.mu.RUnlock()

	copia := make([]models.HorarioAtencion, len(m.horarios))
	copy(copia, m.horarios)
	return copia
}

func (m *Memoria) BuscarHorarioPorID(id int) (models.HorarioAtencion, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, h := range m.horarios {
		if h.ID == id {
			return h, true
		}
	}
	return models.HorarioAtencion{}, false
}

func (m *Memoria) CrearHorario(horario models.HorarioAtencion) models.HorarioAtencion {
	m.mu.Lock()
	defer m.mu.Unlock()

	horario.ID = m.nextHorarioID
	m.nextHorarioID++
	m.horarios = append(m.horarios, horario)
	return horario
}

func (m *Memoria) ActualizarHorario(id int, datos models.HorarioAtencion) (models.HorarioAtencion, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, h := range m.horarios {
		if h.ID == id {
			datos.ID = id
			m.horarios[i] = datos
			return datos, true
		}
	}
	return models.HorarioAtencion{}, false
}

func (m *Memoria) BorrarHorario(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, h := range m.horarios {
		if h.ID == id {
			m.horarios = append(m.horarios[:i], m.horarios[i+1:]...)
			return true
		}
	}
	return false
}
