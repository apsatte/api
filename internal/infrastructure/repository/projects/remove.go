package projects_repository

import (
	"api/internal/domain"
	"api/internal/domain/project"
	"context"
	"time"

	"go.uber.org/zap"
)

func (r *repo) Remove(c context.Context, p *project.Project) error {
	sql := "UPDATE projects.projects SET removed_at = $2 WHERE id = $1;"

	_, err := r.db.Exec(c, sql, p.ID, time.Now().UTC())
	if err != nil {
		r.logger.Error("database error", zap.Error(err))
		return domain.ErrDatabase
	}

	return nil
}
