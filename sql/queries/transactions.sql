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
ORDER BY created_at DESC;

-- name: GetTransactionByID :one
SELECT * FROM transactions
WHERE id=$1
AND user_id=$2
AND deleted_at IS NULL;

-- name: UpdateTransactionByID :one
UPDATE transactions
SET 
    amount=$3,
    label=$4,
    category=$5,
    source=$6,
    destination=$7,
    updated_at=now()
WHERE id=$1
AND user_id=$2
AND deleted_at IS NULL
RETURNING *;

-- name: DeleteTransactionByID :exec
UPDATE transactions
SET
    deleted_at=now()
WHERE id=$1
AND user_id=$2
AND deleted_at IS NULL;

-- name: GetUserBalance :one
SELECT  
    COALESCE(
        SUM(
            CASE 
                WHEN category = 'CREDIT' THEN amount
                WHEN category = 'DEBIT' THEN  -amount
            END
    ),
    0
)::BIGINT AS balance
FROM transactions
WHERE user_id=$1
AND deleted_at IS NULL;


-- name: GetUserExpenses :one
SELECT
COALESCE(SUM(amount), 0)::BIGINT AS total_expenses
FROM transactions
WHERE user_id=$1
AND category='DEBIT'
AND deleted_at IS NULL;

-- name: GetUserIncome :one
SELECT 
    COALESCE(SUM(amount), 0) AS total_income
FROM transactions
WHERE user_id=$1
AND  category='CREDIT'
AND deleted_at IS NULL;











