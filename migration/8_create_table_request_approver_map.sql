-- migrate:up
CREATE TABLE IF NOT EXISTS request_approver_map (
  id integer PRIMARY KEY,
  request_id integer NOT NULL,
  approver_id integer NOT NULL,
  stage integer NOT NULL
)

-- migrate:down