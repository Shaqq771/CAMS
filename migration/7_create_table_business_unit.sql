-- migrate:up
CREATE TABLE IF NOT EXISTS business_unit (
  id integer PRIMARY KEY,
  business_unit_name varchar(50) NOT NULL,
  description varchar(100) NOT NULL,
  created_at datetime NOT NULL DEFAULT now(),
  modified_at datetime NULL DEFAULT NULL,
  created_by varchar(60) NOT NULL,
  modified_by varchar(60) NULL DEFAULT NULL
)

-- migrate:down