-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens (
    token, created_at, updated_at, user_id, expires_at
) VALUES (
    ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, ?, ?
) RETURNING *;

-- name: RevokeRefreshToken :exec
UPDATE refresh_tokens
SET revoked_at = CURRENT_TIMESTAMP
WHERE token = ?;

-- name: GetRefreshToken :one
SELECT token, created_at, updated_at, revoked_at, user_id, expires_at
FROM refresh_tokens
WHERE token = ?;

-- name: DeleteRefreshToken :exec
DELETE FROM refresh_tokens
WHERE token = ?;
