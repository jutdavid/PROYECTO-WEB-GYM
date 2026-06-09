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

		// Atletas CRUD
		r.Get("/atletas", servidor.ListarAtletas)
		r.Post("/atletas", servidor.CrearAtleta)
		r.Get("/atletas/{id}", servidor.ObtenerAtleta)
		r.Put("/atletas/{id}", servidor.ActualizarAtleta)
		r.Delete("/atletas/{id}", servidor.BorrarAtleta)

		// Ciclos de entrenamiento CRUD
		r.Get("/ciclos", servidor.ListarCiclos)
		r.Post("/ciclos", servidor.CrearCiclo)
		r.Get("/ciclos/{id}", servidor.ObtenerCiclo)
		r.Put("/ciclos/{id}", servidor.ActualizarCiclo)
		r.Delete("/ciclos/{id}", servidor.BorrarCiclo)
	})

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
