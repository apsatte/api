package module

import (
	"context"
)

type ModulesRepository interface {
	GetAll(c context.Context) ([]*Module, error)
	GetOneByID(c context.Context, moduleID string) (*Module, error)
}
