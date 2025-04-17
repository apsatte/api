package dishes_repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemove(t *testing.T) {
	ctx, db, repo := setupTestDB(t)

	_, dishes, err := seedMenu(ctx, db)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("success", func(t *testing.T) {
		err := repo.Remove(ctx, dishes[0])
		assert.NoError(t, err)
	})
}
