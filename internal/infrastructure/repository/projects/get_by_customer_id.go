package projects_repository

import (
	"api/internal/domain"
	"api/internal/domain/project"
	"context"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (r *repo) GetByCustomerID(c context.Context, customerID uuid.UUID) ([]*project.Project, error) {
	output := []*project.Project{}

	sql := `SELECT p.id, p.customer_id, p.name, p.logo_url, p.background_url,
					p.service_fee, p.languages, p.created_at,
					tr.lang, tr.description,
					m.id as project_module_id, m.module_id, m.is_active as module_is_active,
					o.id as project_module_option_id, o.module_option_id, o.is_active as option_is_active,
					params.module_option_param_id, params.value
				FROM projects.projects as p
				LEFT JOIN projects.translations as tr ON tr.project_id = p.id
				LEFT JOIN projects.modules as m ON m.project_id = p.id
				LEFT JOIN projects.module_options as o ON o.project_module_id = m.id
				LEFT JOIN projects.module_option_params as params ON params.project_module_option_id = o.id
				WHERE p.customer_id = $1 AND p.removed_at IS NULL;`

	rows, err := r.db.Query(c, sql, customerID)
	if err != nil {
		r.logger.Error("database error", zap.Error(err))
		return nil, domain.ErrDatabase
	}
	defer rows.Close()

	projectMap := make(map[uuid.UUID]*project.Project)

	for rows.Next() {
		var (
			projectID, projectCustomerID     uuid.UUID
			name, logoURL, backgroundURL     string
			serviceFee                       uint
			languages                        []string
			createdAt                        time.Time
			lang, description                *string
			projectModuleID, projectOptionID uuid.UUID
			moduleID, optionID               *string
			paramID, paramValue              *string
			moduleIsActive, optionIsActive   *bool
		)

		err := rows.Scan(
			&projectID, &projectCustomerID,
			&name, &logoURL, &backgroundURL,
			&serviceFee, &languages, &createdAt,
			&lang, &description,
			&projectModuleID, &moduleID, &moduleIsActive,
			&projectOptionID, &optionID, &optionIsActive,
			&paramID, &paramValue,
		)
		if err != nil {
			r.logger.Error("database error", zap.Error(err))
			return nil, domain.ErrDatabase
		}

		prjct, exists := projectMap[projectID]
		if !exists {
			prjct = &project.Project{
				ID:            projectID,
				CustomerID:    projectCustomerID,
				Name:          name,
				LogoURL:       logoURL,
				BackgroundURL: backgroundURL,
				ServiceFee:    serviceFee,
				Languages:     languages,
				CreatedAt:     createdAt,
				Translations:  make(map[string]*project.ProjectTranslation),
				Modules:       []*project.Module{},
			}
			projectMap[projectID] = prjct
		}

		if lang != nil && description != nil {
			if _, ok := prjct.Translations[*lang]; !ok {
				prjct.Translations[*lang] = &project.ProjectTranslation{
					Description: *description,
				}
			}
		}

		if moduleID != nil {
			module := prjct.GetModuleByID(*moduleID)

			if module == nil {
				module = &project.Module{
					ID:       projectModuleID,
					ModuleID: *moduleID,
					IsActive: *moduleIsActive,
				}
				prjct.Modules = append(prjct.Modules, module)
			}

			if optionID != nil {
				option := module.GetOptionByID(*optionID)
				if option == nil {
					option = &project.Option{
						ID:       projectOptionID,
						OptionID: *optionID,
						IsActive: *optionIsActive,
					}
					module.Options = append(module.Options, option)
				}

				if paramID != nil && paramValue != nil {
					param := option.GetParamByID(*paramID)
					if param == nil {
						param = &project.Param{
							ParamID: *paramID,
							Value:   *paramValue,
						}
						option.Params = append(option.Params, param)
					}
				}
			}
		}
	}

	for _, p := range projectMap {
		output = append(output, p)
	}

	return output, nil
}
