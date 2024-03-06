-- name: CreateUserFollow :one
INSERT INTO user_follows (id, created_at, updated_at, user_id, feed_id)
VALUES ($1,$2, $3, $4, $5)
RETURNING *;

-- name: GetUserFollows :many
SELECT * FROM user_follows WHERE user_id=$1;

-- name: DeleteUserFollow :exec
DELETE FROM user_follows WHERE id = $1 AND user_id = $2;
