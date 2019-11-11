-- +migrate Up

CREATE TABLE department
(
    id SERIAL PRIMARY KEY,
    dept_name VARCHAR(50) UNIQUE
);

CREATE TABLE post
(
    id SERIAL PRIMARY KEY,
    post_name VARCHAR(50) UNIQUE
);

-- +migrate Down
DROP TABLE department;
DROP TABLE post;