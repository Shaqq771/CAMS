-- migrate:up
CREATE TABLE IF NOT EXISTS request (
  id integer PRIMARY KEY,
  rule_id integer NOT NULL,
  user_id integer NOT NULL,
  request_module_id integer NOT NULL,
  name varchar(60) NOT NULL,
  email varchar(100) NOT NULL,
  role varchar(50) NOT NULL,
  job_title varchar(50) NOT NULL,
  department varchar(30) NOT NULL,
  description varchar(100) NULL DEFAULT NULL,
  status varchar(20) NOT NULL,
  module varchar(30) NOT NULL,
  type varchar(50) NOT NULL,
  revise_decision varchar(20) NULL DEFAULT NULL,
  created_at datetime NOT NULL DEFAULT now(),
  modified_at datetime NULL DEFAULT NULL,
  created_by varchar(60) NOT NULL,
  modified_by varchar(60) NULL DEFAULT NULL
)

-- migrate:down