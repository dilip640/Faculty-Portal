-- +migrate Up
CREATE TABLE employee
(
    id varchar(20) NOT NULL PRIMARY KEY,
    first_name varchar(20),
    last_name varchar(20),
    email varchar(50),
    passwd text
);

-- +migrate Down
DROP TABLE employee;