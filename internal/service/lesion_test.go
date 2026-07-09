package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"atletismo-api/internal/models"
)

func TestAtletaService_CrearLesion(t *testing.T) {
	casos := []struct {
		nombre    string
		entrada   models.Lesion
		esperaErr error
	}{
		{
			nombre:    "atleta_id vacío → rechazado sin tocar el repositorio",
			entrada:   models.Lesion{AtletaID: 0, Descripcion: "Esguince", Gravedad: "leve", Estado: "activa"},
			esperaErr: ErrAtletaIDInvalido,
		},
		{
			nombre:    "descripcion vacía → rechazado sin tocar el repositorio",
			entrada:   models.Lesion{AtletaID: 1, Descripcion: "", Gravedad: "leve", Estado: "activa"},
			esperaErr: ErrDescripcionVacia,
		},
		{
			nombre:    "gravedad vacía → rechazado sin tocar el repositorio",
			entrada:   models.Lesion{AtletaID: 1, Descripcion: "Esguince", Gravedad: "", Estado: "activa"},
			esperaErr: ErrGravedadVacia,
		},
		{
			nombre:    "estado vacío → rechazado sin tocar el repositorio",
			entrada:   models.Lesion{AtletaID: 1, Descripcion: "Esguince", Gravedad: "leve", Estado: ""},
			esperaErr: ErrEstadoLesionVacio,
		},
		{
			nombre:    "datos válidos → persiste en el repositorio",
			entrada:   models.Lesion{AtletaID: 1, Descripcion: "Esguince de tobillo", Gravedad: "leve", Estado: "activa"},
			esperaErr: nil,
		},
	}

	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			mock := nuevoAtletaMock()
			svc := NewAtletaService(mock)

			res, err := svc.CrearLesion(c.entrada)

			if c.esperaErr != nil {
				require.ErrorIs(t, err, c.esperaErr)
				assert.False(t, mock.llamoCrearLesion, "el repositorio NO debía ser invocado")
			} else {
				require.NoError(t, err)
				assert.True(t, mock.llamoCrearLesion, "el repositorio debía ser invocado")
				assert.NotZero(t, res.ID)
			}
		})
	}
}

func TestAtletaService_ObtenerLesion(t *testing.T) {
	mock := nuevoAtletaMock()
	svc := NewAtletaService(mock)
	creada := mock.CrearLesion(models.Lesion{AtletaID: 1, Descripcion: "Tendinitis", Gravedad: "media", Estado: "activa"})

	t.Run("ID existente → devuelve lesión", func(t *testing.T) {
		l, err := svc.ObtenerLesion(creada.ID)
		require.NoError(t, err)
		assert.Equal(t, "Tendinitis", l.Descripcion)
	})

	t.Run("ID inexistente → ErrLesionNoEncontrada", func(t *testing.T) {
		_, err := svc.ObtenerLesion(9999)
		require.ErrorIs(t, err, ErrLesionNoEncontrada)
	})
}

func TestAtletaService_BorrarLesion(t *testing.T) {
	mock := nuevoAtletaMock()
	svc := NewAtletaService(mock)
	creada := mock.CrearLesion(models.Lesion{AtletaID: 1, Descripcion: "Fractura", Gravedad: "grave", Estado: "activa"})

	t.Run("ID inexistente → error", func(t *testing.T) {
		err := svc.BorrarLesion(9999)
		require.ErrorIs(t, err, ErrLesionNoEncontrada)
	})

	t.Run("ID existente → borra sin error", func(t *testing.T) {
		err := svc.BorrarLesion(creada.ID)
		require.NoError(t, err)
	})
}