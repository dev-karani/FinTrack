-- +goose Up
CREATE TABLE transactions(
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id),
    amount BIGINT NOT NULL,
    label TEXT NOT NULL,
    category TEXT NOT NULL,
    source TEXT NOT NULL,
    destination TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL
    );
-- +goose Down
DROP TABLE transactions;
