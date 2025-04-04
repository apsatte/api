package sessions_repository

import (
	"api/internal/domain"
	"api/internal/domain/customer"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func (r *repo) GetOneByAccessToken(c context.Context, accessToken string) (*customer.Session, error) {
	var s customer.Session

	sql := `SELECT access_token, customer_id, user_agent, ip, created_at
					FROM customers.sessions WHERE access_token = $1 AND removed_at IS NULL;`

	row := r.db.QueryRow(c, sql, accessToken)
	err := row.Scan(&s.AccessToken, &s.CustomerID, &s.UserAgent, &s.IP, &s.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		r.logger.Error("database error", zap.Error(err))
		return nil, domain.ErrDatabase
	}

	return &s, nil
}
