package customers_repository

import (
	"api/internal/domain"
	"api/internal/domain/customer"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func (r *repo) GetOneByEmail(c context.Context, email string) (*customer.Customer, error) {
	var output customer.Customer
	output.Subscription = &customer.Subscription{}

	sql := `SELECT c.id, c.email, c.password, c.name, c.created_at, s.plan_option_id, s.expires_at, s.created_at as sub_created_at
					FROM customers.customers as c
					JOIN customers.subscriptions as s ON s.customer_id = c.id
					WHERE c.email = $1 AND removed_at IS NULL LIMIT 1;`

	row := r.db.QueryRow(c, sql, email)
	err := row.Scan(
		&output.ID, &output.Email, &output.Password,
		&output.Name, &output.CreatedAt,
		&output.Subscription.PlanOptionID, &output.Subscription.ExpiresAt,
		&output.Subscription.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		r.logger.Error("database error", zap.Error(err))
		return nil, domain.ErrDatabase
	}

	return &output, nil
}
