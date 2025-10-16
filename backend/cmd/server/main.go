package main

import (
	"context"
	_ "database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"task-management-app/internal/repository/gormrepo"
	"time"

	"task-management-app/internal/config"
	"task-management-app/internal/db"
	"task-management-app/internal/handler"
	"task-management-app/internal/model"
<<<<<<< HEAD
	"task-management-app/internal/repository"
	"task-management-app/internal/repository/gormrepo"
	"task-management-app/internal/repository/rawrepo"
=======

	"github.com/gorilla/mux"
	_ "gorm.io/gorm"

>>>>>>> backend
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
	log.Println("Successfully connected to PostgreSQL (raw sql)!")
	defer sqlDB.Close()

	gormDB, err := db.NewGormDB(cfg.PostgresDSN())
	if err != nil {
		log.Fatalf("failed connect gorm db: %v", err)
	}
	log.Println("Successfully connected to PostgreSQL (gorm)!")

	if err := gormDB.AutoMigrate(&model.Task{}); err != nil {
		log.Fatalf("gorm automigrate: %v", err)
	}

<<<<<<< HEAD
	var repo repository.TaskRepository
	switch cfg.RepoDriver {
	case "gorm":
		repo = gormrepo.NewGormTaskRepository(gormDB)
	default:
		repo = rawrepo.NewRawTaskRepository(sqlDB)
	}
=======
	//rawRepo := rawrepo.NewRawTaskRepository(sqlDB)
	gormRepo := gormrepo.NewGormTaskRepository(gormDB)

	repo := gormRepo
>>>>>>> backend

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
