-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS access_rules (
    role role NOT NULL,
    endpoint TEXT,
    PRIMARY KEY (role, endpoint)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS access_rules;
-- +goose StatementEnd
