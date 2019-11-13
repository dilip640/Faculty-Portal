-- +migrate Up
CREATE TABLE leave_application
(
    id SERIAL UNIQUE,
    emp_id varchar(20) REFERENCES employee (id) PRIMARY KEY,
    no_of_days integer,
    time_stamp TIMESTAMP DEFAULT NOW(),
    applier varchar(20),
    route_status varchar(20),
    status varchar(20)
);

CREATE TABLE leave
(
    emp_id varchar(20) REFERENCES employee (id),
    no_of_leaves integer,
    year smallint,
    PRIMARY KEY(emp_id, year)
);

CREATE TABLE leave_application_history
(
    leave_id integer REFERENCES leave_application (id),
    emp_id varchar(20) REFERENCES employee (id),
    signed_by varchar(20) REFERENCES employee (id),
    comment text,
    status varchar(20),
    time_stamp TIMESTAMP DEFAULT NOW()
);

-- +migrate Down
DROP TABLE leave;
DROP TABLE leave_application_history;
DROP TABLE leave_application;