package handlers

import (
	"atletismo-api/internal/service"
)

type Server struct {
	//Atleta             *service.AtletaService
	//Certificacion      *service.CertificacionService
	CicloEntrenamiento *service.CicloEntrenamientoService
	//Entrenador         *service.EntrenadorService
	EvaluacionCiclo *service.EvaluacionCicloService
	//HorarioAtencion    *service.HorarioAtencionService
	//Lesion             *service.LesionService
	//MetricaFisica      *service.MetricaFisicaService
	Microciclo *service.MicrocicloService
	Auth       *service.AuthService
}

// func NewServer(atleta *service.AtletaService, certificacion *service.CertificacionService, cicloEntrenamiento *service.CicloEntrenamientoService, entrenador *service.EntrenadorService, evaluacionCiclo *service.EvaluacionCicloService, horarioAtencion *service.HorarioAtencionService, lesion *service.LesionService, metricaFisica *service.MetricaFisicaService, microciclo *service.MicrocicloService, auth *service.AuthService) *Server {
func NewServer(cicloEntrenamiento *service.CicloEntrenamientoService, evaluacionCiclo *service.EvaluacionCicloService, microciclo *service.MicrocicloService, auth *service.AuthService) *Server {
	return &Server{
		//Atleta:             atleta,
		//Certificacion:      certificacion,
		CicloEntrenamiento: cicloEntrenamiento,
		//Entrenador:         entrenador,
		EvaluacionCiclo: evaluacionCiclo,
		//HorarioAtencion:    horarioAtencion,
		//Lesion:             lesion,
		//MetricaFisica:      metricaFisica,
		Microciclo: microciclo,
		Auth:       auth,
	}
}
