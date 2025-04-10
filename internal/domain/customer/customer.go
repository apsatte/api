package customer

import (
	"api/internal/domain"
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	ID           uuid.UUID
	Email        string
	Password     string
	Name         string
	Subscription *Subscription
	CreatedAt    time.Time
}

type Subscription struct {
	PlanOptionID  uuid.UUID
	ProjectsLimit uint
	ExpiresAt     time.Time
	CreatedAt     time.Time
}

func NewCustomer(email, password, name string, subscription *Subscription) (*Customer, error) {
	ID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	if len(name) > 255 {
		return nil, domain.ErrNameTooLong
	}

	if len(password) > 255 {
		return nil, domain.ErrIncorrectPassword
	}

	return &Customer{
		ID:           ID,
		Email:        email,
		Password:     password,
		Name:         name,
		Subscription: subscription,
		CreatedAt:    time.Now().UTC(),
	}, nil
}

func (c *Customer) ComparePassword(password string) error {
	if c.Password == password {
		return nil
	}
	return domain.ErrIncorrectPassword
}

func (c *Customer) Update(name string) {
	c.Name = name
}

func (c *Customer) ChangePassword(oldPassword, NewPassword string) error {
	if err := c.ComparePassword(oldPassword); err != nil {
		return domain.ErrIncorrectPassword
	}

	c.Password = NewPassword
	return nil
}
