package project

import (
	"api/internal/domain"
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID            uuid.UUID
	CustomerID    uuid.UUID
	Name          string
	LogoURL       string
	BackgroundURL string
	ServiceFee    uint
	Languages     []string
	Modules       []*Module
	Translations  map[string]*ProjectTranslation
	CreatedAt     time.Time
}

type ProjectTranslation struct {
	Description string
}

type Module struct {
	ID       uuid.UUID
	ModuleID string
	IsActive bool
	Options  []*Option
}

type Option struct {
	ID       uuid.UUID
	OptionID string
	IsActive bool
	Params   []*Param
}

type Param struct {
	ParamID string
	Value   string
}

func New(
	customerID uuid.UUID,
	name, logoURL, backgroundURL string,
	serviceFee uint,
	languages []string) (*Project, error) {
	ID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &Project{
		ID:            ID,
		CustomerID:    customerID,
		Name:          name,
		LogoURL:       logoURL,
		BackgroundURL: backgroundURL,
		ServiceFee:    serviceFee,
		Languages:     languages,
		CreatedAt:     time.Now().UTC(),
	}, nil
}

func (p *Project) GetModuleByID(moduleID string) *Module {
	for _, module := range p.Modules {
		if module.ModuleID == moduleID {
			return module
		}
	}
	return nil
}

func (p *Project) AddModule(ID string, options []*Option) error {
	for _, module := range p.Modules {
		if module.ModuleID == ID {
			return domain.ErrProjectModuleAlreadyExists
		}
	}

	p.Modules = append(p.Modules, &Module{
		ID:       uuid.New(),
		ModuleID: ID,
		Options:  options,
	})

	return nil
}

///////

func (p *Module) GetOptionByID(optionID string) *Option {
	for _, option := range p.Options {
		if option.OptionID == optionID {
			return option
		}
	}
	return nil
}

///

func (p *Option) GetParamByID(paramID string) *Param {
	for _, param := range p.Params {
		if param.ParamID == paramID {
			return param
		}
	}
	return nil
}
