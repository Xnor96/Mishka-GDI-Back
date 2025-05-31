package dto

import (
	"time"

	"github.com/Mishka-GDI-Back/domain"
)

// Response genérica
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// Categoria Response
type CategoriaResponse struct {
	ID                 int       `json:"id_categoria"`
	Nombre             string    `json:"nombre"`
	FechaCreacion      time.Time `json:"fecha_creacion"`
	FechaActualizacion time.Time `json:"fecha_actualizacion"`
}

// Producto Response
type ProductoResponse struct {
	ID                 int       `json:"id_producto"`
	Codigo             string    `json:"codigo"`
	Nombre             string    `json:"nombre"`
	IDCategoria        *int      `json:"id_categoria"`
	UnidadMedida       string    `json:"unidad_medida"`
	PrecioUnitario     float64   `json:"precio_unitario"`
	StockActual        int       `json:"stock_actual"`
	StockInicial       int       `json:"stock_inicial"`
	FechaCreacion      time.Time `json:"fecha_creacion"`
	FechaActualizacion time.Time `json:"fecha_actualizacion"`
}

// Entrada Producto Response
type EntradaProductoResponse struct {
	ID                 int       `json:"id_entrada"`
	IDProducto         int       `json:"id_producto"`
	FechaEntrada       time.Time `json:"fecha_entrada"`
	Cantidad           int       `json:"cantidad"`
	PrecioUnitario     *float64  `json:"precio_unitario"`
	Observaciones      string    `json:"observaciones"`
	UsuarioRegistro    string    `json:"usuario_registro"`
	FechaCreacion      time.Time `json:"fecha_creacion"`
	FechaActualizacion time.Time `json:"fecha_actualizacion"`
}

// Salida Producto Response
type SalidaProductoResponse struct {
	ID                 int       `json:"id_salida"`
	IDProducto         int       `json:"id_producto"`
	FechaSalida        time.Time `json:"fecha_salida"`
	Cantidad           int       `json:"cantidad"`
	Observaciones      string    `json:"observaciones"`
	UsuarioRegistro    string    `json:"usuario_registro"`
	FechaCreacion      time.Time `json:"fecha_creacion"`
	FechaActualizacion time.Time `json:"fecha_actualizacion"`
}

// Lista de respuestas
type CategoriasResponse struct {
	Success    bool                `json:"success"`
	Message    string              `json:"message"`
	Data       []CategoriaResponse `json:"data"`
	TotalCount int                 `json:"total_count"`
}

type ProductosResponse struct {
	Success    bool               `json:"success"`
	Message    string             `json:"message"`
	Data       []ProductoResponse `json:"data"`
	TotalCount int                `json:"total_count"`
}

type EntradasResponse struct {
	Success    bool                      `json:"success"`
	Message    string                    `json:"message"`
	Data       []EntradaProductoResponse `json:"data"`
	TotalCount int                       `json:"total_count"`
}

type SalidasResponse struct {
	Success    bool                     `json:"success"`
	Message    string                   `json:"message"`
	Data       []SalidaProductoResponse `json:"data"`
	TotalCount int                      `json:"total_count"`
}

// Helper functions para convertir domain models a responses
func CategoriaToResponse(categoria *domain.Categoria) CategoriaResponse {
	return CategoriaResponse{
		ID:                 categoria.ID,
		Nombre:             categoria.Nombre,
		FechaCreacion:      categoria.FechaCreacion,
		FechaActualizacion: categoria.FechaActualizacion,
	}
}

func ProductoToResponse(producto *domain.Producto) ProductoResponse {
	return ProductoResponse{
		ID:                 producto.ID,
		Codigo:             producto.Codigo,
		Nombre:             producto.Nombre,
		IDCategoria:        producto.IDCategoria,
		UnidadMedida:       producto.UnidadMedida,
		PrecioUnitario:     producto.PrecioUnitario,
		StockActual:        producto.StockActual,
		StockInicial:       producto.StockInicial,
		FechaCreacion:      producto.FechaCreacion,
		FechaActualizacion: producto.FechaActualizacion,
	}
}

func EntradaProductoToResponse(entrada *domain.EntradaProducto) EntradaProductoResponse {
	return EntradaProductoResponse{
		ID:                 entrada.ID,
		IDProducto:         entrada.IDProducto,
		FechaEntrada:       entrada.FechaEntrada,
		Cantidad:           entrada.Cantidad,
		PrecioUnitario:     entrada.PrecioUnitario,
		Observaciones:      entrada.Observaciones,
		UsuarioRegistro:    entrada.UsuarioRegistro,
		FechaCreacion:      entrada.FechaCreacion,
		FechaActualizacion: entrada.FechaActualizacion,
	}
}

func SalidaProductoToResponse(salida *domain.SalidaProducto) SalidaProductoResponse {
	return SalidaProductoResponse{
		ID:                 salida.ID,
		IDProducto:         salida.IDProducto,
		FechaSalida:        salida.FechaSalida,
		Cantidad:           salida.Cantidad,
		Observaciones:      salida.Observaciones,
		UsuarioRegistro:    salida.UsuarioRegistro,
		FechaCreacion:      salida.FechaCreacion,
		FechaActualizacion: salida.FechaActualizacion,
	}
}

// Helper functions para convertir listas
func CategoriasToResponse(categorias []domain.Categoria) []CategoriaResponse {
	responses := make([]CategoriaResponse, len(categorias))
	for i, categoria := range categorias {
		responses[i] = CategoriaToResponse(&categoria)
	}
	return responses
}

func ProductosToResponse(productos []domain.Producto) []ProductoResponse {
	responses := make([]ProductoResponse, len(productos))
	for i, producto := range productos {
		responses[i] = ProductoToResponse(&producto)
	}
	return responses
}

func EntradasToResponse(entradas []domain.EntradaProducto) []EntradaProductoResponse {
	responses := make([]EntradaProductoResponse, len(entradas))
	for i, entrada := range entradas {
		responses[i] = EntradaProductoToResponse(&entrada)
	}
	return responses
}

func SalidasToResponse(salidas []domain.SalidaProducto) []SalidaProductoResponse {
	responses := make([]SalidaProductoResponse, len(salidas))
	for i, salida := range salidas {
		responses[i] = SalidaProductoToResponse(&salida)
	}
	return responses
}
