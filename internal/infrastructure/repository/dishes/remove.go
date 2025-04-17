package dishes_repository

import (
	"api/internal/domain"
	"api/internal/domain/menu"
	"context"

	"go.uber.org/zap"
)

func (r *repo) Remove(c context.Context, dish *menu.Dish) error {
	sql := "DELETE FROM menu.dishes WHERE id = $1"

	_, err := r.db.Exec(c, sql, dish.ID)
	if err != nil {
		r.logger.Error("database error", zap.Error(err))
		return domain.ErrDatabase
	}

	return nil
}
