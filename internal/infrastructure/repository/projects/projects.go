package projects_repository

import (
	"api/internal/domain/project"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type repo struct {
	logger *zap.Logger
	db     *pgxpool.Pool
}

func New(logger *zap.Logger, db *pgxpool.Pool) project.ProjectsRepository {
	return &repo{
		logger: logger,
		db:     db,
	}
}
