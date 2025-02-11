-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS actors(
  id TEXT NOT NULL PRIMARY KEY,
  type TEXT NOT NULL, 
  email TEXT NOT NULL,
  preferredUsername TEXT
);
CREATE UNIQUE INDEX idx_actor_id ON actors (id);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS actors;
-- +goose StatementEnd


