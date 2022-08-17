-- migrate:up
CREATE TABLE IF NOT EXISTS product (
  id serial PRIMARY KEY,
  name varchar(50) NOT NULL,
  sku varchar(12) NOT NULL,
  price integer NOT NULL DEFAULT 0,
  uom varchar(10) NOT NULL,
  stock integer NOT NULL DEFAULT 0,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now(),
  deleted_at timestamp NULL DEFAULT NULL
)

-- migrate:down