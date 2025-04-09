package dishes_repository

import (
	"api/internal/domain"
	"api/internal/domain/menu"
	"context"

	"go.uber.org/zap"
)

func (r *repo) Save(c context.Context, dish *menu.Dish) error {
	tx, err := r.db.Begin(c)
	if err != nil {
		r.logger.Error("database error", zap.Error(err))
		return domain.ErrDatabase
	}
	defer tx.Rollback(c)

	sql := `INSERT INTO menu.dishes (id, category_id, position, photo_url, photo_mini_url, price, is_available, created_at)
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
					ON CONFLICT (id) DO UPDATE
					SET category_id = $2, position = $3, photo_url = $4, photo_mini_url = $5, price = $6, is_available = $7;`
	_, err = tx.Exec(c, sql, dish.ID, dish.CategoryID, dish.Position, dish.PhotoURL, dish.PhotoURL, dish.Price, dish.IsAvailable, dish.CreatedAt)
	if err != nil {
		r.logger.Error("database error", zap.Error(err))
		return domain.ErrDatabase
	}

	for lang, tr := range dish.Translations {
		sql := `INSERT INTO menu.dish_translations (dish_id, lang, name, description)
						VALUES ($1, $2, $3, $4)
						ON CONFLICT (dish_id, lang) DO UPDATE
						SET name = $3, description = $4;`
		_, err := tx.Exec(c, sql, dish.ID, lang, tr.Name, tr.Description)
		if err != nil {
			r.logger.Error("database error", zap.Error(err))
			return domain.ErrDatabase
		}
	}

	for _, option := range dish.Options {
		sql := `INSERT INTO menu.options (id, dish_id, position, min_quantity, max_quantity)
						VALUES ($1, $2, $3, $4, $5)
						ON CONFLICT (id) DO UPDATE
						SET position = $3, min_quantity = $4, max_quantity = $5;`
		_, err := tx.Exec(c, sql, option.ID, dish.ID, option.Position, option.MinQuantity, option.MaxQuantity)
		if err != nil {
			r.logger.Error("database error", zap.Error(err))
			return domain.ErrDatabase
		}

		for lang, tr := range option.Translations {
			sql := `INSERT INTO menu.option_translations (option_id, lang, name) VALUES ($1, $2, $3)
							ON CONFLICT (option_id, lang) DO UPDATE SET name = $3;`
			_, err := tx.Exec(c, sql, option.ID, lang, tr.Name)
			if err != nil {
				r.logger.Error("database error", zap.Error(err))
				return domain.ErrDatabase
			}
		}

		// options
		for _, item := range option.Items {
			sql := `INSERT INTO menu.items (id, option_id, position, min_quantity, max_quantity, price_modifier)
							VALUES ($1, $2, $3, $4, $5, $6)
							ON CONFLICT (id) DO UPDATE
							SET position = $3, min_quantity = $4, max_quantity = $5, price_modifier = $6;`
			_, err := tx.Exec(c, sql, item.ID, option.ID, item.Position, item.MinQuantity, item.MaxQuantity, item.PriceModifier)
			if err != nil {
				r.logger.Error("database error", zap.Error(err))
				return domain.ErrDatabase
			}

			for lang, tr := range item.Translations {
				sql := `INSERT INTO menu.item_translations (item_id, lang, name)
								VALUES ($1, $2, $3) ON CONFLICT(item_id, lang) DO UPDATE SET name = $3;`
				_, err := tx.Exec(c, sql, item.ID, lang, tr.Name)
				if err != nil {
					r.logger.Error("database error", zap.Error(err))
					return domain.ErrDatabase
				}
			}
		}
	}

	if err := tx.Commit(c); err != nil {
		r.logger.Error("database error", zap.Error(err))
		return domain.ErrDatabase
	}
	return nil
}
