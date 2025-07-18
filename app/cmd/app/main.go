package main

import (
	_ "awesomeProject666/app/cmd/app/docs"
	"awesomeProject666/app/internal/auth"
	"awesomeProject666/app/internal/config"
	"awesomeProject666/app/internal/database"
	"awesomeProject666/app/internal/delivery/http/v1"
	"awesomeProject666/app/internal/repository"
	"awesomeProject666/app/internal/service"
	"awesomeProject666/app/pkg/logging"
	"context"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title           E-commerce API
// @version         1.0
// @description     Backend сервис для управления пользователями и товарами.
// @termsOfService  http://swagger.io/terms/
// @host      localhost:8080
// @BasePath  /api/v1
// @schemes   http
func main() {
	cfg, err := config.GetConfig(".env")
	if err != nil {
		panic(err)
	}
	logg := logging.InitLogger(cfg.LogLevel)
	logg.Info("Logger initialized successfully")

	pgPool, err := database.NewPostgresPool(cfg.Databases.PostgresDSN)
	if err != nil {
		logg.Fatalf("postgres connection failed: %v", err)
	}
	defer pgPool.Close()
	logg.Info("postgres connection established")

	jwtManager := auth.NewJWTManager(cfg.JWTSecret, 24*time.Hour)
	repos := repository.NewRepository(pgPool)
	services := service.NewServices(repos, cfg)
	h := v1.NewHandler(logg.Logger, services, jwtManager)

	r := h.Init()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	srv := &http.Server{
		Addr:    cfg.AppPort,
		Handler: r,
	}

	go func() {
		logg.Infof("Starting server on port %s", cfg.AppPort)
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logg.Fatalf("Could not start the server: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	logg.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = srv.Shutdown(ctx); err != nil {
		logg.Fatalf("Could not gracefully shutdown the server: %v\n", err)
	}

	logg.Info("Server gracefully stopped")
}
