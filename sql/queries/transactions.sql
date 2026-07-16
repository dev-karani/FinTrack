-- name: CreateTransaction :one
INSERT INTO transactions (id, user_id, amount, label, category, source, destination, created_at, updated_at)
VALUES(
    gen_random_uuid(),
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    now(),
    now()
)
RETURNING *;

-- name: GetAllTransactionsByUserID :many
SELECT * FROM transactions
WHERE user_id= $1
ORDER BY created_at ASC;

-- name: GetTransactionByID :one
SELECT * FROM transactions
WHERE id=$1;
