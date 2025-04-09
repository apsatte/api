package dishes_repository

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetByCategoryID(t *testing.T) {
	ctx, db, repo := setupTestDB(t)

	category, dshs, err := seedMenu(ctx, db)
	if err != nil {
		t.Fail()
	}

	t.Run("success", func(t *testing.T) {
		dishes, err := repo.GetByCategoryID(ctx, category.ID)

		assert.Len(t, dishes, len(dshs))
		assert.NoError(t, err)
	})

	t.Run("success/unknowCategory", func(t *testing.T) {
		dishes, err := repo.GetByCategoryID(ctx, uuid.New())

		assert.Len(t, dishes, 0)
		assert.NoError(t, err)
	})
}
