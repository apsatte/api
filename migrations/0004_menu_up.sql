CREATE SCHEMA menu;

CREATE TABLE menu.categories (
  id         UUID PRIMARY KEY,
  project_id UUID NOT NULL,
  position   INTEGER NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
  removed_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NULL
);

CREATE TABLE menu.category_translations (
  category_id UUID NOT NULL,
  lang        VARCHAR(5) NOT NULL,
  name        VARCHAR(50) NOT NULL,
  
  PRIMARY KEY (category_id, lang),
  
  FOREIGN KEY (category_id) REFERENCES menu.categories (id) ON DELETE CASCADE
);

CREATE TABLE menu.dishes (
  id             UUID PRIMARY KEY,
  category_id    UUID NOT NULL,
  position       INTEGER NOT NULL,
  photo_url      VARCHAR(512) NOT NULL,
  photo_mini_url VARCHAR(512) NOT NULL,
  price          INTEGER NOT NULL,
  is_available   BOOLEAN NOT NULL,
  created_at      TIMESTAMP WITHOUT TIME ZONE NOT NULL,

  FOREIGN KEY (category_id) REFERENCES menu.categories (id) ON DELETE CASCADE
);

CREATE TABLE menu.dish_translations (
  dish_id     UUID NOT NULL,
  lang        VARCHAR(5) NOT NULL,
  name        VARCHAR(50) NOT NULL,
  description VARCHAR(512) NOT NULL,

  PRIMARY KEY(dish_id, lang),
  FOREIGN KEY(dish_id) REFERENCES menu.dishes (id) ON DELETE CASCADE
);

CREATE TABLE menu.options (
  id           UUID PRIMARY KEY,
  dish_id      UUID NOT NULL,
  position     INTEGER NOT NULL,
  min_quantity INTEGER NOT NULL,
  max_quantity INTEGER NOT NULL,

  FOREIGN KEY(dish_id) REFERENCES menu.dishes (id) ON DELETE CASCADE
);

CREATE TABLE menu.option_translations (
  option_id UUID NOT NULL,
  lang      VARCHAR(5) NOT NULL,
  name      VARCHAR(50) NOT NULL,

  PRIMARY KEY(option_id, lang),
  FOREIGN KEY(option_id) REFERENCES menu.options (id) ON DELETE CASCADE
);

CREATE TABLE menu.items (
  id             UUID PRIMARY KEY,
  option_id      UUID NOT NULL,
  position     INTEGER NOT NULL,
  min_quantity   INTEGER NOT NULL,
  max_quantity   INTEGER NOT NULL,
  price_modifier INTEGER NOT NULL
);

CREATE TABLE menu.item_translations (
  item_id UUID NOT NULL,
  lang    VARCHAR(5) NOT NULL,
  name    VARCHAR(50),

  PRIMARY KEY(item_id, lang),
  FOREIGN KEY (item_id) REFERENCES menu.items (id) ON DELETE CASCADE
);
