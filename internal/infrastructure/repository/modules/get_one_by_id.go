package modules_repository

import (
	"api/internal/domain"
	"api/internal/domain/module"
	"context"

	"go.uber.org/zap"
)

func (r *repo) GetOneByID(c context.Context, moduleID string) (*module.Module, error) {
	sql := `SELECT
						m.id, m_tr.lang, m_tr.name,
						o.id, o_tr.lang, o_tr.name,
						p.id, p_tr.lang, p_tr.name
					FROM modules.modules as m
					LEFT JOIN modules.module_translations as m_tr ON m_tr.module_id = m.id
					LEFT JOIN modules.options as o ON o.module_id = m.id
					LEFT JOIN modules.option_translations as o_tr ON o_tr.option_id = o.id
					LEFT JOIN modules.params as p ON p.option_id = o.id
					LEFT JOIN modules.param_translations as p_tr ON p_tr.param_id = p.id
					WHERE m.id = $1;`

	rows, err := r.db.Query(c, sql, moduleID)
	if err != nil {
		r.logger.Error("database error", zap.Error(err))
		return nil, domain.ErrDatabase
	}

	var output *module.Module

	for rows.Next() {
		var moduleID string
		var moduleLang, moduleName *string

		var optionID *string
		var optionLang, optionName *string

		var paramID *string
		var paramLang, paramName *string

		err := rows.Scan(
			&moduleID, &moduleLang, &moduleName,
			&optionID, &optionLang, &optionName,
			&paramID, &paramLang, &paramName,
		)
		if err != nil {
			r.logger.Error("database error", zap.Error(err))
			return nil, domain.ErrDatabase
		}

		if output == nil {
			output = &module.Module{
				ID:           moduleID,
				Options:      []*module.Option{},
				Translations: map[string]*module.ModuleTranslation{},
			}
		}

		if moduleLang != nil {
			if _, ok := output.Translations[*moduleLang]; !ok {
				output.Translations[*moduleLang] = &module.ModuleTranslation{
					Name: *moduleName,
				}
			}
		}

		if optionID != nil {
			option, err := output.GetOneOptionByID(*optionID)
			if err != nil {
				option = &module.Option{
					ID:           *optionID,
					Params:       []*module.Param{},
					Translations: map[string]*module.OptionTranslation{},
				}
				output.AddOption(option)
			}

			if optionLang != nil {
				if _, ok := option.Translations[*optionLang]; !ok {
					option.Translations[*optionLang] = &module.OptionTranslation{
						Name: *optionName,
					}
				}
			}

			if paramID != nil {
				param, err := option.GetOneParamByID(*paramID)
				if err != nil {
					param = &module.Param{
						ID:           *paramID,
						Translations: make(map[string]*module.ParamTranslation),
					}
					option.AddParam(param)
				}

				if paramLang != nil {
					if _, ok := param.Translations[*paramLang]; !ok {
						param.Translations[*paramLang] = &module.ParamTranslation{
							Name: *paramName,
						}
					}
				}
			}
		}
	}

	return output, nil
}
