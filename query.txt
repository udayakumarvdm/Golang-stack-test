CREATE TABLE dates
(
  id bigserial NOT NULL,
  start_date date,
  end_date date,
  PRIMARY KEY (id))
WITH (
  OIDS=FALSE
);
