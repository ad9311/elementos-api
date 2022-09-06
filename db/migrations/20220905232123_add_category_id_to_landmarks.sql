-- migrate:up
ALTER TABLE landmarks
ADD COLUMN "category_id" INT NOT NULL DEFAULT 1;

ALTER TABLE landmarks
ADD CONSTRAINT category_id_fk
FOREIGN KEY(category_id)
REFERENCES categories(id) ON DELETE SET DEFAULT ON UPDATE CASCADE;

-- migrate:down
ALTER TABLE landmarks
DROP CONSTRAINT category_id_fk;

ALTER TABLE landmarks
DROP COLUMN category_id;
