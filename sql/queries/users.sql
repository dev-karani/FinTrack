-- name: CreateUser :one
INSERT INTO users(id, email, created_at, updated_at)
VALUES (
    gen_random_uuid(),
    $1,
    now(),
    now()
)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id= $1;

-- name: UpdateUser :one
UPDATE users
SET 
    email=$2,
    updated_at= now()
WHERE id =$1
RETURNING *;

