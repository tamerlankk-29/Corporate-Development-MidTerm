package main

import (
	"context"
	_ "database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	_ "gorm.io/gorm"
	"task-management-app/internal/config"
	"task-management-app/internal/db"
	"task-management-app/internal/handler"
	"task-management-app/internal/model"
	"task-management-app/internal/repository"
	"task-management-app/internal/repository/gormrepo"
	"task-management-app/internal/repository/rawrepo"
	"task-management-app/internal/service"
	"task-management-app/internal/transport"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed load config: %v", err)
	}

	sqlDB, err := db.NewPostgresDB(cfg.PostgresDSN())
	if err != nil {
		log.Fatalf("failed connect raw db: %v", err)
	}
	defer sqlDB.Close()

	gormDB, err := db.NewGormDB(cfg.PostgresDSN())
	if err != nil {
		log.Fatalf("failed connect gorm db: %v", err)
	}

	if err := gormDB.AutoMigrate(&model.Task{}); err != nil {
		log.Fatalf("gorm automigrate: %v", err)
	}

	var repo repository.TaskRepository
	switch cfg.RepoDriver {
	case "gorm":
		repo = gormrepo.NewGormTaskRepository(gormDB)
	default:
		repo = rawrepo.NewRawTaskRepository(sqlDB)
	}

	svc := service.NewTaskService(repo)
	taskHandler := handler.NewTaskHandler(svc)

	router := transport.NewRouter(func(r *mux.Router) {
		taskHandler.RegisterRoutes(r)
	})

	srv := &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: router,
	}

	go func() {
		log.Printf("server running on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server listen error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown failed: %v", err)
	}
	log.Println("server exited")
}
