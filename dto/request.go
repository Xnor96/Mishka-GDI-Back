package dto

// Categoria DTOs
type CreateCategoriaRequest struct {
	Nombre string `json:"nombre" binding:"required,min=1,max=100"`
}

type UpdateCategoriaRequest struct {
	Nombre string `json:"nombre" binding:"required,min=1,max=100"`
}

// Producto DTOs
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

// Entrada Producto DTOs
type CreateEntradaProductoRequest struct {
	IDProducto      int      `json:"id_producto" binding:"required"`
	FechaEntrada    string   `json:"fecha_entrada" binding:"required"`
	Cantidad        int      `json:"cantidad" binding:"required,min=1"`
	PrecioUnitario  *float64 `json:"precio_unitario"`
	Observaciones   string   `json:"observaciones"`
	UsuarioRegistro string   `json:"usuario_registro" binding:"required,max=100"`
}

// Salida Producto DTOs
type CreateSalidaProductoRequest struct {
	IDProducto      int    `json:"id_producto" binding:"required"`
	FechaSalida     string `json:"fecha_salida" binding:"required"`
	Cantidad        int    `json:"cantidad" binding:"required,min=1"`
	Observaciones   string `json:"observaciones"`
	UsuarioRegistro string `json:"usuario_registro" binding:"required,max=100"`
}
