package domain

import "time"

type Usuario struct {
	ID                 int
	Username           string
	Email              string
	PasswordHash       string
	Nombre             string
	Rol                string
	Activo             bool
	FechaCreacion      time.Time
	FechaActualizacion time.Time
}
