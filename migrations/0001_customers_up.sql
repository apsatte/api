CREATE SCHEMA customers;

CREATE TABLE customers.customers (
  id         UUID PRIMARY KEY,
  email      VARCHAR(255) UNIQUE NOT NULL,
  password   VARCHAR(255) NOT NULL,
  name       VARCHAR(255) NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
  removed_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NULL
);

CREATE TABLE customers.sessions (
  access_token VARCHAR(512) PRIMARY KEY,
  customer_id  UUID NOT NULL,
  ip           VARCHAR(50) NOT NULL,
  user_agent   TEXT NOT NULL,
  created_at   TIMESTAMP WITHOUT TIME ZONE,
  removed_at   TIMESTAMP WITHOUT TIME ZONE DEFAULT NULL,

  FOREIGN KEY (customer_id) REFERENCES customers.customers (id) ON DELETE CASCADE
);


---

CREATE TABLE customers.plans (
  id             UUID PRIMARY KEY,
  projects_limit INTEGER NOT NULL,
  created_at     TIMESTAMP WITHOUT TIME ZONE
);

CREATE TABLE customers.plan_translations (
  plan_id     UUID NOT NULL,
  lang        VARCHAR(5) NOT NULL,

  name        VARCHAR(255) NOT NULL,
  description VARCHAR(2000) NOT NULL,

  FOREIGN KEY (plan_id) REFERENCES customers.plans (id) ON DELETE CASCADE
);

CREATE TABLE customers.plan_options (
  id            UUID PRIMARY KEY,
  plan_id       UUID NOT NULL,
  duration_days INTEGER NOT NULL,
  price         INTEGER NOT NULL,

  FOREIGN KEY (plan_id) REFERENCES customers.plans (id) ON DELETE CASCADE
);


---

CREATE TABLE customers.subscriptions (
  customer_id    UUID PRIMARY KEY,
  plan_option_id UUID NOT NULL,
  expires_at     TIMESTAMP WITHOUT TIME ZONE NOT NULL,
  created_at     TIMESTAMP WITHOUT TIME ZONE NOT NULL,

  FOREIGN KEY (plan_option_id) REFERENCES customers.plan_options (id) ON DELETE CASCADE
);