CREATE SCHEMA modules;

CREATE TABLE modules.modules (
  id VARCHAR(50) PRIMARY KEY
);

CREATE TABLE modules.module_translations (
  module_id VARCHAR(50) NOT NULL,
  lang      VARCHAR(5) NOT NULL,
  name      VARCHAR(50) NOT NULL,

  FOREIGN KEY (module_id) REFERENCES modules.modules (id) ON DELETE CASCADE
);


-----
CREATE TABLE modules.options (
  id        VARCHAR(50) PRIMARY KEY,
  module_id VARCHAR(50) NOT NULL
);

CREATE TABLE modules.option_translations (
  option_id VARCHAR(50) NOT NULL,
  lang      VARCHAR(5) NOT NULL,
  name      VARCHAR(50) NOT NULL,

  FOREIGN KEY (option_id) REFERENCES modules.options (id) ON DELETE CASCADE
);


-----
CREATE TABLE modules.params (
  id        VARCHAR(50) PRIMARY KEY,
  option_id VARCHAR(50) NOT NULL
);

CREATE TABLE modules.param_translations (
  param_id VARCHAR(50) NOT NULL,
  lang     VARCHAR(5) NOT NULL,
  name     VARCHAR(50) NOT NULL,

  FOREIGN KEY (param_id) REFERENCES modules.params (id) ON DELETE CASCADE
);