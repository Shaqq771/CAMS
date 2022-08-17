-- migrate:up
CREATE TABLE IF NOT EXISTS user_product (
  id serial PRIMARY KEY,
  userid integer NOT NULL,
  product_count integer NOT NULL DEFAULT 0,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now(),
  deleted_at timestamp NULL DEFAULT NULL
)

-- migrate:down