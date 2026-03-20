package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Mishka-GDI-Back/application"
	"github.com/Mishka-GDI-Back/infrastructure/config"
	"github.com/Mishka-GDI-Back/infrastructure/database"
	"github.com/Mishka-GDI-Back/infrastructure/http/handler"
	"github.com/Mishka-GDI-Back/infrastructure/http/router"
	"github.com/Mishka-GDI-Back/infrastructure/persistence"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.NewConfig()
	gin.SetMode(cfg.GinMode)

	log.Printf("🚀 Iniciando Mishka Inventory API en puerto %s", cfg.Port)

	db, err := database.NewPostgresConnection(cfg.PostgresURI)
	if err != nil {
		log.Fatalf("❌ Error al conectar con la base de datos: %v", err)
	}
	defer db.Close()

	// ── Repositorios (capa de infraestructura / persistencia) ──────────────
	categoriaRepo := persistence.NewCategoriaRepository(db)
	productoRepo  := persistence.NewProductoRepository(db)
	entradaRepo   := persistence.NewEntradaProductoRepository(db)
	salidaRepo    := persistence.NewSalidaProductoRepository(db)
	controlRepo   := persistence.NewControlDiarioRepository(db)
	resumenRepo   := persistence.NewResumenMensualRepository(db)
	usuarioRepo   := persistence.NewUsuarioRepository(db)
	reportesRepo  := persistence.NewReportesRepository(db)
	alertasRepo   := persistence.NewAlertasRepository(db)

	// ── Servicios (capa de aplicación) ─────────────────────────────────────
	categoriaService := application.NewCategoriaService(categoriaRepo)
	productoService  := application.NewProductoService(productoRepo, categoriaRepo)
	entradaService   := application.NewEntradaProductoService(entradaRepo, productoRepo)
	salidaService    := application.NewSalidaProductoService(salidaRepo, productoRepo)
	controlService   := application.NewControlDiarioService(controlRepo)
	resumenService   := application.NewResumenMensualService(resumenRepo)
	authService      := application.NewAuthService(usuarioRepo)
	reportesService  := application.NewReportesService(reportesRepo)
	alertasService   := application.NewAlertasService(alertasRepo)

	// ── Handlers (capa de infraestructura / HTTP) ───────────────────────────
	categoriaHandler := handler.NewCategoriaHandler(categoriaService)
	productoHandler  := handler.NewProductoHandler(productoService)
	entradaHandler   := handler.NewEntradaHandler(entradaService)
	salidaHandler    := handler.NewSalidaHandler(salidaService)
	controlHandler   := handler.NewControlDiarioHandler(controlService)
	resumenHandler   := handler.NewResumenMensualHandler(resumenService)
	authHandler      := handler.NewAuthHandler(authService)
	reportesHandler  := handler.NewReportesHandler(reportesService)
	alertasHandler   := handler.NewAlertasHandler(alertasService)

	// ── Router ──────────────────────────────────────────────────────────────
	appRouter := router.NewRouter(
		categoriaHandler, productoHandler, entradaHandler, salidaHandler,
		controlHandler, resumenHandler, authHandler, reportesHandler, alertasHandler,
	)
	ginRouter := appRouter.SetupRoutes()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("✅ Servidor iniciado en http://localhost:%s", cfg.Port)
		log.Printf("📚 Health check: http://localhost:%s/health", cfg.Port)
		log.Printf("🔐 Login: POST http://localhost:%s/api/auth/login", cfg.Port)
		log.Printf("🔧 API: http://localhost:%s/api", cfg.Port)

		if err := ginRouter.Run(":" + cfg.Port); err != nil {
			log.Fatalf("❌ Error al iniciar el servidor: %v", err)
		}
	}()

	<-quit
	log.Println("🛑 Cerrando servidor...")
	log.Println("👋 Servidor cerrado exitosamente")
}
