package plans_repository

import (
	"api/internal/domain/plan"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type repo struct {
	logger *zap.Logger
	db     *pgxpool.Pool
}

func New(logger *zap.Logger, db *pgxpool.Pool) plan.PlansRepository {
	return &repo{
		logger: logger,
		db:     db,
	}
}
