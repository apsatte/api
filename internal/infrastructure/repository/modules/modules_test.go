package modules_repository

import (
	"api/internal/domain/module"
	"context"
	"testing"

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
			modules.modules,
			modules.module_translations,
			modules.options,
			modules.option_translations,
			modules.params,
			modules.param_translations
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

func seedModules(ctx context.Context, conn *pgxpool.Pool) (*module.Module, error) {
	moduleID := "delivery"
	optionID := "method"
	paramID := "address"

	_, err := conn.Exec(ctx, `
		INSERT INTO modules.modules (id) VALUES ($1)
		ON CONFLICT DO NOTHING
	`, moduleID)
	if err != nil {
		return nil, err
	}

	translations := map[string]*module.ModuleTranslation{
		"en": {Name: "Delivery"},
		"ru": {Name: "Доставка"},
	}
	for lang, tr := range translations {
		_, err := conn.Exec(ctx, `
			INSERT INTO modules.module_translations (module_id, lang, name)
			VALUES ($1, $2, $3)
			ON CONFLICT DO NOTHING
		`, moduleID, lang, tr.Name)
		if err != nil {
			return nil, err
		}
	}

	_, err = conn.Exec(ctx, `
		INSERT INTO modules.options (id, module_id)
		VALUES ($1, $2)
		ON CONFLICT DO NOTHING
	`, optionID, moduleID)
	if err != nil {
		return nil, err
	}

	optionTranslations := map[string]*module.OptionTranslation{
		"en": {Name: "Method"},
		"ru": {Name: "Метод"},
	}
	for lang, tr := range optionTranslations {
		_, err := conn.Exec(ctx, `
			INSERT INTO modules.option_translations (option_id, lang, name)
			VALUES ($1, $2, $3)
			ON CONFLICT DO NOTHING
		`, optionID, lang, tr.Name)
		if err != nil {
			return nil, err
		}
	}

	_, err = conn.Exec(ctx, `
		INSERT INTO modules.params (id, option_id)
		VALUES ($1, $2)
		ON CONFLICT DO NOTHING
	`, paramID, optionID)
	if err != nil {
		return nil, err
	}

	paramTranslations := map[string]*module.ParamTranslation{
		"en": {Name: "Address"},
		"ru": {Name: "Адрес"},
	}
	for lang, tr := range paramTranslations {
		_, err := conn.Exec(ctx, `
			INSERT INTO modules.param_translations (param_id, lang, name)
			VALUES ($1, $2, $3)
			ON CONFLICT DO NOTHING
		`, paramID, lang, tr.Name)
		if err != nil {
			return nil, err
		}
	}

	mod := &module.Module{
		ID:           moduleID,
		Translations: translations,
		Options: []*module.Option{
			{
				ID:           optionID,
				Translations: optionTranslations,
				Params: []*module.Param{
					{
						ID:           paramID,
						Translations: paramTranslations,
					},
				},
			},
		},
	}

	return mod, nil
}
