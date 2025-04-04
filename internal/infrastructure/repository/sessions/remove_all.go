package sessions_repository

import (
	"api/internal/domain"
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (r *repo) RemoveAll(c context.Context, customerID uuid.UUID) error {
	sql := "UPDATE customers.sessions SET removed_at = CURRENT_TIMESTAMP WHERE customer_id = $1;"

	_, err := r.db.Exec(c, sql, customerID)
	if err != nil {
		r.logger.Error("database error", zap.Error(err))
		return domain.ErrDatabase
	}

	return nil
}
