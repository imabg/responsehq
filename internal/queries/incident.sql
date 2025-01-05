-- name: CreateIncident :one
INSERT INTO incidents(id, name, is_backfilled, body, page_id, company_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetIncidentById :one
SELECT *
FROM incidents
WHERE id = $1;

-- name: GetAllIncidentsAgainstPage :many
SELECT *
FROM incidents
WHERE page_id = $1
  AND is_active = TRUE;

-- name: UpdateIncidentById :exec
UPDATE incidents
SET name=$1,
    body=$2,
    updated_at=$3
WHERE id = $4;

-- name: DeleteIncidentById :exec
DELETE
FROM incidents
WHERE id = $1;
