package domain

// CategoriaRepository define el puerto de persistencia para categorías
type CategoriaRepository interface {
	GetAll() ([]Categoria, error)
	GetByID(id int) (*Categoria, error)
	Create(categoria *Categoria) error
	Update(categoria *Categoria) error
	Delete(id int) error
}

// ProductoRepository define el puerto de persistencia para productos
type ProductoRepository interface {
	GetAll() ([]Producto, error)
	GetByID(id int) (*Producto, error)
	GetByCodigo(codigo string) (*Producto, error)
	Create(producto *Producto) error
	Update(producto *Producto) error
	Delete(id int) error
	GetStockBajo(limite int) ([]Producto, error)
	Search(termino string) ([]Producto, error)
}

// EntradaProductoRepository define el puerto de persistencia para entradas
type EntradaProductoRepository interface {
	GetAll() ([]EntradaConProducto, error)
	GetByID(id int) (*EntradaConProducto, error)
	GetByProductoID(productoID int) ([]EntradaConProducto, error)
	GetByFecha(fecha string) ([]EntradaConProducto, error)
	Create(entrada *EntradaProducto) error
}

// SalidaProductoRepository define el puerto de persistencia para salidas
type SalidaProductoRepository interface {
	GetAll() ([]SalidaConProducto, error)
	GetByID(id int) (*SalidaConProducto, error)
	GetByProductoID(productoID int) ([]SalidaConProducto, error)
	GetByFecha(fecha string) ([]SalidaConProducto, error)
	GetByLugar(lugar string) ([]SalidaConProducto, error)
	Create(salida *SalidaProducto) error
}

// ControlDiarioRepository define el puerto de persistencia para control diario
type ControlDiarioRepository interface {
	GetAll() ([]ControlDiario, error)
	GetByFecha(fecha string) ([]ControlDiario, error)
	GetByFechaHoy() ([]ControlDiario, error)
	GetVerbena() ([]ControlDiario, error)
	Create(control *ControlDiario) error
	GenerarDesdeVentas(fecha string) (*ControlDiario, error)
}

// ResumenMensualRepository define el puerto de persistencia para resumen mensual
type ResumenMensualRepository interface {
	GetByMesAnio(mes, anio int) (*ResumenMensual, error)
	GetActual() (*ResumenMensual, error)
	GetByProductoID(productoID, mes, anio int) (*ResumenProducto, error)
	Generar(mes, anio int) (*ResumenMensual, error)
	Upsert(resumen *ResumenMensual) error
}

// UsuarioRepository define el puerto de persistencia para usuarios
type UsuarioRepository interface {
	GetByUsername(username string) (*Usuario, error)
	GetByID(id int) (*Usuario, error)
	Create(usuario *Usuario) error
	Update(usuario *Usuario) error
}

// ReportesRepository define el puerto de persistencia para reportes
type ReportesRepository interface {
	GetInventarioActual() ([]ReporteInventarioItem, error)
	GetMovimientos(inicio, fin string) ([]ReporteMovimiento, error)
	GetProductosMasVendidos(limite int) ([]ReporteProductoVendido, error)
	GetProductosMasIngresados(limite int) ([]ReporteProductoVendido, error)
	GetValoracionInventario() ([]ReporteValoracion, error)
}

// AlertasRepository define el puerto de persistencia para alertas
type AlertasRepository interface {
	GetStockBajo(limite int) ([]AlertaStockBajo, error)
}
