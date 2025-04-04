package customer

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/google/uuid"
)

type Session struct {
	AccessToken string
	CustomerID  uuid.UUID
	UserAgent   string
	IP          string
	CreatedAt   time.Time
}

func NewSession(customerID uuid.UUID, userAgent, IP string) (*Session, error) {
	accessToken, err := generateRandomString(256)
	if err != nil {
		return nil, err
	}

	return &Session{
		AccessToken: accessToken,
		CustomerID:  customerID,
		UserAgent:   userAgent,
		IP:          IP,
		CreatedAt:   time.Now().UTC(),
	}, nil
}

func generateRandomString(length uint) (string, error) {
	randomData := make([]byte, length)
	_, err := rand.Read(randomData)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(randomData)
	accessToken := hex.EncodeToString(hash[:])

	return accessToken, nil
}
