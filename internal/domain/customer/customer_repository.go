package customer

import (
	"context"

	"github.com/google/uuid"
)

type CustomersRepository interface {
	GetOneByID(c context.Context, ID uuid.UUID) (*Customer, error)
	GetOneByEmail(c context.Context, email string) (*Customer, error)
	Save(c context.Context, customer *Customer) error
}
