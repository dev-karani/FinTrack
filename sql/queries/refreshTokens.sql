-- name: CreateRefreshToken :one 
INSERT INTO refresh_tokens(token, created_at, updated_at, expires_at, user_id, revoked_at)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    NULL
)
RETURNING *;

-- name: RevokeRefreshToken :exec
UPDATE refresh_tokens
SET 
    revoked_at=Now(),
    updated_at=Now()
WHERE token = $1;

-- name: GetRefreshToken :one
SELECT * FROM refresh_tokens
WHERE token=$1
AND revoked_at IS NULL
AND expires_at > NOW();
