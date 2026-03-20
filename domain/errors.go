package domain

import "fmt"

// ErrNotFound indica que el recurso no fue encontrado
type ErrNotFound struct {
	Entity string
	ID     interface{}
}

func (e *ErrNotFound) Error() string {
	return fmt.Sprintf("%s no encontrado(a) con ID: %v", e.Entity, e.ID)
}

// ErrValidation indica un error de validación
type ErrValidation struct {
	Field   string
	Message string
}

func (e *ErrValidation) Error() string {
	return fmt.Sprintf("validación fallida en campo '%s': %s", e.Field, e.Message)
}

// ErrDuplicate indica que ya existe un recurso con ese identificador
type ErrDuplicate struct {
	Entity string
	Field  string
	Value  string
}

func (e *ErrDuplicate) Error() string {
	return fmt.Sprintf("ya existe %s con %s: %s", e.Entity, e.Field, e.Value)
}

// ErrInsufficientStock indica stock insuficiente
type ErrInsufficientStock struct {
	ProductoID   int
	StockActual  int
	CantidadReq  int
}

func (e *ErrInsufficientStock) Error() string {
	return fmt.Sprintf("stock insuficiente: disponible=%d, solicitado=%d", e.StockActual, e.CantidadReq)
}
