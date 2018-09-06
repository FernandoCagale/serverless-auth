-- +goose Up
-- SQL in this section is executed when the migration is applied.

INSERT INTO auth (username, password, updated_at, created_at) VALUES ('golang', '$2b$10$tJf.EcMn0dNhM6jyRQ/vJeYo.J9eqGZSak9cQB1QXpxpvWxFwVF5q', now(), now());

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DELETE FROM auth;