package dishes_repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOneByID(t *testing.T) {
	ctx, db, repo := setupTestDB(t)

	_, dishes, err := seedMenu(ctx, db)
	if err != nil {
		t.Fail()
	}

	t.Run("success", func(t *testing.T) {
		dish, err := repo.GetOneByID(ctx, dishes[0].ID)

		assert.NotNil(t, dish)
		assert.NoError(t, err)
	})
}
