package modules_repository

import (
	"api/internal/domain/module"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type repo struct {
	logger *zap.Logger
	db     *pgxpool.Pool
}

func New(logger *zap.Logger, db *pgxpool.Pool) module.ModulesRepository {
	return &repo{
		logger: logger,
		db:     db,
	}
}
