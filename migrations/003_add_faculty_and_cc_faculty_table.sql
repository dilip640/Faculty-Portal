-- +migrate Up

CREATE TABLE faculty
(
    emp_id varchar(20) REFERENCES employee (id) PRIMARY KEY,
    dept_id integer REFERENCES department (id),
    start_date DATE NOT NULL
);

CREATE TABLE cc_faculty
(
    emp_id varchar(20) REFERENCES employee(id) PRIMARY KEY,
    post_id integer REFERENCES post (id),
    start_date DATE NOT NULL 
);

-- +migrate Down
DROP TABLE faculty;
DROP TABLE cc_faculty;