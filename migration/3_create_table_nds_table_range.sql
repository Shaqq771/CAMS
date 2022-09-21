-- migrate:up
CREATE TABLE IF NOT EXISTS nds_number_range (
  doc_type varchar(100) NOT null PRIMARY KEY,
  plant_id varchar(4) NOT null,
  from_number varchar(10) DEFAULT NULL,
  to_number varchar(10) DEFAULT NULL,
  last_number varchar(10) DEFAULT NULL,
  skip int DEFAULT 0
)

-- migrate:down