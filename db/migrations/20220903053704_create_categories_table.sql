-- migrate:up
CREATE TABLE IF NOT EXISTS categories (
	id serial PRIMARY KEY,
	"name" VARCHAR (50) NOT NULL,
	created_at TIMESTAMPTZ NOT NULL,
  updated_at TIMESTAMPTZ NOT NULL
);


-- migrate:down
DROP TABLE IF EXISTS categories; 
