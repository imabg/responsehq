package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/imabg/responehq/models"
	"github.com/imabg/responehq/pkg/errors"
	"github.com/imabg/responehq/pkg/logger"
	"github.com/imabg/responehq/pkg/respond"
	"github.com/imabg/responehq/pkg/types"
	"github.com/jackc/pgx/v5/pgtype"
	"net/http"
	"time"
)

type Company struct {
	queries *models.Queries
}

type ICompany interface {
	CreateCompany(w http.ResponseWriter, r *http.Request)
	UpdateCompany(w http.ResponseWriter, r *http.Request)
	UpdateCompanyOp(ctx context.Context, params types.UpdateCompanyDTO) error
}

func NewCompany(queries *models.Queries) ICompany {
	return &Company{
		queries: queries,
	}
}

func (c *Company) CreateCompany(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var company types.AddCompanyDTO
	err := respond.GetBody(r, &company)
	if err != nil {
		logger.Error(ctx, "While reading request body", err)
		respond.SendWithError(w, &errors.Error{
			Type:       errors.VALIDATION_ERROR,
			Err:        err,
			Code:       http.StatusBadRequest,
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		})
		return
	}
	newCom, err := c.queries.CreateCompany(ctx, models.CreateCompanyParams{
		ID:             uuid.New().String(),
		Name:           company.Name,
		SubscriptionID: company.SubscriptionID,
	})
	if err != nil {
		logger.DBError(ctx, "while creating new company", err.Error())
		respond.SendWithError(w, &errors.Error{
			Code:       http.StatusInternalServerError,
			Type:       errors.DATABASE_ERROR,
			Message:    err.Error(),
			Err:        err,
			StatusCode: http.StatusInternalServerError,
		})
		return
	}
	respond.Send(ctx, w, respond.Response{
		Code:    http.StatusOK,
		Message: "company created",
		Data:    newCom,
	})
}

func (c *Company) UpdateCompany(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var params types.UpdateCompanyDTO
	err := respond.GetBody(r, &params)
	if err != nil {
		logger.Error(ctx, "While reading request body", err)
		respond.SendWithError(w, &errors.Error{
			Type:    errors.VALIDATION_ERROR,
			Err:     err,
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	if err := c.UpdateCompanyOp(ctx, params); err != nil {
		logger.DBError(ctx, "while updating company", err.Error())
		respond.SendWithError(w, &errors.Error{
			Code:       http.StatusInternalServerError,
			Type:       errors.DATABASE_ERROR,
			Message:    err.Error(),
			Err:        err,
			StatusCode: http.StatusInternalServerError,
		})
		return
	}
	respond.Send(ctx, w, respond.Response{
		Code:    http.StatusOK,
		Message: "company updated",
		Data:    nil,
	})
}

func (c *Company) UpdateCompanyOp(ctx context.Context, params types.UpdateCompanyDTO) error {
	return c.queries.UpdateCompany(ctx, models.UpdateCompanyParams{
		CreatedBy:      params.CreatedBy,
		ID:             params.CompanyID,
		SubscriptionID: params.SubscriptionID,
		IsActive: pgtype.Bool{
			Bool:  params.IsActive,
			Valid: true,
		},
		UpdatedAt: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
	})
}
