package customer

import (
	"api/internal/domain"
	"context"
)

type contextKey string

const sessionKey contextKey = "session"

func GetSessionFromContext(c context.Context) (*Session, error) {
	session, ok := c.Value(sessionKey).(*Session)
	if !ok {
		return nil, domain.ErrUnauthorized
	}

	return session, nil
}

func SetSessionToContext(c context.Context, s *Session) context.Context {
	return context.WithValue(c, sessionKey, s)
}
