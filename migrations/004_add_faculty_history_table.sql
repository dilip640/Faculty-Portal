-- +migrate Up

CREATE TABLE faculty_history
(
    emp_id varchar(20),
    dept_id integer,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL
);

CREATE TABLE cc_faculty_history
(
    emp_id varchar(20),
    post_id integer,
    start_date DATE,
    end_date DATE
);

-- +migrate Down
DROP TABLE faculty_history;
DROP TABLE cc_faculty_history;
