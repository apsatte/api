package dishes_repository

import (
	"api/internal/domain"
	"api/internal/domain/menu"
	"context"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (r *repo) GetOneByID(c context.Context, dishID uuid.UUID) (*menu.Dish, error) {
	sql := `SELECT
						d.id, d.category_id, d.position, d.photo_url,
						d.photo_mini_url, d.price, d.is_available, d.created_at,
						d_tr.lang as dish_tr_lang, d_tr.name as dish_tr_name, d_tr.description as dish_tr_description,
						ops.id as ops_id, ops.position as ops_position, ops.min_quantity as ops_min_quantity, ops.max_quantity as ops_max_quantity, 
						op_tr.lang as op_tr_lang, op_tr.name as op_tr_name,
						i.id as i_id, i.position as i_position, i.min_quantity as i_min_quantity, i.max_quantity as i_max_quantity, i.price_modifier as i_price_modifier,
						i_tr.lang as i_tr_lang, i_tr.name as i_tr_name
					FROM menu.dishes as d
					LEFT JOIN menu.dish_translations as d_tr ON d_tr.dish_id = d.id
					LEFT JOIN menu.options as ops ON ops.dish_id = d.id
					LEFT JOIN menu.option_translations as op_tr ON op_tr.option_id = ops.id
					LEFT JOIN menu.items as i ON i.option_id = ops.id
					LEFT JOIN menu.item_translations as i_tr ON i_tr.item_id = i.id
					WHERE d.id = $1`

	rows, err := r.db.Query(c, sql, dishID)
	if err != nil {
		r.logger.Error("database error", zap.Error(err))
		return nil, domain.ErrDatabase
	}

	var dish *menu.Dish

	for rows.Next() {
		// dish part
		var dishID, categoryID uuid.UUID
		var dishPosition, dishPrice uint
		var dishPhotoURL, dishPhotoMiniURL string
		var dishIsAvailable bool
		var dishCreatedAt time.Time
		var dishLang, dishName, dishDescription *string

		// optionPart
		var optionID *uuid.UUID
		var optionPosition, optionMinQuantity, optionMaxQuantity *uint
		var optionLang, optionName *string

		// optionItemPart
		var itemID *uuid.UUID
		var itemPosition, itemMinQuantity, itemMaxQuantity, itemPriceModifier *uint
		var itemLang, itemName *string

		err := rows.Scan(
			&dishID, &categoryID, &dishPosition,
			&dishPhotoURL, &dishPhotoMiniURL, &dishPrice,
			&dishIsAvailable, &dishCreatedAt,
			&dishLang, &dishName, &dishDescription,

			&optionID, &optionPosition, &optionMinQuantity,
			&optionMaxQuantity, &optionLang, &optionName,

			&itemID, &itemPosition, &itemMinQuantity, &itemMaxQuantity,
			&itemPriceModifier, &itemLang, &itemName,
		)
		if err != nil {
			r.logger.Error("database error", zap.Error(err))
			return nil, domain.ErrDatabase
		}

		if dish == nil {
			dish = &menu.Dish{
				ID:           dishID,
				CategoryID:   categoryID,
				Position:     dishPosition,
				PhotoURL:     dishPhotoURL,
				PhotoMiniURL: dishPhotoMiniURL,
				Price:        dishPrice,
				IsAvailable:  dishIsAvailable,
				Options:      []*menu.DishOption{},
				Translations: make(map[string]*menu.DishTranslation),
				CreatedAt:    dishCreatedAt,
			}
		}

		if dishLang != nil {
			if _, ok := dish.Translations[*dishLang]; !ok {
				dish.Translations[*dishLang] = &menu.DishTranslation{
					Name:        *dishName,
					Description: *dishDescription,
				}
			}
		}

		if optionID != nil {
			// option
			option, err := dish.GetOptionByID(*optionID)
			if err != nil {
				option = &menu.DishOption{
					ID:           *optionID,
					MinQuantity:  *optionMinQuantity,
					MaxQuantity:  *optionMaxQuantity,
					Position:     *optionPosition,
					Translations: make(map[string]*menu.DishOptionTranslation),
					Items:        []*menu.DishOptionItem{},
				}
				dish.AddOption(option)
			}

			// option translations
			if optionLang != nil {
				if _, ok := option.Translations[*optionLang]; !ok {
					option.Translations[*optionLang] = &menu.DishOptionTranslation{
						Name: *optionName,
					}
				}
			}

			if itemID != nil {
				// item
				item, err := option.GetItemByID(*itemID)
				if err != nil {
					item = &menu.DishOptionItem{
						ID:            *itemID,
						MinQuantity:   *itemMinQuantity,
						MaxQuantity:   *itemMaxQuantity,
						PriceModifier: *itemPriceModifier,
						Position:      *itemPosition,
						Translations:  make(map[string]*menu.DishOptionItemTranslation),
					}
					option.AddItem(item)
				}

				// item translations
				if itemLang != nil {
					if _, ok := item.Translations[*itemLang]; !ok {
						item.Translations[*itemLang] = &menu.DishOptionItemTranslation{
							Name: *itemName,
						}
					}
				}
			}
		}
	}

	return dish, nil
}
