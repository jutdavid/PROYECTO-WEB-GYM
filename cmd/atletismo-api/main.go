package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"

	"atletismo-api/internal/config"
	"atletismo-api/internal/handlers"
	"atletismo-api/internal/httpserver"
	"atletismo-api/internal/middleware"
	"atletismo-api/internal/service"
	"atletismo-api/internal/storage"
)

func main() {
	cfg := config.Cargar()
	if err := run(cfg); err != nil {
		log.Fatal(err)
	}
}

func run(cfg config.Config) error {
	// 1. Factory: abre DB, migra, siembra y elige backend.
	recursos, err := storage.Inicializar(cfg.RutaDB, cfg.Backend)
	if err != nil {
		return err
	}
	defer func() { _ = recursos.Cerrar() }()
	log.Printf("Backend de almacenamiento: %s", recursos.BackendUsado)

	// 2. Servicios.
	authService := service.NewAuthService(recursos.Usuarios)
	ciclosService := service.NewCicloEntrenamientoService(recursos.Almacen)
	evaluacionesService := service.NewEvaluacionCicloService(recursos.Almacen)
	microciclosService := service.NewMicrocicloService(recursos.Almacen)
	servidor := handlers.NewServer(ciclosService, evaluacionesService, microciclosService, authService)

	// 3. Router + middleware.
	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/auth/registrar", servidor.Registrar)
		r.Post("/auth/login", servidor.Login)

		r.Group(func(r chi.Router) {
			// Ciclos de entrenamiento CRUD
			r.Get("/ciclos", servidor.ListarCiclos)
			r.Post("/ciclos", servidor.CrearCiclo)
			r.Get("/ciclos/{id}", servidor.ObtenerCiclo)
			r.Put("/ciclos/{id}", servidor.ActualizarCiclo)
			r.Delete("/ciclos/{id}", servidor.BorrarCiclo)

			// Microciclos CRUD
			r.Get("/microciclos", servidor.ListarMicrociclos)
			r.Post("/microciclos", servidor.CrearMicrociclo)
			r.Get("/microciclos/{id}", servidor.ObtenerMicrociclo)
			r.Delete("/microciclos/{id}", servidor.BorrarMicrociclo)

			// Evaluaciones de Ciclo CRUD
			r.Get("/evaluaciones", servidor.ListarEvaluaciones)
			r.Post("/evaluaciones", servidor.CrearEvaluacion)
			r.Get("/evaluaciones/{id}", servidor.ObtenerEvaluacion)
			r.Delete("/evaluaciones/{id}", servidor.BorrarEvaluacion)
		})
	})

	// 4. Servidor HTTP con timeouts.
	srv := httpserver.Nuevo(
		r,
		httpserver.ConPuerto(cfg.Puerto),
		httpserver.ConReadTimeout(cfg.ReadTimeout),
		httpserver.ConWriteTimeout(cfg.WriteTimeout),
	)

	// 5. Graceful shutdown.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	errServidor := make(chan error, 1)
	go func() {
		log.Printf("Servidor escuchando en http://localhost%s", cfg.Puerto)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errServidor <- err
		}
	}()

	select {
	case err := <-errServidor:
		return err
	case <-ctx.Done():
		log.Println("Senal de apagado recibida, cerrando ordenadamente...")
	}

	ctxApagado, cancelar := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelar()
	if err := srv.Shutdown(ctxApagado); err != nil {
		return err
	}
	log.Println("Servidor detenido limpiamente.")
	return nil
}
