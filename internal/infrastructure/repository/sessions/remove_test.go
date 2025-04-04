package sessions_repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemove(t *testing.T) {
	ctx, db, repo := setupTestDB(t)
	_, ssn := seedData(t, ctx, db)

	t.Run("success", func(t *testing.T) {
		err := repo.Remove(ctx, ssn)
		assert.Nil(t, err)
	})
}
