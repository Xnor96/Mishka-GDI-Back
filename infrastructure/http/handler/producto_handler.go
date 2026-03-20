package handler

import (
	"net/http"
	"strconv"

	"github.com/Mishka-GDI-Back/application"
	"github.com/Mishka-GDI-Back/domain"
	"github.com/Mishka-GDI-Back/infrastructure/http/dto"
	"github.com/gin-gonic/gin"
)

type ProductoHandler struct {
	service application.ProductoService
}

func NewProductoHandler(service application.ProductoService) *ProductoHandler {
	return &ProductoHandler{service: service}
}

func (h *ProductoHandler) GetAll(c *gin.Context) {
	productos, err := h.service.GetAll()
	if err != nil {
		handleDomainError(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.ProductosResponse{
		Success:    true,
		Message:    "Productos obtenidos exitosamente",
		Data:       dto.ProductosToResponse(productos),
		TotalCount: len(productos),
	})
}

func (h *ProductoHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: "ID inválido", Error: "El ID debe ser un número entero"})
		return
	}
	producto, err := h.service.GetByID(id)
	if err != nil {
		handleDomainError(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Producto encontrado",
		Data:    dto.ProductoToResponse(producto),
	})
}

func (h *ProductoHandler) Create(c *gin.Context) {
	var req dto.CreateProductoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: "Datos inválidos", Error: err.Error()})
		return
	}
	producto := &domain.Producto{
		Codigo:         req.Codigo,
		Nombre:         req.Nombre,
		IDCategoria:    req.IDCategoria,
		UnidadMedida:   req.UnidadMedida,
		PrecioUnitario: req.PrecioUnitario,
		StockActual:    req.StockActual,
		StockInicial:   req.StockInicial,
	}
	result, err := h.service.Create(producto)
	if err != nil {
		handleDomainError(c, err)
		return
	}
	c.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Message: "Producto creado exitosamente",
		Data:    dto.ProductoToResponse(result),
	})
}

func (h *ProductoHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: "ID inválido", Error: "El ID debe ser un número entero"})
		return
	}
	var req dto.UpdateProductoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: "Datos inválidos", Error: err.Error()})
		return
	}
	producto := &domain.Producto{
		Codigo:         req.Codigo,
		Nombre:         req.Nombre,
		IDCategoria:    req.IDCategoria,
		UnidadMedida:   req.UnidadMedida,
		PrecioUnitario: req.PrecioUnitario,
		StockActual:    req.StockActual,
		StockInicial:   req.StockInicial,
	}
	result, err := h.service.Update(id, producto)
	if err != nil {
		handleDomainError(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Producto actualizado exitosamente",
		Data:    dto.ProductoToResponse(result),
	})
}

func (h *ProductoHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: "ID inválido", Error: "El ID debe ser un número entero"})
		return
	}
	if err := h.service.Delete(id); err != nil {
		handleDomainError(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.Response{Success: true, Message: "Producto eliminado exitosamente"})
}

func (h *ProductoHandler) GetStockBajo(c *gin.Context) {
	limite, _ := strconv.Atoi(c.DefaultQuery("limite", "5"))
	productos, err := h.service.GetStockBajo(limite)
	if err != nil {
		handleDomainError(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.ProductosResponse{
		Success:    true,
		Message:    "Productos con stock bajo",
		Data:       dto.ProductosToResponse(productos),
		TotalCount: len(productos),
	})
}

func (h *ProductoHandler) Search(c *gin.Context) {
	termino := c.Query("q")
	productos, err := h.service.Search(termino)
	if err != nil {
		handleDomainError(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.ProductosResponse{
		Success:    true,
		Message:    "Resultados de búsqueda",
		Data:       dto.ProductosToResponse(productos),
		TotalCount: len(productos),
	})
}
