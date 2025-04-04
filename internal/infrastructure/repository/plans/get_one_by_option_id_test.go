package plans_repository

import (
	"api/internal/domain"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetOneByOptionID(t *testing.T) {
	ctx, db, repo := setupTestDB(t)
	pln, optionID := seedData(t, ctx, db)

	t.Run("success", func(t *testing.T) {
		plan, err := repo.GetOneByOptionID(ctx, optionID)

		assert.Equal(t, plan.ID, pln.ID)
		assert.Nil(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		plan, err := repo.GetOneByOptionID(ctx, uuid.New())

		assert.ErrorIs(t, err, domain.ErrNotFound)
		assert.Nil(t, plan)
	})
}
