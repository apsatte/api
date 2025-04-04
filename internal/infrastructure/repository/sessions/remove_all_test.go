package sessions_repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveAll(t *testing.T) {
	ctx, db, repo := setupTestDB(t)
	cstmr, _ := seedData(t, ctx, db)

	t.Run("success", func(t *testing.T) {
		err := repo.RemoveAll(ctx, cstmr.ID)
		assert.Nil(t, err)
	})
}
