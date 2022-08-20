-- migrate:up
CREATE TABLE IF NOT EXISTS users (
	id serial PRIMARY KEY,
	first_name VARCHAR (20) NOT NULL,
	last_name VARCHAR (20) NOT NULL,
	username VARCHAR (20) UNIQUE NOT NULL,
  email VARCHAR (60) UNIQUE NOT NULL,
	password VARCHAR (60) NOT NULL,
	last_login TIMESTAMPTZ NOT NULL,
	created_at TIMESTAMPTZ NOT NULL,
  updated_at TIMESTAMPTZ NOT NULL
);

-- migrate:down
DROP TABLE IF EXISTS users; 
