package persistence

import (
	"context"
	"fmt"
	"time"

	"github.com/Mishka-GDI-Back/domain"
	"github.com/Mishka-GDI-Back/infrastructure/database"
)

type controlDiarioRepository struct {
	db *database.Database
}

func NewControlDiarioRepository(db *database.Database) domain.ControlDiarioRepository {
	return &controlDiarioRepository{db: db}
}

const controlSelect = `SELECT id_control, fecha, descripcion, monto_entrada, monto_salida, observaciones, es_verbena, usuario_registro, fecha_creacion, fecha_actualizacion FROM control_diario`

func scanControl(row interface{ Scan(dest ...any) error }) (domain.ControlDiario, error) {
	var c domain.ControlDiario
	err := row.Scan(&c.ID, &c.Fecha, &c.Descripcion, &c.MontoEntrada, &c.MontoSalida, &c.Observaciones, &c.EsVerbena, &c.UsuarioRegistro, &c.FechaCreacion, &c.FechaActualizacion)
	return c, err
}

func (r *controlDiarioRepository) GetAll() ([]domain.ControlDiario, error) {
	rows, err := r.db.Pool.Query(context.Background(), controlSelect+" ORDER BY fecha DESC, fecha_creacion DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var controles []domain.ControlDiario
	for rows.Next() {
		c, err := scanControl(rows)
		if err != nil {
			return nil, err
		}
		controles = append(controles, c)
	}
	return controles, nil
}

func (r *controlDiarioRepository) GetByFecha(fecha string) ([]domain.ControlDiario, error) {
	rows, err := r.db.Pool.Query(context.Background(), controlSelect+" WHERE fecha = $1 ORDER BY fecha_creacion DESC", fecha)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var controles []domain.ControlDiario
	for rows.Next() {
		c, err := scanControl(rows)
		if err != nil {
			return nil, err
		}
		controles = append(controles, c)
	}
	return controles, nil
}

func (r *controlDiarioRepository) GetByFechaHoy() ([]domain.ControlDiario, error) {
	return r.GetByFecha(time.Now().Format("2006-01-02"))
}

func (r *controlDiarioRepository) GetVerbena() ([]domain.ControlDiario, error) {
	rows, err := r.db.Pool.Query(context.Background(), controlSelect+" WHERE es_verbena = TRUE ORDER BY fecha DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var controles []domain.ControlDiario
	for rows.Next() {
		c, err := scanControl(rows)
		if err != nil {
			return nil, err
		}
		controles = append(controles, c)
	}
	return controles, nil
}

func (r *controlDiarioRepository) Create(control *domain.ControlDiario) error {
	query := `INSERT INTO control_diario (fecha, descripcion, monto_entrada, monto_salida, observaciones, es_verbena, usuario_registro) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id_control, fecha_creacion, fecha_actualizacion`
	return r.db.Pool.QueryRow(context.Background(), query, control.Fecha, control.Descripcion, control.MontoEntrada, control.MontoSalida, control.Observaciones, control.EsVerbena, control.UsuarioRegistro).Scan(&control.ID, &control.FechaCreacion, &control.FechaActualizacion)
}

func (r *controlDiarioRepository) GenerarDesdeVentas(fecha string) (*domain.ControlDiario, error) {
	var totalVentas float64
	var cantidadVentas int
	err := r.db.Pool.QueryRow(context.Background(), `SELECT COALESCE(SUM(total), 0), COUNT(*) FROM salidas_productos WHERE fecha_salida = $1`, fecha).Scan(&totalVentas, &cantidadVentas)
	if err != nil {
		return nil, err
	}
	fechaTime, err := time.Parse("2006-01-02", fecha)
	if err != nil {
		return nil, err
	}
	control := &domain.ControlDiario{
		Fecha:           fechaTime,
		Descripcion:     fmt.Sprintf("CORTE DE VENTAS - %d venta(s) registrada(s)", cantidadVentas),
		MontoEntrada:    totalVentas,
		MontoSalida:     0,
		Observaciones:   fmt.Sprintf("Generado automaticamente desde %d salida(s)", cantidadVentas),
		EsVerbena:       false,
		UsuarioRegistro: "sistema",
	}
	if err := r.Create(control); err != nil {
		return nil, err
	}
	return control, nil
}
