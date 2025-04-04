package plans_repository

import (
	"api/internal/domain"
	"api/internal/domain/plan"
	"context"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (r *repo) GetOneByOptionID(c context.Context, ID uuid.UUID) (*plan.Plan, error) {
	sql := `SELECT p.id, p.projects_limit, p.created_at,
					tr.lang, tr.name, tr.description,
					o.id as option_id, o.duration_days, o.price
				FROM customers.plans as p
				LEFT JOIN customers.plan_translations as tr ON tr.plan_id = p.id
				LEFT JOIN customers.plan_options as o ON o.plan_id = p.id
				WHERE o.id = $1;
				`
	rows, err := r.db.Query(c, sql, ID)
	if err != nil {
		r.logger.Error("database error", zap.Error(err))
		return nil, domain.ErrDatabase
	}
	defer rows.Close()

	var pln *plan.Plan

	for rows.Next() {
		var planID, optionID uuid.UUID
		var projectsLimit, durationDays, price uint
		var createdAt time.Time
		var lang, name, description string

		err := rows.Scan(
			&planID, &projectsLimit, &createdAt,
			&lang, &name, &description, &optionID,
			&durationDays, &price)
		if err != nil {
			return nil, err
		}

		if pln == nil {
			pln = &plan.Plan{
				ID:            planID,
				ProjectsLimit: projectsLimit,
				Translations:  make(map[string]*plan.PlanTranslation),
				Options:       make(map[uuid.UUID]*plan.PlanOption),
				CreatedAt:     createdAt,
			}
		}

		if lang != "" && name != "" {
			if _, ok := pln.Translations[lang]; !ok {
				pln.Translations[lang] = &plan.PlanTranslation{
					Name:        name,
					Description: description,
				}
			}
		}

		if optionID != uuid.Nil {
			if _, ok := pln.Options[optionID]; !ok {
				pln.Options[optionID] = &plan.PlanOption{
					DurationDays: durationDays,
					Price:        price,
				}
			}
		}
	}

	if pln == nil {
		return nil, domain.ErrNotFound
	}

	return pln, nil
}
