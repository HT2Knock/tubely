-- name: GetVideos :many
SELECT id, created_at, updated_at, title, description, thumbnail_url, video_url, user_id
FROM videos
WHERE user_id = ?
ORDER BY created_at DESC;

-- name: GetVideo :one
SELECT id, created_at, updated_at, title, description, thumbnail_url, video_url, user_id
FROM videos
WHERE id = ?;

-- name: CreateVideo :one
INSERT INTO videos (
    id, title, description, user_id, created_at, updated_at
) VALUES (
    ?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
) RETURNING *;

-- name: UpdateVideo :exec
UPDATE videos
SET title = ?,
    description = ?,
    thumbnail_url = ?,
    video_url = ?,
    user_id = ?
WHERE id = ?;

-- name: DeleteVideo :exec
DELETE FROM videos
WHERE id = ?;
