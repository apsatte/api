package customers_repository

import (
	"api/internal/domain"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetOneByID(t *testing.T) {
	ctx, db, repo := setupTestDB(t)
	_, cstmr, _ := seedData(t, ctx, db)

	t.Run("success", func(t *testing.T) {
		customer, err := repo.GetOneByID(ctx, cstmr.ID)

		assert.Equal(t, customer.Email, cstmr.Email)
		assert.Nil(t, err)
	})

	t.Run("not found", func(t *testing.T) {
		customer, err := repo.GetOneByID(ctx, uuid.New())

		assert.ErrorIs(t, err, domain.ErrNotFound)
		assert.Nil(t, customer)
	})
}
