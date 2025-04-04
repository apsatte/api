package plans_repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	ctx, db, repo := setupTestDB(t)
	seedData(t, ctx, db)

	t.Run("success", func(t *testing.T) {
		plans, err := repo.GetAll(ctx)

		assert.Len(t, plans, 1)
		assert.Nil(t, err)
	})
}
