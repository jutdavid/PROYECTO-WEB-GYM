package storage

import "atletismo-api/internal/models"

// ═══════════════════════════ ATLETAS ═════════════════════════════════════════

func (a *AlmacenSQLite) ListarAtletas() []models.Atleta {
	var lista []models.Atleta
	a.db.Find(&lista)
	return lista
}

func (s *AlmacenSQLite) BuscarAtletaPorID(id int) (models.Atleta, bool) {
	var atleta models.Atleta
	// Agregamos Preload para traer las relaciones automáticamente
	result := s.db.Preload("Metricas").Preload("Lesiones").First(&atleta, id)
	if result.Error != nil {
		return models.Atleta{}, false
	}
	return atleta, true
}

func (a *AlmacenSQLite) CrearAtleta(at models.Atleta) models.Atleta {
	a.db.Create(&at)
	return at
}

func (a *AlmacenSQLite) ActualizarAtleta(id int, datos models.Atleta) (models.Atleta, bool) {
	var existente models.Atleta
	if err := a.db.First(&existente, id).Error; err != nil {
		return models.Atleta{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarAtleta(id int) bool {
	res := a.db.Delete(&models.Atleta{}, id)
	return res.RowsAffected > 0
}

// ═══════════════════════════ MÉTRICAS FÍSICAS ════════════════════════════════

func (a *AlmacenSQLite) ListarMetricas() []models.MetricaFisica {
	var lista []models.MetricaFisica
	a.db.Find(&lista)
	return lista
}

func (s *AlmacenSQLite) BuscarMetricaPorID(id int) (models.MetricaFisica, bool) {
	var metrica models.MetricaFisica
	// Precargamos el Atleta asociado a esta métrica
	result := s.db.Preload("Atleta").First(&metrica, id)
	if result.Error != nil {
		return models.MetricaFisica{}, false
	}
	return metrica, true
}

func (a *AlmacenSQLite) CrearMetrica(m models.MetricaFisica) models.MetricaFisica {
	a.db.Create(&m)
	return m
}

func (a *AlmacenSQLite) ActualizarMetrica(id int, datos models.MetricaFisica) (models.MetricaFisica, bool) {
	var existente models.MetricaFisica
	if err := a.db.First(&existente, id).Error; err != nil {
		return models.MetricaFisica{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarMetrica(id int) bool {
	res := a.db.Delete(&models.MetricaFisica{}, id)
	return res.RowsAffected > 0
}

// ═══════════════════════════ LESIONES ════════════════════════════════════════

func (a *AlmacenSQLite) ListarLesiones() []models.Lesion {
	var lista []models.Lesion
	a.db.Find(&lista)
	return lista
}

func (s *AlmacenSQLite) BuscarLesionPorID(id int) (models.Lesion, bool) {
	var lesion models.Lesion
	// Precargamos el Atleta asociado a esta lesión
	result := s.db.Preload("Atleta").First(&lesion, id)
	if result.Error != nil {
		return models.Lesion{}, false
	}
	return lesion, true
}

func (a *AlmacenSQLite) CrearLesion(l models.Lesion) models.Lesion {
	a.db.Create(&l)
	return l
}

func (a *AlmacenSQLite) ActualizarLesion(id int, datos models.Lesion) (models.Lesion, bool) {
	var existente models.Lesion
	if err := a.db.First(&existente, id).Error; err != nil {
		return models.Lesion{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarLesion(id int) bool {
	res := a.db.Delete(&models.Lesion{}, id)
	return res.RowsAffected > 0
}