package projects_repository

import (
	"api/internal/domain/project"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	ctx, _, repo := setupTestDB(t)

	t.Run("success/create", func(t *testing.T) {
		prjct, err := project.New(uuid.New(), "Untitled", "logo.jpeg", "background.jpeg", []string{"kk", "ru"})
		if err != nil {
			t.Fail()
		}

		err = repo.Save(ctx, prjct)
		assert.NoError(t, err)
	})

	t.Run("success/update", func(t *testing.T) {
		prjct, err := project.New(uuid.New(), "Untitled", "logo.jpeg", "background.jpeg", []string{"kk", "ru"})
		if err != nil {
			t.Fail()
		}

		err = repo.Save(ctx, prjct)
		assert.NoError(t, err)

		prjct.AddModule("delivery", []*project.Option{})

		err = repo.Save(ctx, prjct)
		assert.NoError(t, err)
	})
}
