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
AND deleted_at IS NULL
ORDER BY created_at ASC;

-- name: GetTransactionByID :one
SELECT * FROM transactions
WHERE id=$1
AND deleted_at IS NULL;

-- name: UpdateTransactionByID :one
UPDATE transactions
SET 
    amount=$3,
    label=$4,
    category=$5,
    source=$6,
    destination=$7
WHERE id=$1
AND user_id=$2
AND deleted_at IS NULL
RETURNING *;

-- name: DeleteTransactionBYID :exec
UPDATE transactions
SET
    deleted_at=True
WHERE id=$1
AND user_id=$2
AND deleted_at IS NULL
RETURNING *;
