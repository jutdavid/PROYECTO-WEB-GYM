package service

import "errors"

var (
	ErrNombreVacio            = errors.New("el campo nombre es obligatorio")
	ErrPrecioNegativo         = errors.New("el precio no puede ser negativo")
	ErrNoEncontrado           = errors.New("recurso no encontrado")
	ErrEmailEnUso             = errors.New("el email ya está en uso")
	ErrCredencialesInvalidas  = errors.New("Email o contraseña incorrectos")
	ErrEstadoVacio            = errors.New("el campo nombre es obligatorio")
	ErrComentarioVacio        = errors.New("el campo nombre es obligatorio")
	ErrEnfoqueEspecificoVacio = errors.New("el campo nombre es obligatorio")
)
