package storage

import (
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"atletismo-api/internal/models"
)

func nuevoDBAtletaPrueba(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err, "no se pudo abrir SQLite en memoria")

	err = db.AutoMigrate(&models.Atleta{}, &models.MetricaFisica{}, &models.Lesion{})
	require.NoError(t, err, "AutoMigrate falló")

	return db
}

func TestAlmacenSQLite_CrearYBuscarAtleta(t *testing.T) {
	db := nuevoDBAtletaPrueba(t)
	almacen := NuevoAlmacenSQLite(db)

	creado := almacen.CrearAtleta(models.Atleta{
		Nombre:              "Eliud Kipchoge",
		MetodologiaObjetivo: "Maratón",
		Peso:                57.5,
		CoachID:             1,
	})
	assert.NotZero(t, creado.ID, "GORM debía asignar un ID autogenerado")

	encontrado, ok := almacen.BuscarAtletaPorID(creado.ID)
	require.True(t, ok, "se esperaba encontrar el atleta recién creado")
	assert.Equal(t, "Eliud Kipchoge", encontrado.Nombre)

	_, ok = almacen.BuscarAtletaPorID(9999)
	assert.False(t, ok, "no debía encontrarse un atleta con ID inexistente")
}

// Demuestra la relación Has-Many/Belongs-To funcionando de extremo a extremo.
func TestAlmacenSQLite_RelacionAtletaConMetricasYLesiones(t *testing.T) {
	db := nuevoDBAtletaPrueba(t)
	almacen := NuevoAlmacenSQLite(db)

	atleta := almacen.CrearAtleta(models.Atleta{Nombre: "María Torres", Peso: 58, CoachID: 2})

	almacen.CrearMetrica(models.MetricaFisica{AtletaID: atleta.ID, PorcentajeGrasa: 14, MasaMuscular: 40})
	almacen.CrearLesion(models.Lesion{AtletaID: atleta.ID, Descripcion: "Esguince", Gravedad: "leve", Estado: "activa"})

	encontrado, ok := almacen.BuscarAtletaPorID(atleta.ID)
	require.True(t, ok)
	require.Len(t, encontrado.Metricas, 1, "Has-Many: el atleta debía traer su métrica asociada")
	require.Len(t, encontrado.Lesiones, 1, "Has-Many: el atleta debía traer su lesión asociada")

	metrica, ok := almacen.BuscarMetricaPorID(encontrado.Metricas[0].ID)
	require.True(t, ok)
	require.NotNil(t, metrica.Atleta, "Belongs-To: la métrica debía traer su atleta asociado")
	assert.Equal(t, "María Torres", metrica.Atleta.Nombre)
}

func TestAlmacenSQLite_CrearYBuscarMetrica(t *testing.T) {
	db := nuevoDBAtletaPrueba(t)
	almacen := NuevoAlmacenSQLite(db)
	atleta := almacen.CrearAtleta(models.Atleta{Nombre: "José Ruiz", Peso: 80, CoachID: 1})

	creada := almacen.CrearMetrica(models.MetricaFisica{AtletaID: atleta.ID, PorcentajeGrasa: 18, MasaMuscular: 35})
	assert.NotZero(t, creada.ID)

	encontrada, ok := almacen.BuscarMetricaPorID(creada.ID)
	require.True(t, ok)
	assert.Equal(t, atleta.ID, encontrada.AtletaID)

	_, ok = almacen.BuscarMetricaPorID(9999)
	assert.False(t, ok)
}

func TestAlmacenSQLite_CrearYBuscarLesion(t *testing.T) {
	db := nuevoDBAtletaPrueba(t)
	almacen := NuevoAlmacenSQLite(db)
	atleta := almacen.CrearAtleta(models.Atleta{Nombre: "Ana Cedeño", Peso: 60, CoachID: 1})

	creada := almacen.CrearLesion(models.Lesion{AtletaID: atleta.ID, Descripcion: "Tendinitis", Gravedad: "media", Estado: "activa"})
	assert.NotZero(t, creada.ID)

	encontrada, ok := almacen.BuscarLesionPorID(creada.ID)
	require.True(t, ok)
	assert.Equal(t, "Tendinitis", encontrada.Descripcion)

	_, ok = almacen.BuscarLesionPorID(9999)
	assert.False(t, ok)
}