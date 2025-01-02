-- name: CreateUser :one
INSERT INTO users (id, email, name, password, company_id, subscription_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetUserById :one
SELECT *
FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1
  and is_active = TRUE;

