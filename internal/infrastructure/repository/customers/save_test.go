package customers_repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	ctx, db, repo := setupTestDB(t)
	_, cstmr, _ := seedData(t, ctx, db)

	t.Run("success", func(t *testing.T) {
		err := repo.Save(ctx, cstmr)
		assert.Nil(t, err)
	})
}
