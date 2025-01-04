// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package models

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, email, name, password, company_id, subscription_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, email, company_id, subscription_id, name, password, is_active, created_at, updated_at
`

type CreateUserParams struct {
	ID             string `json:"id"`
	Email          string `json:"email"`
	Name           string `json:"name"`
	Password       string `json:"password"`
	CompanyID      string `json:"companyId"`
	SubscriptionID int32  `json:"subscriptionId"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.ID,
		arg.Email,
		arg.Name,
		arg.Password,
		arg.CompanyID,
		arg.SubscriptionID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.CompanyID,
		&i.SubscriptionID,
		&i.Name,
		&i.Password,
		&i.IsActive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, email, company_id, subscription_id, name, password, is_active, created_at, updated_at
FROM users
WHERE email = $1
  and is_active = TRUE
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.CompanyID,
		&i.SubscriptionID,
		&i.Name,
		&i.Password,
		&i.IsActive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, email, company_id, subscription_id, name, password, is_active, created_at, updated_at
FROM users
WHERE id = $1
`

func (q *Queries) GetUserById(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRow(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.CompanyID,
		&i.SubscriptionID,
		&i.Name,
		&i.Password,
		&i.IsActive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
