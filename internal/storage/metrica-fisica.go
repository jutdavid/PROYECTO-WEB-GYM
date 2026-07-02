package storage

import (
	"atletismo-api/internal/models"
	"time"
)

func (m *Memoria) SeedMetricas() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.metricasFisicas = []models.MetricaFisica{
		{ID: 1, AtletaID: 1, PorcentajeGrasa: 12.4, MasaMuscular: 61.2, FechaMedicion: time.Now()},
		{ID: 2, AtletaID: 2, PorcentajeGrasa: 18.5, MasaMuscular: 45.0, FechaMedicion: time.Now()},
	}
	m.nextMetricaID = 3
}

func (m *Memoria) ListarMetricas() []models.MetricaFisica {
	m.mu.RLock()
	defer m.mu.RUnlock()

	copia := make([]models.MetricaFisica, len(m.metricasFisicas))
	copy(copia, m.metricasFisicas)
	return copia
}

func (m *Memoria) BuscarMetricaPorID(id int) (models.MetricaFisica, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, mf := range m.metricasFisicas {
		if mf.ID == id {
			return mf, true
		}
	}
	return models.MetricaFisica{}, false
}

func (m *Memoria) CrearMetrica(metrica models.MetricaFisica) models.MetricaFisica {
	m.mu.Lock()
	defer m.mu.Unlock()

	metrica.ID = m.nextMetricaID
	m.nextMetricaID++
	m.metricasFisicas = append(m.metricasFisicas, metrica)
	return metrica
}

func (m *Memoria) ActualizarMetrica(id int, datos models.MetricaFisica) (models.MetricaFisica, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, mf := range m.metricasFisicas {
		if mf.ID == id {
			datos.ID = id
			m.metricasFisicas[i] = datos
			return datos, true
		}
	}
	return models.MetricaFisica{}, false
}

func (m *Memoria) BorrarMetrica(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, mf := range m.metricasFisicas {
		if mf.ID == id {
			m.metricasFisicas = append(m.metricasFisicas[:i], m.metricasFisicas[i+1:]...)
			return true
		}
	}
	return false
}
