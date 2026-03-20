package handler

import (
	"errors"
	"net/http"

	"github.com/Mishka-GDI-Back/domain"
	"github.com/Mishka-GDI-Back/infrastructure/http/dto"
	"github.com/gin-gonic/gin"
)

// handleDomainError mapea errores del dominio a respuestas HTTP apropiadas
func handleDomainError(c *gin.Context, err error) {
	var notFound *domain.ErrNotFound
	var validation *domain.ErrValidation
	var duplicate *domain.ErrDuplicate
	var insuffStock *domain.ErrInsufficientStock

	switch {
	case errors.As(err, &notFound):
		c.JSON(http.StatusNotFound, dto.Response{
			Success: false,
			Message: "Recurso no encontrado",
			Error:   err.Error(),
		})
	case errors.As(err, &validation):
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Datos inválidos",
			Error:   err.Error(),
		})
	case errors.As(err, &duplicate):
		c.JSON(http.StatusConflict, dto.Response{
			Success: false,
			Message: "Recurso duplicado",
			Error:   err.Error(),
		})
	case errors.As(err, &insuffStock):
		c.JSON(http.StatusConflict, dto.Response{
			Success: false,
			Message: "Stock insuficiente",
			Error:   err.Error(),
		})
	default:
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "Error interno del servidor",
			Error:   err.Error(),
		})
	}
}
