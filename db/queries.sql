
-- ===================== CicloEntrenamiento =====================

-- name: ListarCicloEntrenamiento :many
SELECT id, atletaID, estado, fechaInicio FROM categorias;

-- name: BuscarCicloEntrenamientoPorID :one
SELECT id, atletaID, estado, fechaInicio FROM cicloEntrenamiento
WHERE id = ?;

-- name: CrearCicloEntrenamiento :one
INSERT INTO cicloEntrenamiento (atletaID, estado, fechaInicio)
VALUES (?, ?, ?)
RETURNING id, atletaID, estado, fechaInicio;

-- name: ActualizarCicloEntrenamiento :one
UPDATE cicloEntrenamiento
SET atletaID = ?, estado = ?, fechaInicio = ?
WHERE id = ?
RETURNING id, atletaID, estado, fechaInicio;

-- name: BorrarCiclo :execrows
DELETE FROM cicloEntrenamiento WHERE id = ?;

-- ===================== EvaluacionCiclo =====================

-- name: ListarEvaluacionCiclo :many
SELECT id, evaluacionCiclo, nivelFatiga, comentarios, fechaEvaluacion FROM evaluacionCiclo;

-- name: BuscarEvaluacionCicloPorID :one
SELECT id, cicloEntrenamientoID, nivelFatiga, comentarios, fechaEvaluacion FROM evaluacionCiclo
WHERE id = ?;

-- name: CrearEvaluacionCiclo :one
INSERT INTO evaluacionCiclo (evaluacionCiclo, nivelFatiga, comentarios, fechaEvaluacion)
VALUES (?, ?, ?, ?)
RETURNING id, evaluacionCiclo, nivelFatiga, comentarios, fechaEvaluacion;

-- name: ActualizarEvaluacionCiclo :one
UPDATE evaluacionCiclo
SET cicloEntrenamientoID = ?, nivelFatiga = ?, comentarios = ?, fechaEvaluacion = ?
WHERE id = ?
RETURNING id, cicloEntrenamientoID, nivelFatiga, comentarios, fechaEvaluacion;

-- name: BorrarEvaluacionCiclo :execrows
DELETE FROM evaluacionCiclo WHERE id = ?;

-- ===================== Microciclo =====================

-- name: ListarMicrociclo :many
SELECT id, cicloEntrenamientoID, numeroSemana, enfoqueEspecifico, fechaInicio FROM microciclo;

-- name: BuscarMicrocicloPorID :one
SELECT id, cicloEntrenamientoID, numeroSemana, enfoqueEspecifico, fechaInicio FROM microciclo
WHERE id = ?;

-- name: CrearMicrociclo :one
INSERT INTO microciclo (cicloEntrenamientoID, numeroSemana, enfoqueEspecifico, fechaInicio)
VALUES (?, ?, ?, ?)
RETURNING id, cicloEntrenamientoID, numeroSemana, enfoqueEspecifico, fechaInicio;

-- name: ActualizarMicrociclo :one
UPDATE microciclo
SET cicloEntrenamientoID = ?, numeroSemana = ?, enfoqueEspecifico = ?, fechaInicio = ?
WHERE id = ?
RETURNING id, cicloEntrenamientoID, numeroSemana, enfoqueEspecifico, fechaInicio;

-- name: BorrarMicrociclo :execrows
DELETE FROM microciclo WHERE id = ?;