package modules_repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	ctx, db, repo := setupTestDB(t)

	_, err := seedModules(ctx, db)
	if err != nil {
		t.Fail()
	}

	t.Run("success", func(t *testing.T) {
		mdls, err := repo.GetAll(ctx)

		assert.Len(t, mdls, 1)
		assert.NoError(t, err)
	})
}
