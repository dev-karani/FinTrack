-- +goose Up
ALTER TABLE transactions
ADD CONSTRAINT valid_category
CHECK (category IN ('DEBIT', 'CREDIT'));

-- +goose Down
ALTER TABLE transactions DROP CONSTRAINT valid_category;
