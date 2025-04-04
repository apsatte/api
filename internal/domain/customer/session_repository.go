package customer

import (
	"context"

	"github.com/google/uuid"
)

type SessionsRepository interface {
	GetOneByAccessToken(c context.Context, accessToken string) (*Session, error)
	Create(c context.Context, session *Session) error
	Remove(c context.Context, session *Session) error
	RemoveAll(c context.Context, customerID uuid.UUID) error
}
