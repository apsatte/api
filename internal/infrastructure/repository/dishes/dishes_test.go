package dishes_repository

import (
	"api/internal/domain/menu"
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func setupTestDB(t *testing.T) (context.Context, *pgxpool.Pool, *repo) {
	db, err := pgxpool.New(context.Background(), "postgresql://test:test@localhost/test")
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec(context.Background(), `
		TRUNCATE TABLE
			menu.categories,
			menu.category_translations,
			menu.dishes,
			menu.dish_translations,
			menu.options,
			menu.option_translations,
			menu.items,
			menu.item_translations
		RESTART IDENTITY CASCADE;
	`)

	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		db.Close()
	})

	return context.Background(), db, &repo{logger: zap.NewExample(), db: db}
}

func seedMenu(ctx context.Context, conn *pgxpool.Pool) (*menu.Category, []*menu.Dish, error) {
	projectID := uuid.New()
	categoryID := uuid.New()
	createdAt := time.Now()

	category := &menu.Category{
		ID:        categoryID,
		ProjectID: projectID,
		Position:  1,
		CreatedAt: createdAt,
		Translations: map[string]*menu.CategoryTranslation{
			"en": {Name: "Burgers"},
			"ru": {Name: "Бургеры"},
		},
	}

	_, err := conn.Exec(ctx, `
		INSERT INTO menu.categories (id, project_id, position, created_at)
		VALUES ($1, $2, $3, $4)
	`, category.ID, category.ProjectID, category.Position, category.CreatedAt)
	if err != nil {
		return nil, nil, err
	}

	for lang, tr := range category.Translations {
		_, err := conn.Exec(ctx, `
			INSERT INTO menu.category_translations (category_id, lang, name)
			VALUES ($1, $2, $3)
		`, category.ID, lang, tr.Name)
		if err != nil {
			return nil, nil, err
		}
	}

	dishes := []*menu.Dish{}
	for i := 1; i <= 2; i++ {
		dishID := uuid.New()
		dish := &menu.Dish{
			ID:           dishID,
			CategoryID:   category.ID,
			Position:     uint(i),
			PhotoURL:     fmt.Sprintf("https://example.com/dish%d.png", i),
			PhotoMiniURL: fmt.Sprintf("https://example.com/dish%d-mini.png", i),
			Price:        uint(500 + i*100),
			IsAvailable:  true,
			CreatedAt:    time.Now(),
			Translations: map[string]*menu.DishTranslation{
				"en": {
					Name:        fmt.Sprintf("Burger %d", i),
					Description: fmt.Sprintf("Tasty burger number %d", i),
				},
				"ru": {
					Name:        fmt.Sprintf("Бургер %d", i),
					Description: fmt.Sprintf("Вкусный бургер номер %d", i),
				},
			},
		}

		_, err := conn.Exec(ctx, `
			INSERT INTO menu.dishes (
				id, category_id, position, photo_url, photo_mini_url,
				price, is_available, created_at
			)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		`, dish.ID, dish.CategoryID, dish.Position, dish.PhotoURL, dish.PhotoMiniURL,
			dish.Price, dish.IsAvailable, dish.CreatedAt)
		if err != nil {
			return nil, nil, err
		}

		for lang, tr := range dish.Translations {
			_, err := conn.Exec(ctx, `
				INSERT INTO menu.dish_translations (dish_id, lang, name, description)
				VALUES ($1, $2, $3, $4)
			`, dish.ID, lang, tr.Name, tr.Description)
			if err != nil {
				return nil, nil, err
			}
		}

		optionID := uuid.New()
		option := &menu.DishOption{
			ID:          optionID,
			Position:    1,
			MinQuantity: 0,
			MaxQuantity: 2,
			Translations: map[string]*menu.DishOptionTranslation{
				"en": {Name: "Extra"},
				"ru": {Name: "Дополнительно"},
			},
		}

		_, err = conn.Exec(ctx, `
			INSERT INTO menu.options (
				id, dish_id, position, min_quantity, max_quantity
			) VALUES ($1, $2, $3, $4, $5)
		`, option.ID, dish.ID, option.Position, option.MinQuantity, option.MaxQuantity)
		if err != nil {
			return nil, nil, err
		}

		for lang, tr := range option.Translations {
			_, err := conn.Exec(ctx, `
				INSERT INTO menu.option_translations (option_id, lang, name)
				VALUES ($1, $2, $3)
			`, option.ID, lang, tr.Name)
			if err != nil {
				return nil, nil, err
			}
		}

		itemID := uuid.New()
		item := &menu.DishOptionItem{
			ID:            itemID,
			Position:      1,
			MinQuantity:   0,
			MaxQuantity:   1,
			PriceModifier: 50,
			Translations: map[string]*menu.DishOptionItemTranslation{
				"en": {Name: "Cheese"},
				"ru": {Name: "Сыр"},
			},
		}

		_, err = conn.Exec(ctx, `
			INSERT INTO menu.items (
				id, option_id, position, min_quantity, max_quantity, price_modifier
			) VALUES ($1, $2, $3, $4, $5, $6)
		`, item.ID, option.ID, item.Position, item.MinQuantity, item.MaxQuantity, item.PriceModifier)
		if err != nil {
			return nil, nil, err
		}

		for lang, tr := range item.Translations {
			_, err := conn.Exec(ctx, `
				INSERT INTO menu.item_translations (item_id, lang, name)
				VALUES ($1, $2, $3)
			`, item.ID, lang, tr.Name)
			if err != nil {
				return nil, nil, err
			}
		}

		option.Items = []*menu.DishOptionItem{item}
		dish.Options = []*menu.DishOption{option}
		dishes = append(dishes, dish)
	}

	return category, dishes, nil
}
