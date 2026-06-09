# Atletismo API

API REST para la gestión de un sistema de atletismo universitario. Permite administrar **Entrenadores**, **Atletas** y **Ciclos de Entrenamiento** con operaciones CRUD completas.

## Tecnologías

- **Go 1.22**
- **Chi v5** — router HTTP ligero
- Almacenamiento en memoria (sin base de datos)

## Estructura del proyecto

```
atletismo-api/
├── go.mod
├── go.sum
├── test_endpoints.sh            ← script de pruebas con curl
├── cmd/
│   └── atletismo-api/
│       └── main.go              ← entry point
└── internal/
    ├── models/
    │   ├── atleta.go
    │   ├── entrenador.go
    │   └── ciclo_entrenamiento.go
    ├── storage/
    │   └── memoria.go           ← almacén en memoria con mutex
    └── handlers/
        ├── atleta.go
        ├── entrenador.go
        ├── ciclo.go
        └── respond.go           ← helpers JSON
```

## Requisitos

- Go 1.22 
- `curl` o Postman para probar los endpoints

## Cómo correrlo

```bash
go mod tidy
go run ./cmd/atletismo-api
```

Verás en consola:

```
Servidor escuchando en http://localhost:8080
```
### Modelos

### Entrenador

```json
{
  "id": 1,
  "nombre": "Carlos Mendoza",
  "especialidad": "Velocidad",
  "capacidad_maxima": 10,
  "carga_actual": 2.5
}
```

### Atleta

```json
{
  "id": 1,
  "nombre": "Pedro Gómez",
  "metodologia_objetivo": "Sprint 100m",
  "peso": 72.5,
  "coach_id": 1
}
```

### CicloEntrenamiento

```json
{
  "id": 1,
  "atleta_id": 1,
  "estado": "activo",
  "fecha_inicio": "2025-01-15T08:00:00Z"
}
```

## Datos pre-cargados al arrancar

### Entrenadores

| ID | Nombre | Especialidad | Capacidad Máx. | Carga Actual |
|----|--------|--------------|----------------|--------------|
| 1 | Carlos Mendoza | Velocidad | 10 | 2.5 |
| 2 | Ana Ríos | Resistencia | 8 | 3.0 |
| 3 | Luis Paredes | Fuerza | 6 | 1.5 |

### Atletas

| ID | Nombre | Metodología Objetivo | Peso | CoachID |
|----|--------|----------------------|------|---------|
| 1 | Pedro Gómez | Sprint 100m | 72.5 | 1 |
| 2 | María Torres | Maratón | 58.0 | 2 |
| 3 | José Ruiz | Salto de longitud | 80.0 | 1 |
| 4 | Laura Vega | Lanzamiento de bala | 90.0 | 3 |
| 5 | Andrés León | 5000m | 65.0 | 2 |

### Ciclos de Entrenamiento

| ID | AtletaID | Estado |
|----|----------|--------|
| 1 | 1 | activo |
| 2 | 2 | completado |
| 3 | 3 | activo |

## Endpoints

### Entrenadores

| Método | Ruta | Descripción | Status |
|--------|------|-------------|--------|
| GET | `/api/v1/entrenadores` | Lista todos | 200 |
| GET | `/api/v1/entrenadores/{id}` | Obtiene uno | 200 / 404 |
| POST | `/api/v1/entrenadores` | Crea uno | 201 / 400 |
| PUT | `/api/v1/entrenadores/{id}` | Actualiza uno | 200 / 400 / 404 |
| DELETE | `/api/v1/entrenadores/{id}` | Elimina uno | 204 / 404 |

### Atletas

| Método | Ruta | Descripción | Status |
|--------|------|-------------|--------|
| GET | `/api/v1/atletas` | Lista todos | 200 |
| GET | `/api/v1/atletas/{id}` | Obtiene uno | 200 / 404 |
| POST | `/api/v1/atletas` | Crea uno | 201 / 400 |
| PUT | `/api/v1/atletas/{id}` | Actualiza uno | 200 / 400 / 404 |
| DELETE | `/api/v1/atletas/{id}` | Elimina uno | 204 / 404 |

### Ciclos de Entrenamiento

| Método | Ruta | Descripción | Status |
|--------|------|-------------|--------|
| GET | `/api/v1/ciclos` | Lista todos | 200 |
| GET | `/api/v1/ciclos/{id}` | Obtiene uno | 200 / 404 |
| POST | `/api/v1/ciclos` | Crea uno | 201 / 400 |
| PUT | `/api/v1/ciclos/{id}` | Actualiza uno | 200 / 400 / 404 |
| DELETE | `/api/v1/ciclos/{id}` | Elimina uno | 204 / 404 |

## Validaciones

### Entrenador
- `nombre` no puede estar vacío → **400**
- `capacidad_maxima` debe ser mayor a 0 → **400**
- `carga_actual` no puede ser negativa → **400**

### Atleta
- `nombre` no puede estar vacío → **400**
- `peso` no puede ser negativo → **400**

### CicloEntrenamiento
- `atleta_id` es obligatorio (debe ser > 0) → **400**
- `estado` no puede estar vacío → **400**
- `fecha_inicio` es opcional; si no se envía, se asigna la fecha actual

## Ejemplos con curl

### Entrenadores

```bash
# Listar todos
curl -i http://localhost:8080/api/v1/entrenadores

# Obtener uno
curl -i http://localhost:8080/api/v1/entrenadores/1

# Crear
curl -i -X POST http://localhost:8080/api/v1/entrenadores \
  -H "Content-Type: application/json" \
  -d '{"nombre":"Sofia Castillo","especialidad":"Natación","capacidad_maxima":5,"carga_actual":1.0}'

# Actualizar
curl -i -X PUT http://localhost:8080/api/v1/entrenadores/1 \
  -H "Content-Type: application/json" \
  -d '{"nombre":"Carlos Mendoza","especialidad":"Velocidad y Fuerza","capacidad_maxima":12,"carga_actual":3.0}'

# Eliminar
curl -i -X DELETE http://localhost:8080/api/v1/entrenadores/3
```

### Atletas

```bash
# Listar todos
curl -i http://localhost:8080/api/v1/atletas

# Obtener uno
curl -i http://localhost:8080/api/v1/atletas/1

# Crear
curl -i -X POST http://localhost:8080/api/v1/atletas \
  -H "Content-Type: application/json" \
  -d '{"nombre":"Diana Flores","metodologia_objetivo":"Triple salto","peso":63.5,"coach_id":1}'

# Actualizar
curl -i -X PUT http://localhost:8080/api/v1/atletas/1 \
  -H "Content-Type: application/json" \
  -d '{"nombre":"Pedro Gómez","metodologia_objetivo":"Sprint 200m","peso":73.0,"coach_id":2}'

# Eliminar
curl -i -X DELETE http://localhost:8080/api/v1/atletas/5
```

### Ciclos de Entrenamiento

```bash
# Listar todos
curl -i http://localhost:8080/api/v1/ciclos

# Obtener uno
curl -i http://localhost:8080/api/v1/ciclos/1

# Crear (con fecha)
curl -i -X POST http://localhost:8080/api/v1/ciclos \
  -H "Content-Type: application/json" \
  -d '{"atleta_id":2,"estado":"activo","fecha_inicio":"2025-01-15T08:00:00Z"}'

# Crear (sin fecha — se asigna automáticamente)
curl -i -X POST http://localhost:8080/api/v1/ciclos \
  -H "Content-Type: application/json" \
  -d '{"atleta_id":3,"estado":"pendiente"}'

# Actualizar
curl -i -X PUT http://localhost:8080/api/v1/ciclos/1 \
  -H "Content-Type: application/json" \
  -d '{"atleta_id":1,"estado":"completado","fecha_inicio":"2025-01-01T00:00:00Z"}'

# Eliminar
curl -i -X DELETE http://localhost:8080/api/v1/ciclos/3
```

## Script de pruebas automáticas

El archivo `test_endpoints.sh` prueba todos los endpoints y validaciones de los tres recursos:

```bash
# Levantar el servidor en una terminal
go run ./cmd/atletismo-api

# Correr los tests en otra terminal
bash test_endpoints.sh
```

Cada prueba imprime el body de respuesta y el status HTTP para verificación rápida.
