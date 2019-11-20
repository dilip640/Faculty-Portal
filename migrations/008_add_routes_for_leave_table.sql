-- +migrate Up

CREATE TABLE application_route
(
    id SERIAL PRIMARY KEY,
    applier varchar(20),
    route_from varchar(20),
    route_to varchar(20),
    ccf_post varchar(30)
);

-- +migrate Down
DROP TABLE application_route;