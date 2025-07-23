-- name: CreateUser :one
INSERT INTO users (id, name, created_at, updated_at)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetUser :one
SELECT *
FROM users
WHERE NAME == $1;