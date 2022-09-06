-- migrate:up
ALTER TABLE landmarks
DROP COLUMN category;

-- migrate:down
ALTER TABLE landmarks
ADD COLUMN category VARCHAR (60) NOT NULL;
