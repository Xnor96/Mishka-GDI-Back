package router

import (
	"github.com/Mishka-GDI-Back/infrastructure/http/handler"
	"github.com/Mishka-GDI-Back/infrastructure/http/middleware"
	"github.com/gin-gonic/gin"
)

type Router struct {
	categoriaHandler *handler.CategoriaHandler
	productoHandler  *handler.ProductoHandler
	entradaHandler   *handler.EntradaHandler
	salidaHandler    *handler.SalidaHandler
	controlHandler   *handler.ControlDiarioHandler
	resumenHandler   *handler.ResumenMensualHandler
	authHandler      *handler.AuthHandler
	reportesHandler  *handler.ReportesHandler
	alertasHandler   *handler.AlertasHandler
}

func NewRouter(
	categoriaHandler *handler.CategoriaHandler,
	productoHandler *handler.ProductoHandler,
	entradaHandler *handler.EntradaHandler,
	salidaHandler *handler.SalidaHandler,
	controlHandler *handler.ControlDiarioHandler,
	resumenHandler *handler.ResumenMensualHandler,
	authHandler *handler.AuthHandler,
	reportesHandler *handler.ReportesHandler,
	alertasHandler *handler.AlertasHandler,
) *Router {
	return &Router{
		categoriaHandler: categoriaHandler,
		productoHandler:  productoHandler,
		entradaHandler:   entradaHandler,
		salidaHandler:    salidaHandler,
		controlHandler:   controlHandler,
		resumenHandler:   resumenHandler,
		authHandler:      authHandler,
		reportesHandler:  reportesHandler,
		alertasHandler:   alertasHandler,
	}
}

func (r *Router) SetupRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(corsMiddleware())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Health check (público)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "message": "Mishka Inventory API is running"})
	})

	api := router.Group("/api")
	{
		// =============================================
		// Auth (público - sin JWT)
		// =============================================
		auth := api.Group("/auth")
		{
			auth.POST("/login", r.authHandler.Login)
			auth.POST("/refresh-token", r.authHandler.RefreshToken)
			auth.POST("/logout", r.authHandler.Logout)
		}

		// =============================================
		// Rutas protegidas con JWT
		// =============================================
		protected := api.Group("/")
		protected.Use(middleware.AuthRequired())
		{
			// Categorías
			categorias := protected.Group("categorias")
			{
				categorias.GET("", r.categoriaHandler.GetAll)
				categorias.GET("/:id", r.categoriaHandler.GetByID)
				categorias.POST("", r.categoriaHandler.Create)
				categorias.PUT("/:id", r.categoriaHandler.Update)
				categorias.DELETE("/:id", r.categoriaHandler.Delete)
			}

			// Productos
			productos := protected.Group("productos")
			{
				productos.GET("", r.productoHandler.GetAll)
				productos.GET("/stock-bajo", r.productoHandler.GetStockBajo)
				productos.GET("/buscar", r.productoHandler.Search)
				productos.GET("/:id", r.productoHandler.GetByID)
				productos.POST("", r.productoHandler.Create)
				productos.PUT("/:id", r.productoHandler.Update)
				productos.DELETE("/:id", r.productoHandler.Delete)
			}

			// Entradas
			entradas := protected.Group("entradas")
			{
				entradas.GET("", r.entradaHandler.GetAll)
				entradas.GET("/producto/:id", r.entradaHandler.GetByProductoID)
				entradas.GET("/fecha/:fecha", r.entradaHandler.GetByFecha)
				entradas.GET("/:id", r.entradaHandler.GetByID)
				entradas.POST("", r.entradaHandler.Create)
			}

			// Salidas
			salidas := protected.Group("salidas")
			{
				salidas.GET("", r.salidaHandler.GetAll)
				salidas.GET("/producto/:id", r.salidaHandler.GetByProductoID)
				salidas.GET("/fecha/:fecha", r.salidaHandler.GetByFecha)
				salidas.GET("/lugar/:lugar", r.salidaHandler.GetByLugar)
				salidas.GET("/:id", r.salidaHandler.GetByID)
				salidas.POST("", r.salidaHandler.Create)
			}

			// Control Diario
			control := protected.Group("control-diario")
			{
				control.GET("", r.controlHandler.GetAll)
				control.GET("/hoy", r.controlHandler.GetHoy)
				control.GET("/verbena", r.controlHandler.GetVerbena)
				control.GET("/fecha/:fecha", r.controlHandler.GetByFecha)
				control.POST("", r.controlHandler.Create)
				control.POST("/generar/:fecha", r.controlHandler.GenerarDesdeVentas)
			}

			// Resumen Mensual
			resumen := protected.Group("resumen-mensual")
			{
				resumen.GET("/actual", r.resumenHandler.GetActual)
				resumen.GET("/producto/:id", r.resumenHandler.GetByProductoID)
				resumen.GET("/:mes/:anio", r.resumenHandler.GetByMesAnio)
				resumen.POST("/generar", r.resumenHandler.Generar)
			}

			// Reportes
			reportes := protected.Group("reportes")
			{
				reportes.GET("/inventario-actual", r.reportesHandler.GetInventarioActual)
				reportes.GET("/movimientos/:inicio/:fin", r.reportesHandler.GetMovimientos)
				reportes.GET("/productos-mas-vendidos", r.reportesHandler.GetProductosMasVendidos)
				reportes.GET("/productos-mas-ingresados", r.reportesHandler.GetProductosMasIngresados)
				reportes.GET("/valoracion-inventario", r.reportesHandler.GetValoracionInventario)
			}

			// Alertas
			alertas := protected.Group("alertas")
			{
				alertas.GET("", r.alertasHandler.GetAlertasActivas)
				alertas.GET("/stock-bajo", r.alertasHandler.GetStockBajo)
				alertas.POST("/configuracion", r.alertasHandler.ConfigurarAlerta)
			}
		}
	}

	return router
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
