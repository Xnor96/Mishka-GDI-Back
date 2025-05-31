package domain

type CategoriaRepository interface {
	GetAll() ([]Categoria, error)
	GetByID(id int) (*Categoria, error)
	Create(categoria *Categoria) error
	Update(categoria *Categoria) error
	Delete(id int) error
}

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

type EntradaProductoRepository interface {
	GetAll() ([]EntradaProducto, error)
	GetByID(id int) (*EntradaProducto, error)
	GetByProductoID(productoID int) ([]EntradaProducto, error)
	GetByFecha(fecha string) ([]EntradaProducto, error)
	Create(entrada *EntradaProducto) error
}

type SalidaProductoRepository interface {
	GetAll() ([]SalidaProducto, error)
	GetByID(id int) (*SalidaProducto, error)
	GetByProductoID(productoID int) ([]SalidaProducto, error)
	GetByFecha(fecha string) ([]SalidaProducto, error)
	Create(salida *SalidaProducto) error
}
