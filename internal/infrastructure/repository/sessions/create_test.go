package sessions_repository

import (
	"api/internal/domain"
	"api/internal/domain/customer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	ctx, db, repo := setupTestDB(t)
	cstmr, session := seedData(t, ctx, db)

	t.Run("success", func(t *testing.T) {
		session, _ := customer.NewSession(cstmr.ID, "", "")

		err := repo.Create(ctx, session)
		assert.Nil(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		err := repo.Create(ctx, session)
		assert.ErrorIs(t, err, domain.ErrDatabase)
	})
}
