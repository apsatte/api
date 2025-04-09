package categories_repository

import (
	"api/internal/domain"
	"api/internal/domain/menu"
	"context"
	"time"

	"go.uber.org/zap"
)

func (r *repo) Remove(c context.Context, category *menu.Category) error {
	sql := `UPDATE menu.categories SET removed_at = $2 WHERE id = $1;`

	_, err := r.db.Exec(c, sql, category.ID, time.Now().UTC())
	if err != nil {
		r.logger.Error("database error", zap.Error(err))
		return domain.ErrDatabase
	}

	return nil
}
