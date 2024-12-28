-- name: CreateStatusPage :one
INSERT INTO status_page (id, url, support_url, logo_url, timezone, history_shows, company_id)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetDetailAgainstId :one
SELECT *
FROM status_page
WHERE id = $1
LIMIT 1;

-- name: GetAllDetails :one
SELECT *
FROM status_page
         INNER JOIN companies ON companies.id = status_page.company_id
         INNER JOIN subscriptions ON subscriptions.id = companies.subscription_id
WHERE status_page.id = $1
LIMIT 1;

-- name: Update :exec
UPDATE status_page
set url=$1,
    support_url=$2,
    logo_url=$3,
    timezone=$4,
    history_shows=$5
WHERE id = $6;

