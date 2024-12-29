package services

import (
	"github.com/google/uuid"
	"github.com/imabg/responehq/models"
	"github.com/imabg/responehq/pkg/logger"
	"github.com/imabg/responehq/pkg/respond"
	"github.com/jackc/pgx/v5/pgtype"
	"net/http"
)

type IUser interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}

type User struct {
	queries *models.Queries
	ctxc    ICompany
}

func NewUser(queries *models.Queries, companyCtx ICompany) IUser {
	return &User{
		queries: queries,
		ctxc:    companyCtx,
	}
}

func (u *User) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := &models.User{}
	err := respond.GetBody(ctx, r, user)
	if err != nil {
		respond.StatusInternalServerError(ctx, w, err)
	}
	hashPwd := encryptPassword(user.Password)
	data, err := u.queries.CreateUser(ctx, models.CreateUserParams{
		Email:     user.Email,
		Name:      user.Name,
		Password:  hashPwd,
		CompanyID: user.CompanyID,
		SubID:     user.SubID,
		ID:        uuid.New().String(),
	})
	if err != nil {
		logger.Error(ctx, "user.Create", err)
		return
	}

	//FIXME: as of now `is_active` and `updated_at` are not updated
	err = u.ctxc.UpdateCompany(r, models.UpdateCompanyParams{
		CreatedBy:      data.ID,
		ID:             data.CompanyID,
		SubscriptionID: data.SubID,
		IsActive:       pgtype.Bool{Bool: true},
	})
	if err != nil {
		logger.Error(ctx, "async op: user.Create", err)
		return
	}

	respond.StatusOk(ctx, w, data)
	return
}
