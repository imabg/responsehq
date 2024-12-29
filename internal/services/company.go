package services

import (
	"github.com/google/uuid"
	"github.com/imabg/responehq/models"
	"github.com/imabg/responehq/pkg/logger"
	"github.com/imabg/responehq/pkg/respond"
	"github.com/jackc/pgx/v5/pgtype"
	"net/http"
	"time"
)

type Company struct {
	queries *models.Queries
}

type ICompany interface {
	CreateCompany(w http.ResponseWriter, r *http.Request)
	UpdateCompany(r *http.Request, params models.UpdateCompanyParams) error
}

func NewCompany(queries *models.Queries) ICompany {
	return &Company{
		queries: queries,
	}
}

func (c *Company) CreateCompany(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	company := &models.CreateCompanyParams{}
	err := respond.GetBody(ctx, r, company)
	if err != nil {
		respond.StatusInternalServerError(ctx, w, err)
	}
	newCom, err := c.queries.CreateCompany(ctx, models.CreateCompanyParams{
		ID:             uuid.New().String(),
		Name:           company.Name,
		SubscriptionID: company.SubscriptionID,
	})
	if err != nil {
		logger.Error(ctx, "Company.CreateCompany", err)
		return
	}
	respond.StatusOk(ctx, w, newCom)
	return
}

func (c *Company) UpdateCompany(r *http.Request, params models.UpdateCompanyParams) error {
	ctx := r.Context()
	if err := c.queries.UpdateCompany(ctx, models.UpdateCompanyParams{
		CreatedBy:      params.CreatedBy,
		ID:             params.ID,
		SubscriptionID: params.SubscriptionID,
		UpdatedAt: pgtype.Timestamp{
			Time: time.Now(),
		},
	}); err != nil {
		logger.Error(ctx, "Company.AddAdminUser", err)
		return err
	}
	return nil
}
