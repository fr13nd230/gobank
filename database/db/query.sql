-- name: NewAccount :one
INSERT INTO accounts(
    owner, currency
) VALUES (
    $1, $2
)
RETURNING *;

-- name: ListAccounts :many
SELECT * 
FROM accounts
ORDER BY created_at
LIMIT $1
OFFSET $2;

-- name: FindAccountById :one
SELECT *
FROM accounts
WHERE id = $1
LIMIT 1;

-- name: UpdateAccountById :one
UPDATE accounts
SET owner = $2, currency = $3, updated_at = now()
WHERE id = $1
RETURNING owner, currency;

-- name: DeleteAccountById :exec
DELETE FROM accounts WHERE id = $1;

-- name: NewTransfer :one
INSERT INTO transfers(
    from_acc, to_acc, amount, status
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: ListTransfers :many
SELECT *
FROM transfers
ORDER BY created_at, updated_at
LIMIT $1
OFFSET $2;

-- name: FindTransferById :one
SELECT *
from transfers
WHERE id = $1
LIMIT 1;

-- name: FindTransfersByAcc :many
SELECT * 
FROM transfers
WHERE from_acc = $1 or to_acc = $1
LIMIT $2
OFFSET $3;

-- name: UpdateTransferById :one
UPDATE transfers SET status = $2, updated_at = now()
WHERE id = $1
RETURNING status;

-- name: DeleteTransferById :exec
DELETE from transfers WHERE id = $1;

-- name: NewEntry :one
INSERT INTO entries (
    account_id, transfer_id, amount
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: ListEntries :many
SELECT * 
FROM entries
ORDER BY created_at
LIMIT $1
OFFSET $2;

-- name: FindEntriesByTrAcc :many
SELECT * 
FROM entries
WHERE account_id = $1 or transfer_id = $1
ORDER BY created_at, updated_at
LIMIT $2
OFFSET $3;

-- name: FindEntryById :one
SELECT *
FROM entries
WHERE id = $1
LIMIT 1;

-- name: UpdateEntryById :one
UPDATE entries SET amount = $2, updated_at = now()
WHERE id = $1
RETURNING amount;

-- name: DeleteEntryById :exec
DELETE FROM entries WHERE id = $1;