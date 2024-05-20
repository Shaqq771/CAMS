-- migrate:up
CREATE TABLE IF NOT EXISTS rule_history (
  id serial PRIMARY KEY,
  type varchar(50) NOT NULL,
  old_value varchar(100) NOT NULL,
  new_value varchar(100) NOT NULL,
  created_at datetime NOT NULL DEFAULT now(),
  modified_at datetime NULL DEFAULT NULL,
  created_by varchar(60) NOT NULL,
  modified_by varchar(60) NULL DEFAULT NULL
)

-- migrate:down