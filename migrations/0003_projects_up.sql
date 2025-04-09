CREATE SCHEMA projects;

CREATE TABLE projects.projects (
  id             UUID PRIMARY KEY,
  customer_id    UUID NOT NULL,
  name           VARCHAR(25) NOT NULL,
  logo_url       VARCHAR(512) NOT NULL,
  background_url VARCHAR(512) NOT NULL,
  service_fee    INTEGER NOT NULL,
  languages      VARCHAR(5)[] NOT NULL,
  created_at     TIMESTAMP WITHOUT TIME ZONE NOT NULL,
  removed_at     TIMESTAMP WITHOUT TIME ZONE DEFAULT NULL
);

CREATE TABLE projects.translations (
  project_id  UUID NOT NULL,
  lang        VARCHAR(5) NOT NULL,
  description VARCHAR(512) NOT NULL,

  FOREIGN KEY (project_id) REFERENCES projects.projects (id) ON DELETE CASCADE
);

CREATE TABLE projects.modules (
  id         UUID PRIMARY KEY,
  project_id UUID NOT NULL,
  module_id  VARCHAR(50) NOT NULL,
  is_active  BOOLEAN NOT NULL,

  UNIQUE(project_id, module_id),

  FOREIGN KEY (project_id) REFERENCES projects.projects (id) ON DELETE CASCADE
);

CREATE TABLE projects.module_options (
  id                UUID PRIMARY KEY,
  project_module_id UUID NOT NULL,
  module_option_id  VARCHAR(50) NOT NULL,
  is_active         BOOLEAN NOT NULL,
  
  FOREIGN KEY (project_module_id) REFERENCES projects.modules (id) ON DELETE CASCADE
);

CREATE TABLE projects.module_option_params (
  project_module_option_id UUID NOT NULL,
  module_option_param_id   VARCHAR(50) NOT NULL,
  value                    VARCHAR(1024),
  PRIMARY KEY(project_module_option_id, module_option_param_id),

  FOREIGN KEY (project_module_option_id) REFERENCES projects.module_options (id) ON DELETE CASCADE
);