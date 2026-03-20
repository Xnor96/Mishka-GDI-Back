package handler

import (
	"net/http"

	"github.com/Mishka-GDI-Back/application"
	"github.com/Mishka-GDI-Back/infrastructure/http/dto"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service application.AuthService
}

func NewAuthHandler(service application.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: "Datos inválidos", Error: err.Error()})
		return
	}
	tokenPair, err := h.service.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.AuthResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, dto.AuthResponse{
		Success:      true,
		Message:      "Inicio de sesión exitoso",
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    tokenPair.ExpiresIn,
		Username:     tokenPair.Username,
		Rol:          tokenPair.Rol,
	})
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: "Datos inválidos", Error: err.Error()})
		return
	}
	tokenPair, err := h.service.RefreshToken(req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.AuthResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, dto.AuthResponse{
		Success:     true,
		Message:     "Token renovado exitosamente",
		AccessToken: tokenPair.AccessToken,
		TokenType:   "Bearer",
		ExpiresIn:   tokenPair.ExpiresIn,
		Username:    tokenPair.Username,
		Rol:         tokenPair.Rol,
	})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	// JWT es stateless — el logout se maneja en el cliente descartando el token
	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Sesión cerrada exitosamente. Descarta el token en el cliente.",
	})
}
