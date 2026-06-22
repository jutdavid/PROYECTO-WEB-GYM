package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"atletismo-api/internal/handlers"
	"atletismo-api/internal/storage"
)

func main() {
	// 1. Crear el almacenamiento y cargar datos iniciales.
	almacen := storage.NuevaMemoria()
	almacen.SeedEntrenadores()
	almacen.SeedAtletas()
	almacen.SeedCiclos()
	almacen.SeedMetricas()
	almacen.SeedLesiones()
	almacen.SeedMicrociclos()
	almacen.SeedEvaluaciones()
	almacen.SeedHorarios()
	almacen.SeedCertificaciones()

	// 2. Crear el Server inyectándole el almacenamiento.
	servidor := handlers.NewServer(almacen)

	// 3. Configurar el router con versionado /api/v1/.
	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {
		// Entrenadores CRUD
		r.Get("/entrenadores", servidor.ListarEntrenadores)
		r.Post("/entrenadores", servidor.CrearEntrenador)
		r.Get("/entrenadores/{id}", servidor.ObtenerEntrenador)
		r.Put("/entrenadores/{id}", servidor.ActualizarEntrenador)
		r.Delete("/entrenadores/{id}", servidor.BorrarEntrenador)

		// Certificaciones CRUD
		r.Get("/certificaciones", servidor.ListarCertificaciones)
		r.Post("/certificaciones", servidor.CrearCertificacion)
		r.Get("/certificaciones/{id}", servidor.ObtenerCertificacion)
		r.Delete("/certificaciones/{id}", servidor.BorrarCertificacion)

		// Horarios de Atención CRUD
		r.Get("/horarios", servidor.ListarHorarios)
		r.Post("/horarios", servidor.CrearHorario)
		r.Get("/horarios/{id}", servidor.ObtenerHorario)
		r.Delete("/horarios/{id}", servidor.BorrarHorario)

		// Atletas CRUD
		r.Get("/atletas", servidor.ListarAtletas)
		r.Post("/atletas", servidor.CrearAtleta)
		r.Get("/atletas/{id}", servidor.ObtenerAtleta)
		r.Put("/atletas/{id}", servidor.ActualizarAtleta)
		r.Delete("/atletas/{id}", servidor.BorrarAtleta)

		// Métricas Físicas CRUD
		r.Get("/metricas", servidor.ListarMetricas)
		r.Post("/metricas", servidor.CrearMetrica)
		r.Get("/metricas/{id}", servidor.ObtenerMetrica)
		r.Delete("/metricas/{id}", servidor.BorrarMetrica)

		// Lesiones CRUD
		r.Get("/lesiones", servidor.ListarLesiones)
		r.Post("/lesiones", servidor.CrearLesion)
		r.Get("/lesiones/{id}", servidor.ObtenerLesion)
		r.Delete("/lesiones/{id}", servidor.BorrarLesion)

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

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
