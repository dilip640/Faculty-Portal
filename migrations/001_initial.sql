-- +migrate Up
CREATE TABLE employee
(
    emp_id varchar(10) NOT NULL PRIMARY KEY,
    first_fame varchar(20),
    last_name varchar(20),
    passwd varchar(50),
    email varchar(50)
);

-- +migrate Down
DROP TABLE employee;