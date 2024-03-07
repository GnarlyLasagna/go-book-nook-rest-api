// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: user_follows.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createUserFollow = `-- name: CreateUserFollow :one
INSERT INTO user_follows (id, created_at, updated_at, user_id, book_id)
VALUES ($1,$2, $3, $4, $5)
RETURNING id, created_at, updated_at, user_id, book_id
`

type CreateUserFollowParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	BookID    uuid.UUID
}

func (q *Queries) CreateUserFollow(ctx context.Context, arg CreateUserFollowParams) (UserFollow, error) {
	row := q.db.QueryRowContext(ctx, createUserFollow,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UserID,
		arg.BookID,
	)
	var i UserFollow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.BookID,
	)
	return i, err
}

const deleteUserFollow = `-- name: DeleteUserFollow :exec
DELETE FROM user_follows WHERE id = $1 AND user_id = $2
`

type DeleteUserFollowParams struct {
	ID     uuid.UUID
	UserID uuid.UUID
}

func (q *Queries) DeleteUserFollow(ctx context.Context, arg DeleteUserFollowParams) error {
	_, err := q.db.ExecContext(ctx, deleteUserFollow, arg.ID, arg.UserID)
	return err
}

const getUserFollows = `-- name: GetUserFollows :many
SELECT id, created_at, updated_at, user_id, book_id FROM user_follows WHERE user_id=$1
`

func (q *Queries) GetUserFollows(ctx context.Context, userID uuid.UUID) ([]UserFollow, error) {
	rows, err := q.db.QueryContext(ctx, getUserFollows, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserFollow
	for rows.Next() {
		var i UserFollow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.BookID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}