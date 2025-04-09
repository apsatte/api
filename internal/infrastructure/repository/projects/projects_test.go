package projects_repository

import (
	"api/internal/domain/project"
	"context"
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
			projects.projects,
			projects.translations,
			projects.modules,
			projects.module_options,
			projects.module_option_params
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

func seed(ctx context.Context, conn *pgxpool.Pool) (*project.Project, error) {
	projectID := uuid.New()
	customerID := uuid.New()
	moduleID := uuid.New()
	optionID := uuid.New()
	createdAt := time.Now().UTC()

	prjct := &project.Project{
		ID:            projectID,
		CustomerID:    customerID,
		Name:          "Burger Place",
		LogoURL:       "https://example.com/logo.png",
		BackgroundURL: "https://example.com/bg.png",
		ServiceFee:    5,
		Languages:     []string{"en", "ru"},
		CreatedAt:     createdAt,
		Translations: map[string]*project.ProjectTranslation{
			"en": {Description: "Best burgers in town"},
			"ru": {Description: "Лучшие бургеры в городе"},
		},
	}

	module := &project.Module{
		ID:       moduleID,
		ModuleID: "delivery",
		IsActive: true,
	}
	prjct.Modules = []*project.Module{module}

	option := &project.Option{
		ID:       optionID,
		OptionID: "show_prices",
		IsActive: true,
		Params: []*project.Param{
			{ParamID: "currency", Value: "USD"},
		},
	}
	module.Options = []*project.Option{option}

	_, err := conn.Exec(ctx, `
		INSERT INTO projects.projects (
			id, customer_id, name, logo_url, background_url,
			service_fee, languages, created_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`, prjct.ID, prjct.CustomerID, prjct.Name, prjct.LogoURL, prjct.BackgroundURL,
		prjct.ServiceFee, prjct.Languages, prjct.CreatedAt)
	if err != nil {
		return nil, err
	}

	for lang, tr := range prjct.Translations {
		_, err := conn.Exec(ctx, `
			INSERT INTO projects.translations (project_id, lang, description)
			VALUES ($1, $2, $3)
		`, prjct.ID, lang, tr.Description)
		if err != nil {
			return nil, err
		}
	}

	_, err = conn.Exec(ctx, `
		INSERT INTO projects.modules (id, project_id, module_id, is_active)
		VALUES ($1, $2, $3, $4)
	`, module.ID, prjct.ID, module.ModuleID, module.IsActive)
	if err != nil {
		return nil, err
	}

	_, err = conn.Exec(ctx, `
		INSERT INTO projects.module_options (id, project_module_id, module_option_id, is_active)
		VALUES ($1, $2, $3, $4)
	`, option.ID, module.ID, option.OptionID, option.IsActive)
	if err != nil {
		return nil, err
	}

	for _, param := range option.Params {
		_, err = conn.Exec(ctx, `
			INSERT INTO projects.module_option_params (
				project_module_option_id, module_option_param_id, value
			)
			VALUES ($1, $2, $3)
		`, option.ID, param.ParamID, param.Value)
		if err != nil {
			return nil, err
		}
	}

	return prjct, nil
}
