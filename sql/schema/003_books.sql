-- +goose Up

CREATE TABLE books (
    id UUID PRIMARY KEY,
    created_at  TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title TEXT NOT NULL,
    author TEXT NOT NULL,
    image TEXT NOT NULL,
    notes TEXT NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down 
DROP TABLE books;
