package categories_repository

import (
	"api/internal/domain"
	"api/internal/domain/menu"
	"context"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (r *repo) GetOneByID(c context.Context, ID uuid.UUID) (*menu.Category, error) {
	sql := `SELECT
						c.id, c.project_id, c.position, c.created_at,
						tr.lang, tr.name
					FROM menu.categories as c
					LEFT JOIN menu.category_translations as tr ON tr.category_id = c.id
					WHERE c.id = $1 AND c.removed_at IS NULL;`
	rows, err := r.db.Query(c, sql, ID)
	if err != nil {
		r.logger.Error("database error", zap.Error(err))
		return nil, domain.ErrDatabase
	}

	var category *menu.Category

	for rows.Next() {
		var ID, projectID uuid.UUID
		var position uint
		var createdAt time.Time
		var lang, name *string

		err := rows.Scan(&ID, &projectID, &position, &createdAt, &lang, &name)
		if err != nil {
			r.logger.Error("database error", zap.Error(err))
			return nil, domain.ErrDatabase
		}

		if category == nil {
			category = &menu.Category{
				ID:           ID,
				Translations: make(map[string]*menu.CategoryTranslation),
				Position:     position,
				CreatedAt:    createdAt,
			}
		}

		if lang != nil && name != nil {
			if _, ok := category.Translations[*lang]; !ok {
				category.Translations[*lang] = &menu.CategoryTranslation{
					Name: *name,
				}
			}
		}
	}

	return category, nil
}
