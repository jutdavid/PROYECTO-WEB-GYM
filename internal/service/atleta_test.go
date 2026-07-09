package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"atletismo-api/internal/models"
)

// ─── Mock de AlmacenAtleta ────────────────────────────────────────────────────

type atletaAlmacenMock struct {
	atletas           []models.Atleta
	metricas          []models.MetricaFisica
	lesiones          []models.Lesion
	nextID            int
	llamoCrear        bool
	llamoCrearMetrica bool
	llamoCrearLesion  bool
}

func nuevoAtletaMock() *atletaAlmacenMock { return &atletaAlmacenMock{nextID: 1} }

// Atletas
func (m *atletaAlmacenMock) ListarAtletas() []models.Atleta { return m.atletas }
func (m *atletaAlmacenMock) BuscarAtletaPorID(id int) (models.Atleta, bool) {
	for _, a := range m.atletas {
		if a.ID == id {
			return a, true
		}
	}
	return models.Atleta{}, false
}
func (m *atletaAlmacenMock) CrearAtleta(a models.Atleta) models.Atleta {
	m.llamoCrear = true
	a.ID = m.nextID
	m.nextID++
	m.atletas = append(m.atletas, a)
	return a
}
func (m *atletaAlmacenMock) ActualizarAtleta(id int, d models.Atleta) (models.Atleta, bool) {
	for i, a := range m.atletas {
		if a.ID == id {
			d.ID = id
			m.atletas[i] = d
			return d, true
		}
	}
	return models.Atleta{}, false
}
func (m *atletaAlmacenMock) BorrarAtleta(id int) bool {
	for i, a := range m.atletas {
		if a.ID == id {
			m.atletas = append(m.atletas[:i], m.atletas[i+1:]...)
			return true
		}
	}
	return false
}

// Metricas
func (m *atletaAlmacenMock) ListarMetricas() []models.MetricaFisica { return m.metricas }
func (m *atletaAlmacenMock) BuscarMetricaPorID(id int) (models.MetricaFisica, bool) {
	for _, met := range m.metricas {
		if met.ID == id {
			return met, true
		}
	}
	return models.MetricaFisica{}, false
}
func (m *atletaAlmacenMock) CrearMetrica(mf models.MetricaFisica) models.MetricaFisica {
	m.llamoCrearMetrica = true
	mf.ID = m.nextID
	m.nextID++
	m.metricas = append(m.metricas, mf)
	return mf
}
func (m *atletaAlmacenMock) ActualizarMetrica(id int, d models.MetricaFisica) (models.MetricaFisica, bool) {
	for i, met := range m.metricas {
		if met.ID == id {
			d.ID = id
			m.metricas[i] = d
			return d, true
		}
	}
	return models.MetricaFisica{}, false
}
func (m *atletaAlmacenMock) BorrarMetrica(id int) bool {
	for i, met := range m.metricas {
		if met.ID == id {
			m.metricas = append(m.metricas[:i], m.metricas[i+1:]...)
			return true
		}
	}
	return false
}

// Lesiones
func (m *atletaAlmacenMock) ListarLesiones() []models.Lesion { return m.lesiones }
func (m *atletaAlmacenMock) BuscarLesionPorID(id int) (models.Lesion, bool) {
	for _, l := range m.lesiones {
		if l.ID == id {
			return l, true
		}
	}
	return models.Lesion{}, false
}
func (m *atletaAlmacenMock) CrearLesion(l models.Lesion) models.Lesion {
	m.llamoCrearLesion = true
	l.ID = m.nextID
	m.nextID++
	m.lesiones = append(m.lesiones, l)
	return l
}
func (m *atletaAlmacenMock) ActualizarLesion(id int, d models.Lesion) (models.Lesion, bool) {
	for i, l := range m.lesiones {
		if l.ID == id {
			d.ID = id
			m.lesiones[i] = d
			return d, true
		}
	}
	return models.Lesion{}, false
}
func (m *atletaAlmacenMock) BorrarLesion(id int) bool {
	for i, l := range m.lesiones {
		if l.ID == id {
			m.lesiones = append(m.lesiones[:i], m.lesiones[i+1:]...)
			return true
		}
	}
	return false
}

// ═══════════════════════════ TEST: CrearAtleta ═══════════════════════════════

func TestAtletaService_CrearAtleta(t *testing.T) {
	casos := []struct {
		nombre    string
		entrada   models.Atleta
		esperaErr error
	}{
		{
			nombre:    "nombre vacío → rechazado sin tocar el repositorio",
			entrada:   models.Atleta{Nombre: "", Peso: 70, CoachID: 1},
			esperaErr: ErrNombreVacio,
		},
		{
			nombre:    "peso negativo → rechazado sin tocar el repositorio (regla de negocio clave)",
			entrada:   models.Atleta{Nombre: "Pedro Gómez", Peso: -5, CoachID: 1},
			esperaErr: ErrPesoInvalido,
		},
		{
			nombre:    "coach_id vacío → rechazado sin tocar el repositorio",
			entrada:   models.Atleta{Nombre: "Pedro Gómez", Peso: 70, CoachID: 0},
			esperaErr: ErrCoachIDInvalido,
		},
		{
			nombre:    "datos válidos → persiste en el repositorio",
			entrada:   models.Atleta{Nombre: "Pedro Gómez", Peso: 72.5, CoachID: 1},
			esperaErr: nil,
		},
	}

	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			mock := nuevoAtletaMock()
			svc := NewAtletaService(mock)

			res, err := svc.CrearAtleta(c.entrada)

			if c.esperaErr != nil {
				require.ErrorIs(t, err, c.esperaErr)
				assert.False(t, mock.llamoCrear, "el repositorio NO debía ser invocado ante datos inválidos")
			} else {
				require.NoError(t, err)
				assert.True(t, mock.llamoCrear, "el repositorio debía ser invocado")
				assert.NotZero(t, res.ID)
				assert.Equal(t, c.entrada.Nombre, res.Nombre)
			}
		})
	}
}

func TestAtletaService_ObtenerAtleta(t *testing.T) {
	mock := nuevoAtletaMock()
	svc := NewAtletaService(mock)
	creado := mock.CrearAtleta(models.Atleta{Nombre: "María Torres", Peso: 58, CoachID: 2})

	t.Run("ID existente → devuelve atleta", func(t *testing.T) {
		a, err := svc.ObtenerAtleta(creado.ID)
		require.NoError(t, err)
		assert.Equal(t, "María Torres", a.Nombre)
	})

	t.Run("ID inexistente → ErrAtletaNoEncontrado", func(t *testing.T) {
		_, err := svc.ObtenerAtleta(9999)
		require.ErrorIs(t, err, ErrAtletaNoEncontrado)
	})
}

func TestAtletaService_BorrarAtleta(t *testing.T) {
	mock := nuevoAtletaMock()
	svc := NewAtletaService(mock)
	creado := mock.CrearAtleta(models.Atleta{Nombre: "José Ruiz", Peso: 80, CoachID: 1})

	t.Run("ID inexistente → error", func(t *testing.T) {
		err := svc.BorrarAtleta(9999)
		require.ErrorIs(t, err, ErrAtletaNoEncontrado)
	})

	t.Run("ID existente → borra sin error", func(t *testing.T) {
		err := svc.BorrarAtleta(creado.ID)
		require.NoError(t, err)
	})
}