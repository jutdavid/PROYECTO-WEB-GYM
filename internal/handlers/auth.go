package handlers

import (
	"encoding/json"
	"net/http"
)

type credenciales struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *Server) Registrar(w http.ResponseWriter, r *http.Request) {
	var creds credenciales
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		ResponderError(w, http.StatusBadRequest, "JSON invalido: "+err.Error())
		return
	}
	usuario, err := s.Auth.Registrar(creds.Email, creds.Password)
	if err != nil {
		ResponderError(w, statusDeError(err), err.Error())
		return
	}
	ResponderJSON(w, http.StatusCreated, usuario)
}

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	var creds credenciales
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		ResponderError(w, http.StatusBadRequest, "JSON invalido: "+err.Error())
		return
	}
	token, err := s.Auth.Login(creds.Email, creds.Password)
	if err != nil {
		ResponderError(w, statusDeError(err), err.Error())
		return
	}
	ResponderJSON(w, http.StatusOK, map[string]string{"token": token})
}
