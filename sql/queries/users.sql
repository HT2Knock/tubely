-- name: GetUsers :many
SELECT id, created_at, updated_at, password, email
FROM users;

-- name: GetUserByEmail :one
SELECT id, created_at, updated_at, password, email
FROM users
WHERE email = ?;

-- name: GetUserByRefreshToken :one
SELECT u.id, u.created_at, u.updated_at, u.password, u.email
FROM users u
JOIN refresh_tokens rt ON u.id = rt.user_id
WHERE rt.token = ?;

-- name: CreateUser :one
INSERT INTO users (
    id, email, password, created_at, updated_at
) VALUES (
    ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
) RETURNING *;

-- name: GetUser :one
SELECT id, created_at, updated_at, password, email
FROM users
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;
