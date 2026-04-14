-- +goose Up
ALTER TABLE projects 
ADD COLUMN updated_at TIMESTAMP NOT NULL DEFAULT NOW();

-- +goose Down
ALTER TABLE projects DROP COLUMN IF EXISTS updated_at;
