package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"atletismo-api/internal/middleware"
	"atletismo-api/internal/models"
	"atletismo-api/internal/service"
)

// ─── Fake de AlmacenAtleta en memoria ─────────────────────────────────────────

type atletaFake struct {
	datos  []models.Atleta
	nextID int
}

func nuevoAtletaFake() *atletaFake { return &atletaFake{nextID: 1} }

func (f *atletaFake) ListarAtletas() []models.Atleta { return f.datos }
func (f *atletaFake) BuscarAtletaPorID(id int) (models.Atleta, bool) {
	for _, a := range f.datos {
		if a.ID == id {
			return a, true
		}
	}
	return models.Atleta{}, false
}
func (f *atletaFake) CrearAtleta(a models.Atleta) models.Atleta {
	a.ID = f.nextID
	f.nextID++
	f.datos = append(f.datos, a)
	return a
}
func (f *atletaFake) ActualizarAtleta(id int, d models.Atleta) (models.Atleta, bool) {
	for i, a := range f.datos {
		if a.ID == id {
			d.ID = id
			f.datos[i] = d
			return d, true
		}
	}
	return models.Atleta{}, false
}
func (f *atletaFake) BorrarAtleta(id int) bool {
	for i, a := range f.datos {
		if a.ID == id {
			f.datos = append(f.datos[:i], f.datos[i+1:]...)
			return true
		}
	}
	return false
}

// MetricaFisica y Lesion: no se ejercitan aquí, solo cumplen la interfaz.
func (f *atletaFake) ListarMetricas() []models.MetricaFisica { return nil }
func (f *atletaFake) BuscarMetricaPorID(id int) (models.MetricaFisica, bool) {
	return models.MetricaFisica{}, false
}
func (f *atletaFake) CrearMetrica(m models.MetricaFisica) models.MetricaFisica { return m }
func (f *atletaFake) ActualizarMetrica(id int, d models.MetricaFisica) (models.MetricaFisica, bool) {
	return d, true
}
func (f *atletaFake) BorrarMetrica(id int) bool       { return true }
func (f *atletaFake) ListarLesiones() []models.Lesion { return nil }
func (f *atletaFake) BuscarLesionPorID(id int) (models.Lesion, bool) {
	return models.Lesion{}, false
}
func (f *atletaFake) CrearLesion(l models.Lesion) models.Lesion { return l }
func (f *atletaFake) ActualizarLesion(id int, d models.Lesion) (models.Lesion, bool) {
	return d, true
}
func (f *atletaFake) BorrarLesion(id int) bool { return true }

// ─── Fake de UserRepository ───────────────────────────────────────────────────

type usuarioAtletaFake struct {
	datos  []models.Usuario
	nextID int
}

func nuevoUsuarioAtletaFake() *usuarioAtletaFake { return &usuarioAtletaFake{nextID: 1} }

func (u *usuarioAtletaFake) CrearUsuario(usr models.Usuario) (models.Usuario, error) {
	usr.ID = u.nextID
	u.nextID++
	u.datos = append(u.datos, usr)
	return usr, nil
}
func (u *usuarioAtletaFake) BuscarUsuarioPorEmail(email string) (models.Usuario, bool) {
	for _, usr := range u.datos {
		if usr.Email == email {
			return usr, true
		}
	}
	return models.Usuario{}, false
}

// ─── Helper: genera un token JWT válido para los tests ───────────────────────

func generarTokenAtletaTest(t *testing.T, auth *service.AuthService) string {
	t.Helper()
	u := models.Usuario{ID: 1, Email: "test@atletismo.com", Password: "x"}
	token, err := auth.GenerarToken(u)
	require.NoError(t, err, "no se pudo generar token de prueba")
	return "Bearer " + token
}

// ─── Helper: router mínimo solo con el módulo de Atletas ─────────────────────

func nuevoRouterAtletaTest(t *testing.T) (http.Handler, *service.AuthService) {
	t.Helper()

	atletaService := service.NewAtletaService(nuevoAtletaFake())
	authService := service.NewAuthService(nuevoUsuarioAtletaFake())

	servidor := &Server{
		Atleta: atletaService,
		Auth:   authService,
	}

	r := chi.NewRouter()
	r.Route("/api/v1", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authService))
			r.Get("/atletas", servidor.ListarAtletas)
			r.Post("/atletas", servidor.CrearAtleta)
			r.Get("/atletas/{id}", servidor.ObtenerAtleta)
		})
	})

	return r, authService
}

// ─── Tests ────────────────────────────────────────────────────────────────────

// 401 — GET /atletas SIN token debe devolver 401 (obligatorio).
func TestHandler_GetAtletas_SinToken_Devuelve401(t *testing.T) {
	router, _ := nuevoRouterAtletaTest(t)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/atletas", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

// 201 — POST /atletas con datos válidos y token → 201 Created.
func TestHandler_CrearAtleta_Valido_Devuelve201(t *testing.T) {
	router, auth := nuevoRouterAtletaTest(t)
	token := generarTokenAtletaTest(t, auth)

	cuerpo, _ := json.Marshal(models.Atleta{
		Nombre:              "Pedro Gómez",
		MetodologiaObjetivo: "Sprint 100m",
		Peso:                72.5,
		CoachID:             1,
	})

	req := httptest.NewRequest(http.MethodPost, "/api/v1/atletas", bytes.NewReader(cuerpo))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code, "body: %s", rec.Body.String())

	var respuesta models.Atleta
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &respuesta))
	assert.NotZero(t, respuesta.ID)
	assert.Equal(t, "Pedro Gómez", respuesta.Nombre)
}

// 400 — POST /atletas con peso negativo → 400 Bad Request.
func TestHandler_CrearAtleta_PesoNegativo_Devuelve400(t *testing.T) {
	router, auth := nuevoRouterAtletaTest(t)
	token := generarTokenAtletaTest(t, auth)

	cuerpo, _ := json.Marshal(models.Atleta{
		Nombre:  "Pedro Gómez",
		Peso:    -10,
		CoachID: 1,
	})

	req := httptest.NewRequest(http.MethodPost, "/api/v1/atletas", bytes.NewReader(cuerpo))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code, "body: %s", rec.Body.String())
}