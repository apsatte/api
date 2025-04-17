package main

import (
	account_usecase "api/internal/application/account"
	category_usecase "api/internal/application/category"
	dish_usecase "api/internal/application/dish"
	project_usecase "api/internal/application/project"
	categories_repository "api/internal/infrastructure/repository/categories"
	customers_repository "api/internal/infrastructure/repository/customers"
	dishes_repository "api/internal/infrastructure/repository/dishes"
	projects_repository "api/internal/infrastructure/repository/projects"
	sessions_repository "api/internal/infrastructure/repository/sessions"
	"api/internal/transport/http"
	account_handler "api/internal/transport/http/handlers/account"
	menu_handler "api/internal/transport/http/handlers/menu"
	projects_handler "api/internal/transport/http/handlers/project"
	"api/pkg/configuration"
	"api/pkg/logger"
	"api/pkg/postgresql"
	"api/pkg/storage"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {
	config, err := configuration.Load()
	if err != nil {
		log.Fatal(err)
	}

	logger, err := logger.New(&config.Logger)
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgresql.New(&config.DB)
	if err != nil {
		logger.Fatal("failed to connect postgresql", zap.Error(err))
	}

	storage, err := storage.New(&config.S3)
	if err != nil {
		logger.Fatal("failed to connect S3 Storage", zap.Error(err))
	}

	// repositories
	categoryRepo := categories_repository.New(logger, db)
	dishRepo := dishes_repository.New(logger, db)
	customersRepo := customers_repository.New(logger, db)
	sessionsRepo := sessions_repository.New(logger, db)
	projectsRepo := projects_repository.New(logger, db)

	// usecases
	accountUseCase := account_usecase.New(customersRepo, sessionsRepo)
	projectsUseCase := project_usecase.New(projectsRepo, customersRepo, storage, config.Server.CdnBaseUrl)
	categoryUseCase := category_usecase.New(categoryRepo, customersRepo, projectsRepo)
	dishUseCase := dish_usecase.New(categoryRepo, customersRepo, projectsRepo, dishRepo, storage)

	// handlers
	accountHandler := account_handler.New(accountUseCase)
	projectHandler := projects_handler.New(projectsUseCase)
	menuHandler := menu_handler.New(categoryUseCase, dishUseCase)

	// http server
	httpServer := http.New(sessionsRepo)
	httpServer.SetupRouter(
		accountHandler,
		projectHandler,
		menuHandler,
	)

	go func() {
		if err := httpServer.Run(config.Server.HttpSocket); err != nil {
			logger.Fatal("failed to start http server")
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		logger.Error("server forced to shutdown", zap.Error(err))
	}
}
