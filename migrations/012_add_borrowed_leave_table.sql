-- +migrate Up
CREATE TABLE borrowed_leave
(
    leave_id integer REFERENCES leave_application (id) PRIMARY KEY,
    no_of_days integer
);

-- +migrate Down
DROP TABLE borrowed_leave;