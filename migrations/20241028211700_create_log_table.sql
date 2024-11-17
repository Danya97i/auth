-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users_logs (
    id BIGSERIAL PRIMARY KEY,
    action TEXT,
    user_id BIGINT,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_logs;
-- +goose StatementEnd
