package projects_repository

import (
	"api/internal/domain"
	"api/internal/domain/project"
	"context"

	"go.uber.org/zap"
)

func (r *repo) Save(c context.Context, p *project.Project) error {
	tx, err := r.db.Begin(c)
	if err != nil {
		r.logger.Error("database error", zap.Error(err))
		return domain.ErrDatabase
	}
	defer tx.Rollback(c)

	sql := `INSERT INTO projects.projects (id, customer_id, name, logo_url, background_url, service_fee, languages, created_at)
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
					ON CONFLICT (id) DO UPDATE
					SET name = $3, logo_url = $4, background_url = $5, service_fee = $6, languages = $7`
	_, err = tx.Exec(c, sql, p.ID, p.CustomerID, p.Name, p.LogoURL, p.BackgroundURL, p.ServiceFee, p.Languages, p.CreatedAt)
	if err != nil {
		r.logger.Error("database error", zap.Error(err))
		return domain.ErrDatabase
	}

	for lang, tr := range p.Translations {
		sql := `INSERT INTO projects.translations (project_id, lang, description)
						VALUES ($1, $2, $3)
						ON CONFLICT (project_id, lang) DO UPDATE SET description = $3;`
		_, err := tx.Exec(c, sql, p.ID, lang, tr.Description)
		if err != nil {
			r.logger.Error("database error", zap.Error(err))
			return domain.ErrDatabase
		}
	}

	for _, module := range p.Modules {
		sql := `INSERT INTO projects.modules (id, project_id, module_id, is_active)
						VALUES ($1, $2, $3, $4)
						ON CONFLICT (id) DO UPDATE
						SET is_active = $4`
		_, err := tx.Exec(c, sql, module.ID, p.ID, module.ModuleID, module.IsActive)
		if err != nil {
			r.logger.Error("database error", zap.Error(err))
			return domain.ErrDatabase
		}

		for _, option := range module.Options {
			sql := `INSERT INTO projects.module_options (id, project_module_id, module_option_id, is_active)
							VALUES($1, $2, $3, $4)
							ON CONFLICT (id) DO UPDATE
							SET is_active = $4;`
			_, err := tx.Exec(c, sql, option.ID, module.ID, option.OptionID, option.IsActive)
			if err != nil {
				r.logger.Error("database error", zap.Error(err))
				return domain.ErrDatabase
			}

			for _, param := range option.Params {
				sql := `INSERT INTO projects.module_option_params (project_module_option_id, module_option_param_id, value)
								VALUES ($1, $2, $3) ON CONFLICT (id) DO UPDATE value = $3;`
				_, err := tx.Exec(c, sql, option.ID, param.ParamID, param.Value)
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
