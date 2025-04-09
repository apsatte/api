package categories_repository

import (
	"api/internal/domain"
	"api/internal/domain/menu"
	"context"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (r *repo) GetByProjectID(c context.Context, projectID uuid.UUID) ([]*menu.Category, error) {
	output := []*menu.Category{}
	sql := `SELECT
						c.id, c.project_id, c.position, c.created_at,
						tr.lang, tr.name
					FROM menu.categories as c
					LEFT JOIN menu.category_translations as tr ON tr.category_id = c.id
					WHERE c.project_id = $1 AND c.removed_at IS NULL;`
	rows, err := r.db.Query(c, sql, projectID)
	if err != nil {
		r.logger.Error("database error", zap.Error(err))
		return nil, domain.ErrDatabase
	}

	categoryMap := map[uuid.UUID]*menu.Category{}

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

		category, exists := categoryMap[ID]
		if !exists {
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

	return output, nil
}
