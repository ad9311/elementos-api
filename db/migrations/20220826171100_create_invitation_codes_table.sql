-- migrate:up
CREATE TABLE IF NOT EXISTS invitations (
	id serial PRIMARY KEY,
	code VARCHAR (6) UNIQUE NOT NULL,
  expires_at TIMESTAMPTZ NOT NULL,
	created_at TIMESTAMPTZ NOT NULL,
  updated_at TIMESTAMPTZ NOT NULL
);

-- migrate:down
DROP TABLE IF EXISTS invitations;
