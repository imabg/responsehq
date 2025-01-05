-- name: CreateSubscriber :one
INSERT INTO subscribers (id, type, value, is_verified, page_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetAllSubscribersAgainstPage :many
SELECT *
FROM subscribers
WHERE page_id = $1;

-- name: GetSubscriberById :one
SELECT *
FROM subscribers
WHERE id = $1
  AND is_archived = FALSE;

-- name: GetSubscriberBasedOnType :many
SELECT *
FROM subscribers
WHERE type = $1
  AND is_archived = $2;

-- name: UpdateSubscriberById :exec
UPDATE subscribers
SET is_archived=$1
WHERE id = $2;