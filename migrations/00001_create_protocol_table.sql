-- +goose Up
CREATE TABLE protocols (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    default_port INTEGER,
    description TEXT
);

-- +goose Down
DROP TABLE protocols;