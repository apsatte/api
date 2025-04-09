package projects_repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOneByID(t *testing.T) {
	ctx, db, repo := setupTestDB(t)

	prjct, err := seed(ctx, db)
	if err != nil {
		t.Fail()
	}

	t.Run("success", func(t *testing.T) {
		project, err := repo.GetOneByID(ctx, prjct.ID)

		assert.Equal(t, project.ID, prjct.ID)
		assert.NoError(t, err)
	})
}
