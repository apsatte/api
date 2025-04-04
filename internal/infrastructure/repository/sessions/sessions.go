package sessions_repository

import (
	"api/internal/domain/customer"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type repo struct {
	logger *zap.Logger
	db     *pgxpool.Pool
}

func New(logger *zap.Logger, db *pgxpool.Pool) customer.SessionsRepository {
	return &repo{
		logger: logger,
		db:     db,
	}
}
