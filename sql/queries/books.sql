-- name: CreateBook :one
INSERT INTO books (id, created_at, updated_at, name, url, user_id)
VALUES ($1,$2, $3, $4, $5, $6)
RETURNING *;

-- name: GetBooks :many
SELECT * FROM books;


