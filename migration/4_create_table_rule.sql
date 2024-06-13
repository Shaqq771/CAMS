-- migrate:up
CREATE TABLE IF NOT EXISTS rule (
  id integer PRIMARY KEY,
  module_id integer NOT NULL,
  module varchar(30) NOT NULL,
  type varchar(50) NOT NULL,
  description varchar(100) NULL DEFAULT NULL,
  stage integer NOT NULL,
  value varchar(100) NOT NULL,
  method varchar(30) NOT NULL,
  count_levelling integer NULL,
  approver varchar(200) NOT NULL,
  reject_permission boolean NOT NULL,
  req integer NOT NULL,
  email_approver boolean NOT NULL,
  due_time integer NULL,
  case_overdue varchar(50) NULL,
  revise_method varchar(50) NOT NULL,
  revise_decision varchar(50) NOT NULL,
  delegation varchar(60) NULL,
  delegation_time integer NULL,
  created_at datetime NOT NULL DEFAULT now(),
  modified_at datetime NULL DEFAULT NULL,
  created_by varchar(60) NOT NULL,
  modified_by varchar(60) NULL DEFAULT NULL
)

-- migrate:down