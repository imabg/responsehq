-- name: CreateCompany :one
INSERT INTO companies (id, name, created_by, subscription_id)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetAllDetailById :one
SELECT *
FROM companies,
     subscriptions
         JOIN subscriptions ON companies.subscription_id = subscription.id
WHERE companies.id = $1
LIMIT 1;

-- name: GetDetailById :one
SELECT *
FROM companies
WHERE id = $1
LIMIT 1;

-- name: UpdateCompany :exec
UPDATE companies
set subscription_id=$1,
    is_active=$2,
    created_by = $3,
    updated_at = $4
WHERE id = $5;

-- name: DeleteCompany :exec
DELETE
FROM companies
WHERE id = $1;