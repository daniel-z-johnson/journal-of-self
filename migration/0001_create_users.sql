-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY, 
    username text UNIQUE, 
    email text UNIQUE, 
    password text, 
    created_at timestamp with time zone, 
    updated_at timestamp with time zone, 
    icon text
);

-- +goose Down
DROP TABLE users;
