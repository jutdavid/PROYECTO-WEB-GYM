package storage

import (
	"atletismo-api/internal/models"
	"time"
)

func (m *Memoria) SeedCertificaciones() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.certificaciones = []models.Certificacion{
		{ID: 1, EntrenadorID: 1, Nombre: "Entrenador IAAF Nivel 2", Institucion: "World Athletics", FechaEmision: time.Now()},
	}
	m.nextCertID = 2
}

func (m *Memoria) ListarCertificaciones() []models.Certificacion {
	m.mu.RLock()
	defer m.mu.RUnlock()

	copia := make([]models.Certificacion, len(m.certificaciones))
	copy(copia, m.certificaciones)
	return copia
}

func (m *Memoria) BuscarCertificacionPorID(id int) (models.Certificacion, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, c := range m.certificaciones {
		if c.ID == id {
			return c, true
		}
	}
	return models.Certificacion{}, false
}

func (m *Memoria) CrearCertificacion(cert models.Certificacion) models.Certificacion {
	m.mu.Lock()
	defer m.mu.Unlock()

	cert.ID = m.nextCertID
	m.nextCertID++
	m.certificaciones = append(m.certificaciones, cert)
	return cert
}

func (m *Memoria) ActualizarCertificacion(id int, datos models.Certificacion) (models.Certificacion, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, c := range m.certificaciones {
		if c.ID == id {
			datos.ID = id
			m.certificaciones[i] = datos
			return datos, true
		}
	}
	return models.Certificacion{}, false
}

func (m *Memoria) BorrarCertificacion(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, c := range m.certificaciones {
		if c.ID == id {
			m.certificaciones = append(m.certificaciones[:i], m.certificaciones[i+1:]...)
			return true
		}
	}
	return false
}
