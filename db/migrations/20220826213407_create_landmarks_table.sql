-- migrate:up
CREATE TABLE IF NOT EXISTS landmarks (
	id serial PRIMARY KEY,
	"name" VARCHAR (60) UNIQUE NOT NULL,
	native_name VARCHAR (60) UNIQUE NOT NULL,
	category VARCHAR (60) NOT NULL,
	"description" TEXT NOT NULL,
  wiki_url TEXT NOT NULL,
  "location" TEXT [] NOT NULL,
  "img_urls" TEXT [] NOT NULL,
  default_landmark BOOLEAN NOT NULL DEFAULT false,
  "user_id" INT NOT NULL DEFAULT 1,
	created_at TIMESTAMPTZ NOT NULL,
  updated_at TIMESTAMPTZ NOT NULL,
  CONSTRAINT user_id_fk
  FOREIGN KEY(user_id)
	REFERENCES users(id) ON DELETE SET DEFAULT ON UPDATE CASCADE
);

-- migrate:down
DROP TABLE IF EXISTS landmarks;
