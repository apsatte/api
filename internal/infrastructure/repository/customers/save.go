package customers_repository

import (
	"api/internal/domain"
	"api/internal/domain/customer"
	"context"

	"go.uber.org/zap"
)

func (r *repo) Save(c context.Context, customer *customer.Customer) error {
	tx, err := r.db.Begin(c)
	if err != nil {
		r.logger.Error("database error", zap.Error(err))
		return domain.ErrDatabase
	}
	defer tx.Rollback(c)

	sql := `INSERT INTO customers.customers (id, email, password, name, created_at)
					VALUES ($1, $2, $3, $4, $5)
					ON CONFLICT (id) DO UPDATE
					SET email = $2, password = $3, name = $4;`

	_, err = tx.Exec(
		c, sql,
		customer.ID, customer.Email, customer.Password,
		customer.Name, customer.CreatedAt)
	if err != nil {
		r.logger.Error("database error", zap.Error(err))
		return domain.ErrDatabase
	}

	sql = `INSERT INTO customers.subscriptions (customer_id, plan_option_id, expires_at, created_at)
				 VALUES ($1, $2, $3, $4)
				 ON CONFLICT (customer_id) DO UPDATE
				 SET plan_option_id = $2, expires_at = $3;`
	if err := tx.Commit(c); err != nil {
		r.logger.Error("database error", zap.Error(err))
		return domain.ErrDatabase
	}

	return nil
}
