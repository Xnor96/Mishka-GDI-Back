package persistence

import (
	"context"
	"time"

	"github.com/Mishka-GDI-Back/domain"
	"github.com/Mishka-GDI-Back/infrastructure/database"
	"github.com/jackc/pgx/v5"
)

type salidaProductoRepository struct {
	db *database.Database
}

func NewSalidaProductoRepository(db *database.Database) domain.SalidaProductoRepository {
	return &salidaProductoRepository{db: db}
}

const salidaSelectJoin = `
	SELECT sp.id_salida, sp.id_producto, sp.fecha_salida, sp.cantidad,
	       sp.precio_venta, sp.descuento, sp.total, sp.lugar_venta, sp.tipo_pago,
	       sp.observaciones, sp.usuario_registro, sp.fecha_creacion, sp.fecha_actualizacion,
	       p.nombre, p.codigo, COALESCE(c.nombre, '') AS nombre_categoria
	FROM salidas_productos sp
	JOIN productos p ON sp.id_producto = p.id_producto
	LEFT JOIN categorias c ON p.id_categoria = c.id_categoria`

func scanSalidaConProducto(rows pgx.Rows) (domain.SalidaConProducto, error) {
	var s domain.SalidaConProducto
	err := rows.Scan(
		&s.ID, &s.IDProducto, &s.FechaSalida, &s.Cantidad,
		&s.PrecioVenta, &s.Descuento, &s.Total, &s.LugarVenta, &s.TipoPago,
		&s.Observaciones, &s.UsuarioRegistro, &s.FechaCreacion, &s.FechaActualizacion,
		&s.NombreProducto, &s.CodigoProducto, &s.NombreCategoria,
	)
	return s, err
}

func (r *salidaProductoRepository) GetAll() ([]domain.SalidaConProducto, error) {
	rows, err := r.db.Pool.Query(context.Background(), salidaSelectJoin+" ORDER BY sp.fecha_salida DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var salidas []domain.SalidaConProducto
	for rows.Next() {
		s, err := scanSalidaConProducto(rows)
		if err != nil {
			return nil, err
		}
		salidas = append(salidas, s)
	}
	return salidas, nil
}

func (r *salidaProductoRepository) GetByID(id int) (*domain.SalidaConProducto, error) {
	rows, err := r.db.Pool.Query(context.Background(), salidaSelectJoin+" WHERE sp.id_salida = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, &domain.ErrNotFound{Entity: "salida", ID: id}
	}
	s, err := scanSalidaConProducto(rows)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *salidaProductoRepository) GetByProductoID(productoID int) ([]domain.SalidaConProducto, error) {
	rows, err := r.db.Pool.Query(context.Background(), salidaSelectJoin+" WHERE sp.id_producto = $1 ORDER BY sp.fecha_salida DESC", productoID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var salidas []domain.SalidaConProducto
	for rows.Next() {
		s, err := scanSalidaConProducto(rows)
		if err != nil {
			return nil, err
		}
		salidas = append(salidas, s)
	}
	return salidas, nil
}

func (r *salidaProductoRepository) GetByFecha(fecha string) ([]domain.SalidaConProducto, error) {
	rows, err := r.db.Pool.Query(context.Background(), salidaSelectJoin+" WHERE sp.fecha_salida = $1 ORDER BY sp.fecha_creacion DESC", fecha)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var salidas []domain.SalidaConProducto
	for rows.Next() {
		s, err := scanSalidaConProducto(rows)
		if err != nil {
			return nil, err
		}
		salidas = append(salidas, s)
	}
	return salidas, nil
}

func (r *salidaProductoRepository) GetByLugar(lugar string) ([]domain.SalidaConProducto, error) {
	rows, err := r.db.Pool.Query(context.Background(), salidaSelectJoin+" WHERE UPPER(sp.lugar_venta) = UPPER($1) ORDER BY sp.fecha_salida DESC", lugar)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var salidas []domain.SalidaConProducto
	for rows.Next() {
		s, err := scanSalidaConProducto(rows)
		if err != nil {
			return nil, err
		}
		salidas = append(salidas, s)
	}
	return salidas, nil
}

func (r *salidaProductoRepository) Create(salida *domain.SalidaProducto) error {
	fechaSalida, err := time.Parse("2006-01-02", salida.FechaSalida.Format("2006-01-02"))
	if err != nil {
		return err
	}
	query := `INSERT INTO salidas_productos (id_producto, fecha_salida, cantidad, precio_venta, descuento, total, lugar_venta, tipo_pago, observaciones, usuario_registro) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id_salida, fecha_creacion, fecha_actualizacion`
	err = r.db.Pool.QueryRow(context.Background(), query, salida.IDProducto, fechaSalida, salida.Cantidad, salida.PrecioVenta, salida.Descuento, salida.Total, salida.LugarVenta, salida.TipoPago, salida.Observaciones, salida.UsuarioRegistro).Scan(&salida.ID, &salida.FechaCreacion, &salida.FechaActualizacion)
	if err != nil {
		return err
	}
	salida.FechaSalida = fechaSalida
	return nil
}
