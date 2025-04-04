package sessions_repository

import (
	"api/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOneByAccessToken(t *testing.T) {
	ctx, db, repo := setupTestDB(t)
	_, session := seedData(t, ctx, db)

	t.Run("success", func(t *testing.T) {
		ssn, err := repo.GetOneByAccessToken(ctx, session.AccessToken)

		assert.Equal(t, session.CustomerID, ssn.CustomerID)
		assert.Nil(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		ssn, err := repo.GetOneByAccessToken(ctx, "invalid token")

		assert.ErrorIs(t, err, domain.ErrNotFound)
		assert.Nil(t, ssn)
	})
}
