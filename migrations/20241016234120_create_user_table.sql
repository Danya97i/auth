-- +goose Up
-- +goose StatementBegin
create type role as enum ('unknown', 'admin', 'user');

CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role role NOT NULL DEFAULT 'unknown',
    created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamptz
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
