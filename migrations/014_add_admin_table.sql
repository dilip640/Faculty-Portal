-- +migrate Up

CREATE TABLE admin
(
    emp_id varchar(20) REFERENCES employee (id) PRIMARY KEY
);

-- +migrate Down
DROP TABLE admin;