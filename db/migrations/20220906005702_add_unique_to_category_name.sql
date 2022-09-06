-- migrate:up
CREATE UNIQUE INDEX categories_name
ON categories(name);

ALTER TABLE categories
ADD CONSTRAINT categories_username_key
UNIQUE USING INDEX categories_name;

-- migrate:down
ALTER TABLE categories
DROP CONSTRAINT categories_username_key;
