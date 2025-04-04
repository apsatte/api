package account_usecase

import (
	"api/internal/domain/customer"
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type Subscription struct {
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}

type AuthOutput struct {
	Customer     Customer     `json:"customer"`
	Subscription Subscription `json:"subscription"`
}

func NewAuthOutput(c *customer.Customer) *AuthOutput {
	return &AuthOutput{
		Customer: Customer{
			ID:   c.ID,
			Name: c.Name,
		},
		Subscription: Subscription{
			ExpiresAt: c.Subscription.ExpiresAt,
			CreatedAt: c.Subscription.CreatedAt,
		},
	}
}
