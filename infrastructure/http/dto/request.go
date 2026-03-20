package dto

// =============================================
// Categoria DTOs
// =============================================

type CreateCategoriaRequest struct {
	Nombre string `json:"nombre" binding:"required,min=1,max=100"`
}

type UpdateCategoriaRequest struct {
	Nombre string `json:"nombre" binding:"required,min=1,max=100"`
}

// =============================================
// Producto DTOs
// =============================================

type CreateProductoRequest struct {
	Codigo         string  `json:"codigo" binding:"required,min=1,max=50"`
	Nombre         string  `json:"nombre" binding:"required,min=1,max=200"`
	IDCategoria    *int    `json:"id_categoria"`
	UnidadMedida   string  `json:"unidad_medida" binding:"max=20"`
	PrecioUnitario float64 `json:"precio_unitario" binding:"min=0"`
	StockActual    int     `json:"stock_actual" binding:"min=0"`
	StockInicial   int     `json:"stock_inicial" binding:"min=0"`
}

type UpdateProductoRequest struct {
	Codigo         string  `json:"codigo" binding:"required,min=1,max=50"`
	Nombre         string  `json:"nombre" binding:"required,min=1,max=200"`
	IDCategoria    *int    `json:"id_categoria"`
	UnidadMedida   string  `json:"unidad_medida" binding:"max=20"`
	PrecioUnitario float64 `json:"precio_unitario" binding:"min=0"`
	StockActual    int     `json:"stock_actual" binding:"min=0"`
	StockInicial   int     `json:"stock_inicial" binding:"min=0"`
}

// =============================================
// Entrada Producto DTOs
// =============================================

type CreateEntradaProductoRequest struct {
	IDProducto      int      `json:"id_producto" binding:"required"`
	FechaEntrada    string   `json:"fecha_entrada" binding:"required"`
	Cantidad        int      `json:"cantidad" binding:"required,min=1"`
	PrecioUnitario  *float64 `json:"precio_unitario"`
	Observaciones   string   `json:"observaciones"`
	UsuarioRegistro string   `json:"usuario_registro" binding:"required,max=100"`
}

// =============================================
// Salida Producto DTOs
// =============================================

type CreateSalidaProductoRequest struct {
	IDProducto      int     `json:"id_producto" binding:"required"`
	FechaSalida     string  `json:"fecha_salida" binding:"required"`
	Cantidad        int     `json:"cantidad" binding:"required,min=1"`
	PrecioVenta     float64 `json:"precio_venta" binding:"min=0"`
	Descuento       float64 `json:"descuento" binding:"min=0"`
	LugarVenta      string  `json:"lugar_venta" binding:"max=100"`
	TipoPago        string  `json:"tipo_pago" binding:"max=50"`
	Observaciones   string  `json:"observaciones"`
	UsuarioRegistro string  `json:"usuario_registro" binding:"required,max=100"`
}

// =============================================
// Control Diario DTOs
// =============================================

type CreateControlDiarioRequest struct {
	Fecha           string  `json:"fecha" binding:"required"`
	Descripcion     string  `json:"descripcion" binding:"required,min=1"`
	MontoEntrada    float64 `json:"monto_entrada" binding:"min=0"`
	MontoSalida     float64 `json:"monto_salida" binding:"min=0"`
	Observaciones   string  `json:"observaciones"`
	EsVerbena       bool    `json:"es_verbena"`
	UsuarioRegistro string  `json:"usuario_registro" binding:"required,max=100"`
}

// =============================================
// Resumen Mensual DTOs
// =============================================

type GenerarResumenRequest struct {
	Mes                  int     `json:"mes" binding:"required,min=1,max=12"`
	Anio                 int     `json:"anio" binding:"required,min=2020"`
	TotalIngresos        float64 `json:"total_ingresos" binding:"min=0"`
	TotalGastosFijos     float64 `json:"total_gastos_fijos" binding:"min=0"`
	TotalGastosVariables float64 `json:"total_gastos_variables" binding:"min=0"`
	Observaciones        string  `json:"observaciones"`
}

// =============================================
// Auth DTOs
// =============================================

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// =============================================
// Alertas DTOs
// =============================================

type ConfigurarAlertaRequest struct {
	LimiteStockBajo int `json:"limite_stock_bajo" binding:"required,min=1"`
}
