package customers_repository

import (
	"api/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOneByEmail(t *testing.T) {
	ctx, db, repo := setupTestDB(t)

	_, cstmr, _ := seedData(t, ctx, db)

	t.Run("success", func(t *testing.T) {
		customer, err := repo.GetOneByEmail(ctx, cstmr.Email)

		assert.Equal(t, customer.ID, cstmr.ID)
		assert.Nil(t, err)
	})

	t.Run("not found", func(t *testing.T) {
		customer, err := repo.GetOneByEmail(ctx, "noname@bad.email")

		assert.ErrorIs(t, err, domain.ErrNotFound)
		assert.Nil(t, customer)
	})
}
