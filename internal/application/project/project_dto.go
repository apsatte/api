package project_usecase

import (
	"api/internal/domain/project"

	"github.com/google/uuid"
)

type ProjectOutput struct {
	ID            uuid.UUID                            `json:"id"`
	Name          string                               `json:"name"`
	LogoURL       string                               `json:"logo_url"`
	BackgroundURL string                               `json:"background_url"`
	ServiceFee    uint                                 `json:"service_fee"`
	Languages     []string                             `json:"languages"`
	Modules       []*ProjectOutputModule               `json:"modules"`
	Translations  map[string]*ProjectOutputTranslation `json:"translations"`
}

type ProjectOutputTranslation struct {
	Description string `json:"description"`
}

type ProjectOutputModule struct {
	ID       uuid.UUID              `json:"id"`
	ModuleID string                 `json:"module_id"`
	IsActive bool                   `json:"is_active"`
	Options  []*ProjectOutputOption `json:"options"`
}

type ProjectOutputOption struct {
	ID       uuid.UUID             `json:"id"`
	OptionID string                `json:"option_id"`
	IsActive bool                  `json:"is_active"`
	Params   []*ProjectOutputParam `json:"params"`
}

type ProjectOutputParam struct {
	ParamID string `json:"param_id"`
	Value   string `json:"value"`
}

func MapProjectsToOutput(cdnBaseURL string, projects []*project.Project) []*ProjectOutput {
	result := make([]*ProjectOutput, 0, len(projects))
	for _, p := range projects {
		result = append(result, MapProjectToOutput(cdnBaseURL, p))
	}
	return result
}

func MapProjectToOutput(cdnBaseURL string, p *project.Project) *ProjectOutput {
	if p == nil {
		return nil
	}

	backgroundURL := cdnBaseURL + "/" + p.BackgroundURL
	logoURL := cdnBaseURL + "/" + p.LogoURL

	return &ProjectOutput{
		ID:            p.ID,
		Name:          p.Name,
		LogoURL:       logoURL,
		BackgroundURL: backgroundURL,
		ServiceFee:    p.ServiceFee,
		Languages:     p.Languages,
		Modules:       mapModules(p.Modules),
		Translations:  mapProjectTranslations(p.Translations),
	}
}

func mapProjectTranslations(translations map[string]*project.ProjectTranslation) map[string]*ProjectOutputTranslation {
	if translations == nil {
		return nil
	}

	result := make(map[string]*ProjectOutputTranslation, len(translations))
	for lang, t := range translations {
		result[lang] = &ProjectOutputTranslation{
			Description: t.Description,
		}
	}
	return result
}

func mapModules(modules []*project.Module) []*ProjectOutputModule {
	result := make([]*ProjectOutputModule, 0, len(modules))
	for _, m := range modules {
		result = append(result, &ProjectOutputModule{
			ID:       m.ID,
			ModuleID: m.ModuleID,
			IsActive: m.IsActive,
			Options:  mapOptions(m.Options),
		})
	}
	return result
}

func mapOptions(options []*project.Option) []*ProjectOutputOption {
	result := make([]*ProjectOutputOption, 0, len(options))
	for _, o := range options {
		result = append(result, &ProjectOutputOption{
			ID:       o.ID,
			OptionID: o.OptionID,
			IsActive: o.IsActive,
			Params:   mapParams(o.Params),
		})
	}
	return result
}

func mapParams(params []*project.Param) []*ProjectOutputParam {
	result := make([]*ProjectOutputParam, 0, len(params))
	for _, p := range params {
		result = append(result, &ProjectOutputParam{
			ParamID: p.ParamID,
			Value:   p.Value,
		})
	}
	return result
}
