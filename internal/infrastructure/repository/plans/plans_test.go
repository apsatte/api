package plans_repository

import (
	"api/internal/domain/customer"
	"api/internal/domain/plan"
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func setupTestDB(t *testing.T) (context.Context, *pgxpool.Pool, *repo) {
	db, err := pgxpool.New(context.Background(), "postgresql://test:test@localhost/test")
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec(context.Background(), `
		TRUNCATE TABLE
			customers.customers,
			customers.sessions,
			customers.plans,
			customers.plan_options,
			customers.subscriptions
		RESTART IDENTITY CASCADE;
	`)

	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		db.Close()
	})

	return context.Background(), db, &repo{logger: zap.NewExample(), db: db}
}

func seedData(t *testing.T, c context.Context, db *pgxpool.Pool) (*plan.Plan, uuid.UUID) {
	optionID := uuid.New()
	planOptions := map[uuid.UUID]*plan.PlanOption{
		optionID: &plan.PlanOption{
			DurationDays: 365,
			Price:        150,
		},
	}
	planTranslations := map[string]*plan.PlanTranslation{
		"en": &plan.PlanTranslation{
			Name:        "Standart",
			Description: "empty",
		},
	}
	plan, _ := plan.NewPlan(3, planTranslations, planOptions)

	_, err := db.Exec(
		c,
		"INSERT INTO customers.plans (id, projects_limit, created_at) VALUES ($1, $2, $3);",
		plan.ID, plan.ProjectsLimit, plan.CreatedAt,
	)

	require.Nil(t, err)

	_, err = db.Exec(
		c,
		"INSERT INTO customers.plan_options (id, plan_id, duration_days, price) VALUES ($1, $2, $3, $4);",
		optionID, plan.ID, planOptions[optionID].DurationDays, planOptions[optionID].Price,
	)
	require.Nil(t, err)

	_, err = db.Exec(
		c,
		"INSERT INTO customers.plan_translations (plan_id, lang, name, description) VALUES ($1, $2, $3, $4);",
		plan.ID, "en", planTranslations["en"].Name, planTranslations["en"].Description,
	)
	require.Nil(t, err)

	// customer
	subscription := &customer.Subscription{
		PlanOptionID: optionID,
		ExpiresAt:    time.Now().UTC().Add(time.Hour * 24 * 365),
		CreatedAt:    time.Now().UTC(),
	}
	cstmr, _ := customer.NewCustomer("user@example.com", "12346578", "John Doe", subscription)

	_, err = db.Exec(
		c,
		"INSERT INTO customers.customers (id, email, password, name, created_at) VALUES ($1, $2, $3, $4, $5);",
		cstmr.ID, cstmr.Email, cstmr.Password, cstmr.Name, cstmr.CreatedAt,
	)
	require.Nil(t, err)

	_, err = db.Exec(
		c,
		"INSERT INTO customers.subscriptions (customer_id, plan_option_id, expires_at, created_at) VALUES ($1, $2, $3, $4);",
		cstmr.ID, optionID, subscription.ExpiresAt, subscription.CreatedAt,
	)
	require.Nil(t, err)

	session, _ := customer.NewSession(cstmr.ID, "", "")
	_, err = db.Exec(
		c,
		"INSERT INTO customers.sessions (access_token, customer_id, ip, user_agent, created_at) VALUES ($1, $2, $3, $4, $5);",
		session.AccessToken, session.CustomerID, session.IP, session.UserAgent, session.CreatedAt,
	)
	require.Nil(t, err)

	return plan, optionID
}
