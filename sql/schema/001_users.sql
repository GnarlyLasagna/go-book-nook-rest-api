
-- +goose Up

CREATE TABLE users (
    id UUID PRIMARY KEY,
    created_at  TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL,
	image TEXT
);

-- +goose Down 
DROP TABLE users;
