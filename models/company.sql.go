// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: company.sql

package models

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createCompany = `-- name: CreateCompany :one
INSERT INTO companies (id, name, created_by, subscription_id) VALUES ($1, $2, $3, $4) RETURNING id, name, created_by, is_active, subscription_id, created_at, updated_at
`

type CreateCompanyParams struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	CreatedBy      string    `json:"createdBy"`
	SubscriptionID int32     `json:"subscriptionId"`
}

func (q *Queries) CreateCompany(ctx context.Context, arg CreateCompanyParams) (Company, error) {
	row := q.db.QueryRow(ctx, createCompany,
		arg.ID,
		arg.Name,
		arg.CreatedBy,
		arg.SubscriptionID,
	)
	var i Company
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedBy,
		&i.IsActive,
		&i.SubscriptionID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteCompany = `-- name: DeleteCompany :exec
DELETE FROM companies WHERE id = $1
`

func (q *Queries) DeleteCompany(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteCompany, id)
	return err
}

const getAllDetailById = `-- name: GetAllDetailById :one
SELECT companies.id, name, created_by, companies.is_active, subscription_id, companies.created_at, companies.updated_at, subscriptions.id, subscriptions.is_active, subscriptions.plan, subscriptions.created_at, subscriptions.updated_at, subscriptions.id, subscriptions.is_active, subscriptions.plan, subscriptions.created_at, subscriptions.updated_at FROM companies, subscriptions
JOIN subscriptions ON companies.subscription_id = subscription.id WHERE companies.id = $1 LIMIT 1
`

type GetAllDetailByIdRow struct {
	ID             uuid.UUID        `json:"id"`
	Name           string           `json:"name"`
	CreatedBy      string           `json:"createdBy"`
	IsActive       pgtype.Bool      `json:"isActive"`
	SubscriptionID int32            `json:"subscriptionId"`
	CreatedAt      pgtype.Timestamp `json:"createdAt"`
	UpdatedAt      pgtype.Timestamp `json:"updatedAt"`
	ID_2           int64            `json:"id2"`
	IsActive_2     pgtype.Bool      `json:"isActive2"`
	Plan           Plans            `json:"plan"`
	CreatedAt_2    pgtype.Timestamp `json:"createdAt2"`
	UpdatedAt_2    pgtype.Timestamp `json:"updatedAt2"`
	ID_3           int64            `json:"id3"`
	IsActive_3     pgtype.Bool      `json:"isActive3"`
	Plan_2         Plans            `json:"plan2"`
	CreatedAt_3    pgtype.Timestamp `json:"createdAt3"`
	UpdatedAt_3    pgtype.Timestamp `json:"updatedAt3"`
}

func (q *Queries) GetAllDetailById(ctx context.Context, id uuid.UUID) (GetAllDetailByIdRow, error) {
	row := q.db.QueryRow(ctx, getAllDetailById, id)
	var i GetAllDetailByIdRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedBy,
		&i.IsActive,
		&i.SubscriptionID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ID_2,
		&i.IsActive_2,
		&i.Plan,
		&i.CreatedAt_2,
		&i.UpdatedAt_2,
		&i.ID_3,
		&i.IsActive_3,
		&i.Plan_2,
		&i.CreatedAt_3,
		&i.UpdatedAt_3,
	)
	return i, err
}

const getDetailById = `-- name: GetDetailById :one
SELECT  id, name, created_by, is_active, subscription_id, created_at, updated_at FROM companies
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetDetailById(ctx context.Context, id uuid.UUID) (Company, error) {
	row := q.db.QueryRow(ctx, getDetailById, id)
	var i Company
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedBy,
		&i.IsActive,
		&i.SubscriptionID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateCompany = `-- name: UpdateCompany :exec
UPDATE companies set subscription_id=$1, is_active=$2 WHERE id = $2
`

type UpdateCompanyParams struct {
	SubscriptionID int32       `json:"subscriptionId"`
	IsActive       pgtype.Bool `json:"isActive"`
}

func (q *Queries) UpdateCompany(ctx context.Context, arg UpdateCompanyParams) error {
	_, err := q.db.Exec(ctx, updateCompany, arg.SubscriptionID, arg.IsActive)
	return err
}
