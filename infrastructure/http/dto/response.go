package dto

import (
	"time"

	"github.com/Mishka-GDI-Back/domain"
)

// =============================================
// Response genérica
// =============================================

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// =============================================
// Categoria Response
// =============================================

type CategoriaResponse struct {
	ID                 int       `json:"id_categoria"`
	Nombre             string    `json:"nombre"`
	FechaCreacion      time.Time `json:"fecha_creacion"`
	FechaActualizacion time.Time `json:"fecha_actualizacion"`
}

type CategoriasResponse struct {
	Success    bool                `json:"success"`
	Message    string              `json:"message"`
	Data       []CategoriaResponse `json:"data"`
	TotalCount int                 `json:"total_count"`
}

// =============================================
// Producto Response
// =============================================

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

type ProductosResponse struct {
	Success    bool               `json:"success"`
	Message    string             `json:"message"`
	Data       []ProductoResponse `json:"data"`
	TotalCount int                `json:"total_count"`
}

// =============================================
// Entrada Producto Response (con datos del producto)
// =============================================

type EntradaProductoResponse struct {
	ID                 int       `json:"id_entrada"`
	IDProducto         int       `json:"id_producto"`
	CodigoProducto     string    `json:"codigo_producto"`
	NombreProducto     string    `json:"nombre_producto"`
	NombreCategoria    string    `json:"nombre_categoria"`
	FechaEntrada       time.Time `json:"fecha_entrada"`
	Cantidad           int       `json:"cantidad"`
	PrecioUnitario     *float64  `json:"precio_unitario"`
	Observaciones      string    `json:"observaciones"`
	UsuarioRegistro    string    `json:"usuario_registro"`
	FechaCreacion      time.Time `json:"fecha_creacion"`
	FechaActualizacion time.Time `json:"fecha_actualizacion"`
}

type EntradasResponse struct {
	Success    bool                      `json:"success"`
	Message    string                    `json:"message"`
	Data       []EntradaProductoResponse `json:"data"`
	TotalCount int                       `json:"total_count"`
}

// =============================================
// Salida Producto Response (con datos del producto + campos de venta)
// =============================================

type SalidaProductoResponse struct {
	ID                 int       `json:"id_salida"`
	IDProducto         int       `json:"id_producto"`
	CodigoProducto     string    `json:"codigo_producto"`
	NombreProducto     string    `json:"nombre_producto"`
	NombreCategoria    string    `json:"nombre_categoria"`
	FechaSalida        time.Time `json:"fecha_salida"`
	Cantidad           int       `json:"cantidad"`
	PrecioVenta        float64   `json:"precio_venta"`
	Descuento          float64   `json:"descuento"`
	Total              float64   `json:"total"`
	LugarVenta         string    `json:"lugar_venta"`
	TipoPago           string    `json:"tipo_pago"`
	Observaciones      string    `json:"observaciones"`
	UsuarioRegistro    string    `json:"usuario_registro"`
	FechaCreacion      time.Time `json:"fecha_creacion"`
	FechaActualizacion time.Time `json:"fecha_actualizacion"`
}

type SalidasResponse struct {
	Success    bool                     `json:"success"`
	Message    string                   `json:"message"`
	Data       []SalidaProductoResponse `json:"data"`
	TotalCount int                      `json:"total_count"`
}

// =============================================
// Control Diario Response
// =============================================

type ControlDiarioResponse struct {
	ID                 int       `json:"id_control"`
	Fecha              time.Time `json:"fecha"`
	Descripcion        string    `json:"descripcion"`
	MontoEntrada       float64   `json:"monto_entrada"`
	MontoSalida        float64   `json:"monto_salida"`
	Observaciones      string    `json:"observaciones"`
	EsVerbena          bool      `json:"es_verbena"`
	UsuarioRegistro    string    `json:"usuario_registro"`
	FechaCreacion      time.Time `json:"fecha_creacion"`
	FechaActualizacion time.Time `json:"fecha_actualizacion"`
}

type ControlDiariosResponse struct {
	Success      bool                    `json:"success"`
	Message      string                  `json:"message"`
	Data         []ControlDiarioResponse `json:"data"`
	TotalCount   int                     `json:"total_count"`
	TotalEntrada float64                 `json:"total_entrada"`
	TotalSalida  float64                 `json:"total_salida"`
	Balance      float64                 `json:"balance"`
}

// =============================================
// Resumen Mensual Response
// =============================================

type ResumenMensualResponse struct {
	ID                   int       `json:"id_resumen"`
	Mes                  int       `json:"mes"`
	NombreMes            string    `json:"nombre_mes"`
	Anio                 int       `json:"anio"`
	TotalIngresos        float64   `json:"total_ingresos"`
	TotalGastosFijos     float64   `json:"total_gastos_fijos"`
	TotalGastosVariables float64   `json:"total_gastos_variables"`
	Balance              float64   `json:"balance"`
	Observaciones        string    `json:"observaciones"`
	FechaGeneracion      time.Time `json:"fecha_generacion"`
	FechaActualizacion   time.Time `json:"fecha_actualizacion"`
}

// =============================================
// Auth Response
// =============================================

type AuthResponse struct {
	Success      bool   `json:"success"`
	Message      string `json:"message"`
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	TokenType    string `json:"token_type,omitempty"`
	ExpiresIn    int64  `json:"expires_in,omitempty"`
	Username     string `json:"username,omitempty"`
	Rol          string `json:"rol,omitempty"`
}

// =============================================
// Reportes Responses
// =============================================

type ReporteInventarioItem struct {
	IDProducto     int     `json:"id_producto"`
	Codigo         string  `json:"codigo"`
	Nombre         string  `json:"nombre"`
	Categoria      string  `json:"categoria"`
	UnidadMedida   string  `json:"unidad_medida"`
	PrecioUnitario float64 `json:"precio_unitario"`
	StockActual    int     `json:"stock_actual"`
	ValorTotal     float64 `json:"valor_total"`
}

type ReporteMovimientoItem struct {
	Fecha      time.Time `json:"fecha"`
	Tipo       string    `json:"tipo"`
	Codigo     string    `json:"codigo"`
	Nombre     string    `json:"nombre"`
	Categoria  string    `json:"categoria"`
	Cantidad   int       `json:"cantidad"`
	Precio     float64   `json:"precio"`
	Total      float64   `json:"total"`
	LugarVenta string    `json:"lugar_venta,omitempty"`
	TipoPago   string    `json:"tipo_pago,omitempty"`
}

type ReporteProductoVendidoItem struct {
	IDProducto    int     `json:"id_producto"`
	Codigo        string  `json:"codigo"`
	Nombre        string  `json:"nombre"`
	Categoria     string  `json:"categoria"`
	TotalVendido  int     `json:"total_vendido"`
	TotalIngresos float64 `json:"total_ingresos"`
}

type ReporteValoracionItem struct {
	IDCategoria     int     `json:"id_categoria"`
	NombreCategoria string  `json:"nombre_categoria"`
	TotalProductos  int     `json:"total_productos"`
	TotalUnidades   int     `json:"total_unidades"`
	ValorTotal      float64 `json:"valor_total"`
}

// =============================================
// Alertas Response
// =============================================

type AlertaStockBajoItem struct {
	IDProducto     int     `json:"id_producto"`
	Codigo         string  `json:"codigo"`
	Nombre         string  `json:"nombre"`
	Categoria      string  `json:"categoria"`
	StockActual    int     `json:"stock_actual"`
	StockInicial   int     `json:"stock_inicial"`
	PrecioUnitario float64 `json:"precio_unitario"`
}

type AlertasResponse struct {
	Success     bool                  `json:"success"`
	Message     string                `json:"message"`
	StockBajo   []AlertaStockBajoItem `json:"stock_bajo"`
	TotalAlerts int                   `json:"total_alertas"`
}

// =============================================
// Helper functions: domain → response
// =============================================

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

func EntradaConProductoToResponse(entrada *domain.EntradaConProducto) EntradaProductoResponse {
	return EntradaProductoResponse{
		ID:                 entrada.ID,
		IDProducto:         entrada.IDProducto,
		CodigoProducto:     entrada.CodigoProducto,
		NombreProducto:     entrada.NombreProducto,
		NombreCategoria:    entrada.NombreCategoria,
		FechaEntrada:       entrada.FechaEntrada,
		Cantidad:           entrada.Cantidad,
		PrecioUnitario:     entrada.PrecioUnitario,
		Observaciones:      entrada.Observaciones,
		UsuarioRegistro:    entrada.UsuarioRegistro,
		FechaCreacion:      entrada.FechaCreacion,
		FechaActualizacion: entrada.FechaActualizacion,
	}
}

func SalidaConProductoToResponse(salida *domain.SalidaConProducto) SalidaProductoResponse {
	return SalidaProductoResponse{
		ID:                 salida.ID,
		IDProducto:         salida.IDProducto,
		CodigoProducto:     salida.CodigoProducto,
		NombreProducto:     salida.NombreProducto,
		NombreCategoria:    salida.NombreCategoria,
		FechaSalida:        salida.FechaSalida,
		Cantidad:           salida.Cantidad,
		PrecioVenta:        salida.PrecioVenta,
		Descuento:          salida.Descuento,
		Total:              salida.Total,
		LugarVenta:         salida.LugarVenta,
		TipoPago:           salida.TipoPago,
		Observaciones:      salida.Observaciones,
		UsuarioRegistro:    salida.UsuarioRegistro,
		FechaCreacion:      salida.FechaCreacion,
		FechaActualizacion: salida.FechaActualizacion,
	}
}

func ControlDiarioToResponse(c *domain.ControlDiario) ControlDiarioResponse {
	return ControlDiarioResponse{
		ID:                 c.ID,
		Fecha:              c.Fecha,
		Descripcion:        c.Descripcion,
		MontoEntrada:       c.MontoEntrada,
		MontoSalida:        c.MontoSalida,
		Observaciones:      c.Observaciones,
		EsVerbena:          c.EsVerbena,
		UsuarioRegistro:    c.UsuarioRegistro,
		FechaCreacion:      c.FechaCreacion,
		FechaActualizacion: c.FechaActualizacion,
	}
}

var nombresMeses = map[int]string{
	1: "Enero", 2: "Febrero", 3: "Marzo", 4: "Abril",
	5: "Mayo", 6: "Junio", 7: "Julio", 8: "Agosto",
	9: "Septiembre", 10: "Octubre", 11: "Noviembre", 12: "Diciembre",
}

func ResumenMensualToResponse(r *domain.ResumenMensual) ResumenMensualResponse {
	return ResumenMensualResponse{
		ID:                   r.ID,
		Mes:                  r.Mes,
		NombreMes:            nombresMeses[r.Mes],
		Anio:                 r.Anio,
		TotalIngresos:        r.TotalIngresos,
		TotalGastosFijos:     r.TotalGastosFijos,
		TotalGastosVariables: r.TotalGastosVariables,
		Balance:              r.Balance,
		Observaciones:        r.Observaciones,
		FechaGeneracion:      r.FechaGeneracion,
		FechaActualizacion:   r.FechaActualizacion,
	}
}

// =============================================
// Helper functions: domain reportes → response
// =============================================

func ReporteInventarioToResponse(item *domain.ReporteInventarioItem) ReporteInventarioItem {
	return ReporteInventarioItem{
		IDProducto:     item.IDProducto,
		Codigo:         item.Codigo,
		Nombre:         item.Nombre,
		Categoria:      item.Categoria,
		UnidadMedida:   item.UnidadMedida,
		PrecioUnitario: item.PrecioUnitario,
		StockActual:    item.StockActual,
		ValorTotal:     item.ValorTotal,
	}
}

func ReporteMovimientoToResponse(item *domain.ReporteMovimiento) ReporteMovimientoItem {
	return ReporteMovimientoItem{
		Fecha:      item.Fecha,
		Tipo:       item.Tipo,
		Codigo:     item.Codigo,
		Nombre:     item.Nombre,
		Categoria:  item.Categoria,
		Cantidad:   item.Cantidad,
		Precio:     item.Precio,
		Total:      item.Total,
		LugarVenta: item.LugarVenta,
		TipoPago:   item.TipoPago,
	}
}

func ReporteProductoVendidoToResponse(item *domain.ReporteProductoVendido) ReporteProductoVendidoItem {
	return ReporteProductoVendidoItem{
		IDProducto:    item.IDProducto,
		Codigo:        item.Codigo,
		Nombre:        item.Nombre,
		Categoria:     item.Categoria,
		TotalVendido:  item.TotalVendido,
		TotalIngresos: item.TotalIngresos,
	}
}

func ReporteValoracionToResponse(item *domain.ReporteValoracion) ReporteValoracionItem {
	return ReporteValoracionItem{
		IDCategoria:     item.IDCategoria,
		NombreCategoria: item.NombreCategoria,
		TotalProductos:  item.TotalProductos,
		TotalUnidades:   item.TotalUnidades,
		ValorTotal:      item.ValorTotal,
	}
}

func AlertaStockBajoToResponse(item *domain.AlertaStockBajo) AlertaStockBajoItem {
	return AlertaStockBajoItem{
		IDProducto:     item.IDProducto,
		Codigo:         item.Codigo,
		Nombre:         item.Nombre,
		Categoria:      item.Categoria,
		StockActual:    item.StockActual,
		StockInicial:   item.StockInicial,
		PrecioUnitario: item.PrecioUnitario,
	}
}

// =============================================
// Resumen Producto Response
// =============================================

type ResumenProductoResponse struct {
	IDProducto    int     `json:"id_producto"`
	Mes           int     `json:"mes"`
	Anio          int     `json:"anio"`
	TotalEntradas int     `json:"total_entradas"`
	MontoEntradas float64 `json:"monto_entradas"`
	TotalSalidas  int     `json:"total_salidas"`
	MontoSalidas  float64 `json:"monto_salidas"`
}

func ResumenProductoToResponse(r *domain.ResumenProducto) ResumenProductoResponse {
	return ResumenProductoResponse{
		IDProducto:    r.IDProducto,
		Mes:           r.Mes,
		Anio:          r.Anio,
		TotalEntradas: r.TotalEntradas,
		MontoEntradas: r.MontoEntradas,
		TotalSalidas:  r.TotalSalidas,
		MontoSalidas:  r.MontoSalidas,
	}
}

// =============================================
// Helper functions: listas
// =============================================

func CategoriasToResponse(categorias []domain.Categoria) []CategoriaResponse {
	responses := make([]CategoriaResponse, len(categorias))
	for i, c := range categorias {
		responses[i] = CategoriaToResponse(&c)
	}
	return responses
}

func ProductosToResponse(productos []domain.Producto) []ProductoResponse {
	responses := make([]ProductoResponse, len(productos))
	for i, p := range productos {
		responses[i] = ProductoToResponse(&p)
	}
	return responses
}

func EntradasConProductoToResponse(entradas []domain.EntradaConProducto) []EntradaProductoResponse {
	responses := make([]EntradaProductoResponse, len(entradas))
	for i, e := range entradas {
		responses[i] = EntradaConProductoToResponse(&e)
	}
	return responses
}

func SalidasConProductoToResponse(salidas []domain.SalidaConProducto) []SalidaProductoResponse {
	responses := make([]SalidaProductoResponse, len(salidas))
	for i, s := range salidas {
		responses[i] = SalidaConProductoToResponse(&s)
	}
	return responses
}

func ControlDiariosToResponse(controles []domain.ControlDiario) []ControlDiarioResponse {
	responses := make([]ControlDiarioResponse, len(controles))
	for i, c := range controles {
		responses[i] = ControlDiarioToResponse(&c)
	}
	return responses
}

func ReportesInventarioToResponse(items []domain.ReporteInventarioItem) []ReporteInventarioItem {
	responses := make([]ReporteInventarioItem, len(items))
	for i, item := range items {
		responses[i] = ReporteInventarioToResponse(&item)
	}
	return responses
}

func ReportesMovimientoToResponse(items []domain.ReporteMovimiento) []ReporteMovimientoItem {
	responses := make([]ReporteMovimientoItem, len(items))
	for i, item := range items {
		responses[i] = ReporteMovimientoToResponse(&item)
	}
	return responses
}

func ReportesProductoVendidoToResponse(items []domain.ReporteProductoVendido) []ReporteProductoVendidoItem {
	responses := make([]ReporteProductoVendidoItem, len(items))
	for i, item := range items {
		responses[i] = ReporteProductoVendidoToResponse(&item)
	}
	return responses
}

func ReportesValoracionToResponse(items []domain.ReporteValoracion) []ReporteValoracionItem {
	responses := make([]ReporteValoracionItem, len(items))
	for i, item := range items {
		responses[i] = ReporteValoracionToResponse(&item)
	}
	return responses
}

func AlertasStockBajoToResponse(items []domain.AlertaStockBajo) []AlertaStockBajoItem {
	responses := make([]AlertaStockBajoItem, len(items))
	for i, item := range items {
		responses[i] = AlertaStockBajoToResponse(&item)
	}
	return responses
}
