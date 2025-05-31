# Mishka-GDI-Back

Sistema de gestión de inventario backend para la aplicación Mishka. API REST desarrollada en Go con Gin-Gonic y PostgreSQL.

## 🚀 Características

- ✅ **CRUD completo** para Categorías, Productos, Entradas y Salidas
- ✅ **Búsqueda avanzada** de productos por código o nombre
- ✅ **Control de stock** automático con entradas y salidas
- ✅ **Alertas de stock bajo** configurables
- ✅ **Arquitectura limpia** con separación de responsabilidades
- ✅ **Validaciones robustas** en todas las capas
- ✅ **Manejo de errores** centralizado
- ✅ **CORS** configurado para frontend
- ✅ **Pool de conexiones** PostgreSQL optimizado

## 🛠️ Tecnologías

- **Go 1.23.3**
- **Gin-Gonic** (Framework web)
- **PostgreSQL** (Base de datos)
- **pgx/v5** (Driver PostgreSQL con pool de conexiones)
- **Arquitectura Clean Architecture**

## 📁 Estructura del Proyecto

```
Mishka-GDI-Back/
├── cmd/
│   └── main.go                 # Punto de entrada de la aplicación
├── config/
│   └── config.go              # Configuración de la aplicación
├── db/
│   └── postgres.go            # Conexión a PostgreSQL
├── domain/
│   ├── categoria.go           # Entidad Categoria
│   ├── producto.go            # Entidad Producto
│   ├── entrada_producto.go    # Entidad EntradaProducto
│   ├── salida_producto.go     # Entidad SalidaProducto
│   └── repository.go          # Interfaces de repositorios
├── repository/
│   ├── categoria_repository.go
│   ├── producto_repository.go
│   ├── entrada_repository.go
│   └── salida_repository.go
├── service/
│   ├── categoria_service.go
│   ├── producto_service.go
│   ├── entrada_service.go
│   └── salida_service.go
├── handler/
│   ├── categoria_handler.go
│   ├── producto_handler.go
│   ├── entrada_handler.go
│   └── salida_handler.go
├── dto/
│   ├── request.go             # DTOs de entrada
│   └── response.go            # DTOs de respuesta
├── router/
│   └── router.go              # Configuración de rutas
├── .vscode/
│   └── launch.json            # Configuración de debug para VS Code
├── ddl.sql                    # Script de creación de BD
├── dml.sql                    # Script de datos de prueba
├── endpoints.txt              # Documentación de endpoints
├── go.mod                     # Dependencias de Go
└── README.md
```

## 🔧 Configuración

### Prerrequisitos

1. **Go 1.23.3+**
2. **PostgreSQL 12+**
3. **Git**

### Instalación

1. **Clonar el repositorio**
```bash
git clone https://github.com/tu-usuario/Mishka-GDI-Back.git
cd Mishka-GDI-Back
```

2. **Instalar dependencias**
```bash
go mod tidy
```

3. **Configurar la base de datos**
```bash
# Crear la base de datos
psql -U postgres -c "CREATE DATABASE mishka;"

# Ejecutar el script DDL
psql -U postgres -d mishka -f ddl.sql

# Cargar datos de prueba (opcional)
psql -U postgres -d mishka -f dml.sql
```

4. **Configurar variables de entorno en VS Code**

El archivo `.vscode/launch.json` ya está configurado con las variables necesarias:

```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Mishka Backend",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}/cmd/main.go",
            "env": {
                "POSTGRES_URI": "postgres://usuario:password@localhost:5432/mishka?sslmode=disable",
                "PORT": "8080",
                "GIN_MODE": "debug"
            },
            "console": "integratedTerminal"
        }
    ]
}
```

**⚠️ Importante:** Actualiza la `POSTGRES_URI` con tus credenciales de PostgreSQL.

## 🚀 Ejecución

### Desde VS Code (Recomendado)
1. Abre el proyecto en VS Code
2. Ve a Run and Debug (Ctrl+Shift+D)
3. Selecciona "Launch Mishka Backend"
4. Presiona F5 o click en "Start Debugging"

### Desde terminal
```bash
# Configurar variables de entorno
export POSTGRES_URI="postgres://usuario:password@localhost:5432/mishka?sslmode=disable"
export PORT="8080"
export GIN_MODE="debug"

# Ejecutar la aplicación
go run cmd/main.go
```

## 📚 API Endpoints

### Health Check
- `GET /health` - Verificar estado del servidor

### Categorías
- `GET /api/categorias` - Listar todas las categorías
- `GET /api/categorias/{id}` - Obtener categoría por ID
- `POST /api/categorias` - Crear nueva categoría
- `PUT /api/categorias/{id}` - Actualizar categoría
- `DELETE /api/categorias/{id}` - Eliminar categoría

### Productos
- `GET /api/productos` - Listar todos los productos
- `GET /api/productos/{id}` - Obtener producto por ID
- `POST /api/productos` - Crear nuevo producto
- `PUT /api/productos/{id}` - Actualizar producto
- `DELETE /api/productos/{id}` - Eliminar producto
- `GET /api/productos/stock-bajo?limite=5` - Productos con stock bajo
- `GET /api/productos/buscar?q=termino` - Buscar productos

### Entradas
- `GET /api/entradas` - Listar todas las entradas
- `GET /api/entradas/{id}` - Obtener entrada por ID
- `POST /api/entradas` - Registrar nueva entrada
- `GET /api/entradas/producto/{id}` - Entradas por producto
- `GET /api/entradas/fecha/{fecha}` - Entradas por fecha (YYYY-MM-DD)

### Salidas
- `GET /api/salidas` - Listar todas las salidas
- `GET /api/salidas/{id}` - Obtener salida por ID
- `POST /api/salidas` - Registrar nueva salida
- `GET /api/salidas/producto/{id}` - Salidas por producto
- `GET /api/salidas/fecha/{fecha}` - Salidas por fecha (YYYY-MM-DD)

## 📝 Ejemplos de Uso

### Crear una categoría
```bash
curl -X POST http://localhost:8080/api/categorias \
  -H "Content-Type: application/json" \
  -d '{"nombre": "COLLAR"}'
```

### Crear un producto
```bash
curl -X POST http://localhost:8080/api/productos \
  -H "Content-Type: application/json" \
  -d '{
    "codigo": "C001",
    "nombre": "Collar de Perlas",
    "id_categoria": 1,
    "unidad_medida": "UNIDAD",
    "precio_unitario": 250.00,
    "stock_actual": 10,
    "stock_inicial": 10
  }'
```

### Registrar entrada de producto
```bash
curl -X POST http://localhost:8080/api/entradas \
  -H "Content-Type: application/json" \
  -d '{
    "id_producto": 1,
    "fecha_entrada": "2025-01-15",
    "cantidad": 5,
    "precio_unitario": 250.00,
    "observaciones": "Reposición de inventario",
    "usuario_registro": "admin"
  }'
```

### Registrar salida de producto
```bash
curl -X POST http://localhost:8080/api/salidas \
  -H "Content-Type: application/json" \
  -d '{
    "id_producto": 1,
    "fecha_salida": "2025-01-16",
    "cantidad": 2,
    "observaciones": "Venta en tienda",
    "usuario_registro": "vendedor1"
  }'
```

## 🔄 Control de Stock

El sistema maneja automáticamente el stock de productos:

- **Entradas**: Incrementan el `stock_actual` del producto
- **Salidas**: Decrementan el `stock_actual` del producto
- **Validaciones**: 
  - No se permiten salidas con stock insuficiente
  - Los stocks no pueden ser negativos
  - Las cantidades deben ser positivas

## 🚧 Próximas Características

- [ ] Autenticación JWT
- [ ] Reportes avanzados
- [ ] Control diario automatizado
- [ ] Resúmenes mensuales
- [ ] Alertas configurables
- [ ] Exportación de datos
- [ ] Dashboard de métricas
- [ ] Backup automático

## 🤝 Contribución

1. Fork el proyecto
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

## 📄 Licencia

Este proyecto está bajo la Licencia MIT. Ver el archivo `LICENSE` para más detalles.

## 📞 Contacto

Para preguntas o soporte, contacta al equipo de desarrollo.

---
**Mishka Inventory Management System** - Desarrollado con ❤️ en Go