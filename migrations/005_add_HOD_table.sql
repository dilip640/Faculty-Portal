-- +migrate Up

CREATE TABLE hod
(
    faculty_id varchar(20) REFERENCES faculty (emp_id) PRIMARY KEY,
    dept_id integer UNIQUE REFERENCES faculty (dept_id),
    start_date DATE NOT NULL,
    end_date DATE
);

CREATE TABLE hod_history
(
    faculty_id varchar(20) REFERENCES faculty (emp_id) PRIMARY KEY,
    dept_id integer,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL
);

-- +migrate Down
DROP TABLE hod;
DROP TABLE hod_history;