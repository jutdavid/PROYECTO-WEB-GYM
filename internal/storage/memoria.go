package storage

import (
	"sync"

	"atletismo-api/internal/models"
)

// Memoria es un almacén unificado del sistema de atletismo.
type Memoria struct {
	atletas      []models.Atleta
	nextAtletaID uint

	entrenadores     []models.Entrenador
	nextEntrenadorID uint

	ciclos      []models.CicloEntrenamiento
	nextCicloID uint

	mu sync.Mutex
}

// NuevaMemoria crea un almacén vacío y listo para usar.
func NuevaMemoria() *Memoria {
	return &Memoria{
		atletas:          []models.Atleta{},
		nextAtletaID:     1,
		entrenadores:     []models.Entrenador{},
		nextEntrenadorID: 1,
		ciclos:           []models.CicloEntrenamiento{},
		nextCicloID:      1,
	}
}

func (m *Memoria) SeedEntrenadores() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.entrenadores = []models.Entrenador{
		{ID: 1, Nombre: "Carlos Mendoza", Especialidad: "Velocidad", CapacidadMaxima: 10, CargaActual: 2.5},
		{ID: 2, Nombre: "Ana Ríos", Especialidad: "Resistencia", CapacidadMaxima: 8, CargaActual: 3.0},
		{ID: 3, Nombre: "Luis Paredes", Especialidad: "Fuerza", CapacidadMaxima: 6, CargaActual: 1.5},
	}
	m.nextEntrenadorID = 4
}

func (m *Memoria) ListarEntrenadores() []models.Entrenador {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.Entrenador, len(m.entrenadores))
	copy(copia, m.entrenadores)
	return copia
}

func (m *Memoria) BuscarEntrenadorPorID(id uint) (models.Entrenador, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, e := range m.entrenadores {
		if e.ID == id {
			return e, true
		}
	}
	return models.Entrenador{}, false
}

func (m *Memoria) CrearEntrenador(e models.Entrenador) models.Entrenador {
	m.mu.Lock()
	defer m.mu.Unlock()

	e.ID = m.nextEntrenadorID
	m.nextEntrenadorID++
	m.entrenadores = append(m.entrenadores, e)
	return e
}

func (m *Memoria) ActualizarEntrenador(id uint, datos models.Entrenador) (models.Entrenador, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, e := range m.entrenadores {
		if e.ID == id {
			datos.ID = id
			m.entrenadores[i] = datos
			return datos, true
		}
	}
	return models.Entrenador{}, false
}

func (m *Memoria) BorrarEntrenador(id uint) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, e := range m.entrenadores {
		if e.ID == id {
			m.entrenadores = append(m.entrenadores[:i], m.entrenadores[i+1:]...)
			return true
		}
	}
	return false
}

func (m *Memoria) SeedAtletas() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.atletas = []models.Atleta{
		{ID: 1, Nombre: "Pedro Gómez", MetodologiaObjetivo: "Sprint 100m", Peso: 72.5, CoachID: 1},
		{ID: 2, Nombre: "María Torres", MetodologiaObjetivo: "Maratón", Peso: 58.0, CoachID: 2},
		{ID: 3, Nombre: "José Ruiz", MetodologiaObjetivo: "Salto de longitud", Peso: 80.0, CoachID: 1},
		{ID: 4, Nombre: "Laura Vega", MetodologiaObjetivo: "Lanzamiento de bala", Peso: 90.0, CoachID: 3},
		{ID: 5, Nombre: "Andrés León", MetodologiaObjetivo: "5000m", Peso: 65.0, CoachID: 2},
	}
	m.nextAtletaID = 6
}

func (m *Memoria) ListarAtletas() []models.Atleta {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.Atleta, len(m.atletas))
	copy(copia, m.atletas)
	return copia
}

func (m *Memoria) BuscarAtletaPorID(id uint) (models.Atleta, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, a := range m.atletas {
		if a.ID == id {
			return a, true
		}
	}
	return models.Atleta{}, false
}

func (m *Memoria) CrearAtleta(a models.Atleta) models.Atleta {
	m.mu.Lock()
	defer m.mu.Unlock()

	a.ID = m.nextAtletaID
	m.nextAtletaID++
	m.atletas = append(m.atletas, a)
	return a
}

func (m *Memoria) ActualizarAtleta(id uint, datos models.Atleta) (models.Atleta, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, a := range m.atletas {
		if a.ID == id {
			datos.ID = id
			m.atletas[i] = datos
			return datos, true
		}
	}
	return models.Atleta{}, false
}

func (m *Memoria) BorrarAtleta(id uint) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, a := range m.atletas {
		if a.ID == id {
			m.atletas = append(m.atletas[:i], m.atletas[i+1:]...)
			return true
		}
	}
	return false
}

func (m *Memoria) SeedCiclos() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.ciclos = []models.CicloEntrenamiento{
		{ID: 1, AtletaID: 1, Estado: "activo"},
		{ID: 2, AtletaID: 2, Estado: "completado"},
		{ID: 3, AtletaID: 3, Estado: "activo"},
	}
	m.nextCicloID = 4
}

func (m *Memoria) ListarCiclos() []models.CicloEntrenamiento {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.CicloEntrenamiento, len(m.ciclos))
	copy(copia, m.ciclos)
	return copia
}

func (m *Memoria) BuscarCicloPorID(id uint) (models.CicloEntrenamiento, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, c := range m.ciclos {
		if c.ID == id {
			return c, true
		}
	}
	return models.CicloEntrenamiento{}, false
}

func (m *Memoria) CrearCiclo(c models.CicloEntrenamiento) models.CicloEntrenamiento {
	m.mu.Lock()
	defer m.mu.Unlock()

	c.ID = m.nextCicloID
	m.nextCicloID++
	m.ciclos = append(m.ciclos, c)
	return c
}

func (m *Memoria) ActualizarCiclo(id uint, datos models.CicloEntrenamiento) (models.CicloEntrenamiento, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, c := range m.ciclos {
		if c.ID == id {
			datos.ID = id
			m.ciclos[i] = datos
			return datos, true
		}
	}
	return models.CicloEntrenamiento{}, false
}

func (m *Memoria) BorrarCiclo(id uint) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, c := range m.ciclos {
		if c.ID == id {
			m.ciclos = append(m.ciclos[:i], m.ciclos[i+1:]...)
			return true
		}
	}
	return false
}
