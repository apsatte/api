package dishes_repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	ctx, db, repo := setupTestDB(t)
	_, dishes, err := seedMenu(ctx, db)
	if err != nil {
		t.Fail()
	}

	t.Run("success", func(t *testing.T) {
		err = repo.Save(ctx, dishes[0])

		assert.NoError(t, err)
	})
}
