package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Mishka-GDI-Back/config"
	"github.com/Mishka-GDI-Back/db"
	"github.com/Mishka-GDI-Back/handler"
	"github.com/Mishka-GDI-Back/repository"
	"github.com/Mishka-GDI-Back/router"
	"github.com/Mishka-GDI-Back/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// Cargar configuración
	cfg := config.NewConfig()

	// Configurar modo de Gin
	gin.SetMode(cfg.GinMode)

	log.Printf("🚀 Iniciando Mishka Inventory API en puerto %s", cfg.Port)
	log.Printf("🔗 Conectando a base de datos en: %s", cfg.PostgresURI)

	// Conectar a la base de datos
	database, err := db.NewPostgresConnection(cfg.PostgresURI)
	if err != nil {
		log.Fatalf("❌ Error al conectar con la base de datos: %v", err)
	}
	defer database.Close()

	// Inicializar repositorios
	categoriaRepo := repository.NewCategoriaRepository(database)
	productoRepo := repository.NewProductoRepository(database)
	entradaRepo := repository.NewEntradaProductoRepository(database)
	salidaRepo := repository.NewSalidaProductoRepository(database)

	// Inicializar servicios
	categoriaService := service.NewCategoriaService(categoriaRepo)
	productoService := service.NewProductoService(productoRepo, categoriaRepo)
	entradaService := service.NewEntradaProductoService(entradaRepo, productoRepo)
	salidaService := service.NewSalidaProductoService(salidaRepo, productoRepo)

	// Inicializar handlers
	categoriaHandler := handler.NewCategoriaHandler(categoriaService)
	productoHandler := handler.NewProductoHandler(productoService)
	entradaHandler := handler.NewEntradaHandler(entradaService)
	salidaHandler := handler.NewSalidaHandler(salidaService)

	// Configurar router
	appRouter := router.NewRouter(categoriaHandler, productoHandler, entradaHandler, salidaHandler)
	ginRouter := appRouter.SetupRoutes()

	// Canal para capturar señales del sistema
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Ejecutar servidor en una goroutine
	go func() {
		log.Printf("✅ Servidor iniciado en http://localhost:%s", cfg.Port)
		log.Printf("📚 Health check disponible en http://localhost:%s/health", cfg.Port)
		log.Printf("🔧 API disponible en http://localhost:%s/api", cfg.Port)

		if err := ginRouter.Run(":" + cfg.Port); err != nil {
			log.Fatalf("❌ Error al iniciar el servidor: %v", err)
		}
	}()

	// Esperar señal de terminación
	<-quit
	log.Println("🛑 Cerrando servidor...")
	log.Println("👋 Servidor cerrado exitosamente")
}
