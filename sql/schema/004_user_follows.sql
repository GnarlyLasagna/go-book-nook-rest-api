-- +goose Up
CREATE TABLE user_follows (
    id UUID PRIMARY KEY,
    created_at  TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    books_id  UUID NOT NULL REFERENCES books(id) ON DELETE CASCADE,
    UNIQUE(user_id, books_id)
);

-- +goose Down 
DROP TABLE user_follows;
