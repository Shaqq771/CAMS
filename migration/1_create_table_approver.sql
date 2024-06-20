CREATE TABLE IF NOT EXISTS approver (
  id INT AUTO_INCREMENT PRIMARY KEY,
  approver_user_id integer NOT NULL,
  business_unit_id integer NOT NULL,
  name varchar(60) NOT NULL,
  email varchar(100) NOT NULL,
  role varchar(50) NOT NULL,
  job_title varchar(50) NOT NULL,
  department varchar(30) NOT NULL,
  location varchar(50) NOT NULL,
  business_unit varchar(30) NOT NULL,
  flag_skip_status boolean NOT NULL,
  delegate_status boolean NOT NULL,
  description varchar(100),
  created_at datetime NOT NULL DEFAULT now(),
  modified_at datetime NULL DEFAULT NULL,
  created_by varchar(60) NULL DEFAULT NULL,
  modified_by varchar(60) NULL DEFAULT NULL
)