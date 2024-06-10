-- migrate:up
CREATE TABLE IF NOT EXISTS module (
  id integer PRIMARY KEY,
  module_name varchar(30) NOT NULL,
  type_count integer NOT NULL,
  description varchar(100) NOT NULL,
  created_at datetime NOT NULL DEFAULT now(),
  modified_at datetime NULL DEFAULT NULL,
  created_by varchar(60) NOT NULL,
  modified_by varchar(60) NULL DEFAULT NULL
)

-- migrate:down