package middleware

import (
	"net/http"
	"strings"

	"github.com/Mishka-GDI-Back/application"
	"github.com/Mishka-GDI-Back/infrastructure/http/dto"
	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
				Success: false,
				Error:   "Token de autorización requerido",
			})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
				Success: false,
				Error:   "Formato de token inválido. Use: Bearer <token>",
			})
			return
		}

		claims, err := application.ValidateToken(parts[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
				Success: false,
				Error:   "Token inválido o expirado",
			})
			return
		}

		if claims.Type != "access" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
				Success: false,
				Error:   "Se requiere un access token",
			})
			return
		}

		// Guardar datos del usuario en el contexto
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("rol", claims.Rol)
		c.Next()
	}
}

func AdminRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		rol, exists := c.Get("rol")
		if !exists || rol != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, dto.Response{
				Success: false,
				Error:   "Acceso denegado. Se requiere rol de administrador",
			})
			return
		}
		c.Next()
	}
}
