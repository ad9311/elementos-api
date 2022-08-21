-- migrate:up
CREATE TABLE IF NOT EXISTS invitation_codes (
	id serial PRIMARY KEY,
	code VARCHAR (6) UNIQUE NOT NULL,
  validity TIMESTAMPTZ NOT NULL,
	created_at TIMESTAMPTZ NOT NULL,
  updated_at TIMESTAMPTZ NOT NULL
);

-- migrate:down
DROP TABLE IF EXISTS invitation_codes; 
