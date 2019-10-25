-- +migrate Up
CREATE TABLE faculty
(
    emp_id varchar(20) NOT NULL REFERENCES employee(id),
    start_date DATE,
    end_date DATE,
    dept varchar(10),
    PRIMARY KEY(emp_id, start_date)
);

CREATE TABLE cc_faculty
(
    emp_id varchar(20) NOT NULL REFERENCES employee(id),
    start_date DATE,
    end_date DATE,
    post varchar(50),
    PRIMARY KEY(emp_id, start_date)
);

-- +migrate Down
DROP TABLE faculty;
DROP TABLE cc_faculty;