package router

import (
	"github.com/Mishka-GDI-Back/handler"
	"github.com/gin-gonic/gin"
)

type Router struct {
	categoriaHandler *handler.CategoriaHandler
	productoHandler  *handler.ProductoHandler
	entradaHandler   *handler.EntradaHandler
	salidaHandler    *handler.SalidaHandler
}

func NewRouter(categoriaHandler *handler.CategoriaHandler, productoHandler *handler.ProductoHandler, entradaHandler *handler.EntradaHandler, salidaHandler *handler.SalidaHandler) *Router {
	return &Router{
		categoriaHandler: categoriaHandler,
		productoHandler:  productoHandler,
		entradaHandler:   entradaHandler,
		salidaHandler:    salidaHandler,
	}
}

func (r *Router) SetupRoutes() *gin.Engine {
	router := gin.Default()

	// Middleware para CORS
	router.Use(corsMiddleware())

	// Middleware para logging
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Mishka Inventory API is running",
		})
	})

	// API routes
	api := router.Group("/api")
	{
		// Categorías
		categorias := api.Group("/categorias")
		{
			categorias.GET("", r.categoriaHandler.GetAll)
			categorias.GET("/:id", r.categoriaHandler.GetByID)
			categorias.POST("", r.categoriaHandler.Create)
			categorias.PUT("/:id", r.categoriaHandler.Update)
			categorias.DELETE("/:id", r.categoriaHandler.Delete)
		}

		// Productos
		productos := api.Group("/productos")
		{
			productos.GET("", r.productoHandler.GetAll)
			productos.GET("/:id", r.productoHandler.GetByID)
			productos.POST("", r.productoHandler.Create)
			productos.PUT("/:id", r.productoHandler.Update)
			productos.DELETE("/:id", r.productoHandler.Delete)
			productos.GET("/stock-bajo", r.productoHandler.GetStockBajo)
			productos.GET("/buscar", r.productoHandler.Search)
		}

		// Entradas
		entradas := api.Group("/entradas")
		{
			entradas.GET("", r.entradaHandler.GetAll)
			entradas.GET("/:id", r.entradaHandler.GetByID)
			entradas.POST("", r.entradaHandler.Create)
			entradas.GET("/producto/:id", r.entradaHandler.GetByProductoID)
			entradas.GET("/fecha/:fecha", r.entradaHandler.GetByFecha)
		}

		// Salidas
		salidas := api.Group("/salidas")
		{
			salidas.GET("", r.salidaHandler.GetAll)
			salidas.GET("/:id", r.salidaHandler.GetByID)
			salidas.POST("", r.salidaHandler.Create)
			salidas.GET("/producto/:id", r.salidaHandler.GetByProductoID)
			salidas.GET("/fecha/:fecha", r.salidaHandler.GetByFecha)
		}

		// TODO: Rutas para reportes, control diario, etc.
		// Estas se implementarán en las siguientes iteraciones
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
