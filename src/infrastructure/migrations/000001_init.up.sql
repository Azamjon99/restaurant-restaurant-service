-- restaurant schema
CREATE SCHEMA IF NOT EXISTS restaurant;

CREATE TABLE IF NOT EXISTS restaurant.menus (
	id varchar(36) PRIMARY KEY,
    entity_id VARCHAR(36),
    name VARCHAR(255),
	created_at timestamp,
	updated_at timestamp
);

CREATE INDEX IF NOT EXISTS idx_entity_menu ON restaurant.menus(entity_id);

CREATE TABLE IF NOT EXISTS restaurant.categories (
	id varchar(36) PRIMARY KEY,
    menu_id VARCHAR(36),
    name VARCHAR(255),
	description TEXT,
	created_at timestamp,
	updated_at timestamp
);

CREATE INDEX IF NOT EXISTS idx_menu_category ON restaurant.categories(menu_id);

CREATE TABLE IF NOT EXISTS restaurant.items (
	id varchar(36) PRIMARY KEY,
    category_id VARCHAR(36),
    name VARCHAR(255),
	description TEXT,
	image_url varchar(255),
	price int,
	created_at timestamp,
	updated_at timestamp
);

CREATE INDEX IF NOT EXISTS idx_category_item ON restaurant.items(category_id);
CREATE TABLE IF NOT EXISTS restaurant.orders (
	id varchar(36) PRIMARY KEY,
    restaurant_id VARCHAR(36),
    status VARCHAR(255),
	eater jsonb,
	delivery jsonb,
	payment jsonb,
	items jsonb,
	created_at timestamp,
	updated_at timestamp
);

CREATE INDEX IF NOT EXISTS idx_restaurant_order ON restaurant.orders(restaurant_id);

CREATE TABLE IF NOT EXISTS restaurant.restaurants (
	id varchar(36) PRIMARY KEY,
    entity_id VARCHAR(36),
    name VARCHAR(255),
	description text,
	address jsonb,
	type VARCHAR(100),
	created_at timestamp,
	updated_at timestamp
);

CREATE INDEX IF NOT EXISTS idx_entity_restaurant ON restaurant.restaurants(entity_id);

CREATE TABLE IF NOT EXISTS restaurant.images (
	id varchar(36) PRIMARY KEY,
    restaurant_id VARCHAR(36),
    url VARCHAR(255),
	type VARCHAR(100),
	created_at timestamp
);

CREATE INDEX IF NOT EXISTS idx_restaurant_image ON restaurant.images(restaurant_id);