-- name: CreateSubscription :one
INSERT INTO subscriptions (plan)
VALUES ($1)
RETURNING *;

-- name: GetSubscriptionById :one
SELECT *
FROM subscriptions
WHERE id = $1
LIMIT 1;

-- name: UpdateSubscriptionById :exec
UPDATE subscriptions
set plan = $1
WHERE id = $2;

-- name: MarkSubscriptionInactive :one
UPDATE subscriptions
set is_active = false
WHERE id = $1
RETURNING *;

-- name: ListAllSubscriptions :many
SELECT *
FROM subscriptions
WHERE is_active = true
ORDER BY created_at;


