-- name: CreatePage :one
INSERT INTO pages (id, url, support_url, logo_url, timezone, send_notification, history_shows, company_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetDetailAgainstId :one
SELECT *
FROM pages
WHERE id = $1;

-- name: GetAllDetails :one
SELECT *
FROM pages
         INNER JOIN companies ON companies.id = pages.company_id
         INNER JOIN subscriptions ON subscriptions.id = companies.subscription_id
WHERE pages.id = $1
LIMIT 1;

-- name: Update :exec
UPDATE pages
set url=$1,
    support_url=$2,
    logo_url=$3,
    timezone=$4,
    history_shows=$5
WHERE id = $6;

-- name: UpdateNotificationStatus :exec
UPDATE pages
SET send_notification=$1
WHERE id = $2;
