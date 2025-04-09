package module

import (
	"api/internal/domain"
)

type Module struct {
	ID           string
	Options      []*Option
	Translations map[string]*ModuleTranslation
}

type ModuleTranslation struct {
	Name string
}

func (m *Module) GetOneOptionByID(optionID string) (*Option, error) {
	for _, option := range m.Options {
		if option.ID == optionID {
			return option, nil
		}
	}

	return nil, domain.ErrNotFound
}

func (m *Module) AddOption(option *Option) error {
	for _, opt := range m.Options {
		if opt.ID == option.ID {
			return domain.ErrModuleOptionAlreadyExists
		}
	}

	m.Options = append(m.Options, option)

	return nil
}

type Option struct {
	ID           string
	Params       []*Param
	Translations map[string]*OptionTranslation
}

type OptionTranslation struct {
	Name string
}

func (o *Option) GetOneParamByID(paramID string) (*Param, error) {
	for _, param := range o.Params {
		if param.ID == paramID {
			return param, nil
		}
	}

	return nil, domain.ErrNotFound
}

func (o *Option) AddParam(param *Param) error {
	for _, prm := range o.Params {
		if prm.ID == param.ID {
			return domain.ErrOptionParamAlreadyExists
		}
	}

	o.Params = append(o.Params, param)
	return nil
}

type Param struct {
	ID           string
	Translations map[string]*ParamTranslation
}

type ParamTranslation struct {
	Name string
}
