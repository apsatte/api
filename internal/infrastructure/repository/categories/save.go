package categories_repository

import (
	"api/internal/domain"
	"api/internal/domain/menu"
	"context"

	"go.uber.org/zap"
)

func (r *repo) Save(c context.Context, category *menu.Category) error {
	tx, err := r.db.Begin(c)
	if err != nil {
		r.logger.Error("database error", zap.Error(err))
		return domain.ErrDatabase
	}
	defer tx.Rollback(c)

	sql := `INSERT INTO menu.categories (id, project_id, position, created_at)
					VALUES ($1, $2, $3, $4)
					ON CONFLICT (id) DO UPDATE SET position = $3;`
	_, err = tx.Exec(c, sql, category.ID, category.ProjectID, category.Position, category.CreatedAt)
	if err != nil {
		r.logger.Error("database error", zap.Error(err))
		return domain.ErrDatabase
	}

	for lang, tr := range category.Translations {
		sql := `INSERT INTO menu.category_translations (category_id, lang, name)
						VALUES ($1, $2, $3)
						ON CONFLICT (category_id, lang) DO UPDATE SET name = $3;`
		_, err := tx.Exec(c, sql, category.ID, lang, tr.Name)
		if err != nil {
			r.logger.Error("database error", zap.Error(err))
			return domain.ErrDatabase
		}
	}

	if err := tx.Commit(c); err != nil {
		r.logger.Error("database error", zap.Error(err))
		return domain.ErrDatabase
	}

	return nil
}
