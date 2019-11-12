-- +migrate Up

CREATE TABLE hod
(
    emp_id varchar(20) REFERENCES employee (id) PRIMARY KEY,
    dept_id integer REFERENCES department (id) UNIQUE,
    start_date DATE NOT NULL,
    end_date DATE
);

CREATE TABLE hod_history
(
    emp_id varchar(20),
    dept_id integer,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL
);

-- +migrate Down
DROP TABLE hod;
DROP TABLE hod_history;