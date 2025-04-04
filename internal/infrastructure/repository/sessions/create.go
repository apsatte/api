package sessions_repository

import (
	"api/internal/domain"
	"api/internal/domain/customer"
	"context"

	"go.uber.org/zap"
)

func (r *repo) Create(c context.Context, s *customer.Session) error {
	sql := `INSERT INTO customers.sessions (access_token, customer_id, user_agent, ip, created_at)
					VALUES ($1, $2, $3, $4, $5);`

	_, err := r.db.Exec(c, sql, s.AccessToken, s.CustomerID, s.UserAgent, s.IP, s.CreatedAt)
	if err != nil {
		r.logger.Error("database error", zap.Error(err))
		return domain.ErrDatabase
	}

	return nil
}
