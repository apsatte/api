package projects_repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetByCustomerID(t *testing.T) {
	ctx, db, repo := setupTestDB(t)

	prjct, err := seed(ctx, db)
	if err != nil {
		t.Fail()
	}

	t.Run("success", func(t *testing.T) {
		projects, err := repo.GetByCustomerID(ctx, prjct.CustomerID)

		assert.Len(t, projects, 1)
		assert.NoError(t, err)
	})
}
