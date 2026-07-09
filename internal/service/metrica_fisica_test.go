package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"atletismo-api/internal/models"
)

func TestAtletaService_CrearMetrica(t *testing.T) {
	casos := []struct {
		nombre    string
		entrada   models.MetricaFisica
		esperaErr error
	}{
		{
			nombre:    "atleta_id vacío → rechazado sin tocar el repositorio",
			entrada:   models.MetricaFisica{AtletaID: 0, PorcentajeGrasa: 12, MasaMuscular: 40},
			esperaErr: ErrAtletaIDInvalido,
		},
		{
			nombre:    "porcentaje_grasa inválido → rechazado sin tocar el repositorio",
			entrada:   models.MetricaFisica{AtletaID: 1, PorcentajeGrasa: -5, MasaMuscular: 40},
			esperaErr: ErrPorcentajeGrasaInvalido,
		},
		{
			nombre:    "masa_muscular inválida → rechazado sin tocar el repositorio",
			entrada:   models.MetricaFisica{AtletaID: 1, PorcentajeGrasa: 12, MasaMuscular: 0},
			esperaErr: ErrMasaMuscularInvalida,
		},
		{
			nombre:    "datos válidos → persiste en el repositorio",
			entrada:   models.MetricaFisica{AtletaID: 1, PorcentajeGrasa: 12.5, MasaMuscular: 42.3},
			esperaErr: nil,
		},
	}

	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			mock := nuevoAtletaMock()
			svc := NewAtletaService(mock)

			res, err := svc.CrearMetrica(c.entrada)

			if c.esperaErr != nil {
				require.ErrorIs(t, err, c.esperaErr)
				assert.False(t, mock.llamoCrearMetrica, "el repositorio NO debía ser invocado")
			} else {
				require.NoError(t, err)
				assert.True(t, mock.llamoCrearMetrica, "el repositorio debía ser invocado")
				assert.NotZero(t, res.ID)
			}
		})
	}
}

func TestAtletaService_ObtenerMetrica(t *testing.T) {
	mock := nuevoAtletaMock()
	svc := NewAtletaService(mock)
	creada := mock.CrearMetrica(models.MetricaFisica{AtletaID: 1, PorcentajeGrasa: 15, MasaMuscular: 38})

	t.Run("ID existente → devuelve métrica", func(t *testing.T) {
		m, err := svc.ObtenerMetrica(creada.ID)
		require.NoError(t, err)
		assert.Equal(t, 1, m.AtletaID)
	})

	t.Run("ID inexistente → ErrMetricaNoEncontrada", func(t *testing.T) {
		_, err := svc.ObtenerMetrica(9999)
		require.ErrorIs(t, err, ErrMetricaNoEncontrada)
	})
}

func TestAtletaService_BorrarMetrica(t *testing.T) {
	mock := nuevoAtletaMock()
	svc := NewAtletaService(mock)
	creada := mock.CrearMetrica(models.MetricaFisica{AtletaID: 1, PorcentajeGrasa: 20, MasaMuscular: 45})

	t.Run("ID inexistente → error", func(t *testing.T) {
		err := svc.BorrarMetrica(9999)
		require.ErrorIs(t, err, ErrMetricaNoEncontrada)
	})

	t.Run("ID existente → borra sin error", func(t *testing.T) {
		err := svc.BorrarMetrica(creada.ID)
		require.NoError(t, err)
	})
}