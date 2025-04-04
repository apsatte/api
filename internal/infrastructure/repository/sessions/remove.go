package sessions_repository

import (
	"api/internal/domain"
	"api/internal/domain/customer"
	"context"

	"go.uber.org/zap"
)

func (r *repo) Remove(c context.Context, session *customer.Session) error {
	sql := "UPDATE customers.sessions SET removed_at = CURRENT_TIMESTAMP WHERE access_token = $1;"

	_, err := r.db.Exec(c, sql, session.AccessToken)
	if err != nil {
		r.logger.Error("database error", zap.Error(err))
		return domain.ErrDatabase
	}

	return nil
}
