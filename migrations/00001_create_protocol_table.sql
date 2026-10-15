-- +goose Up
CREATE TABLE protocol (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    default_port INTEGER,
    description TEXT
);

-- +goose Down
DROP TABLE protocol;