package services

import (
	"github.com/google/uuid"
	"github.com/imabg/responehq/models"
	"github.com/imabg/responehq/pkg/errors"
	"github.com/imabg/responehq/pkg/logger"
	"github.com/imabg/responehq/pkg/respond"
	"github.com/imabg/responehq/pkg/types"
	"github.com/mdobak/go-xerrors"
	"net/http"
)

type IUser interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}

type User struct {
	queries *models.Queries
	ctxC    ICompany
}

func NewUser(queries *models.Queries, companyCtx ICompany) IUser {
	return &User{
		queries: queries,
		ctxC:    companyCtx,
	}
}

func (u *User) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var user types.AddUserDTO
	err := respond.GetBody(r, &user)
	if err != nil {
		logger.Error(ctx, "while reading request body: Create user", err)
		respond.SendWithError(w, &errors.Error{
			Code:       http.StatusBadRequest,
			Message:    err.Error(),
			Err:        err,
			StatusCode: http.StatusBadRequest,
		})
		return
	}
	hashPwd := encryptPassword(user.Password)
	if hashPwd == "" {
		logger.Error(ctx, "while encrypting password: Create user", xerrors.New("password is empty"))
		return
	}

	data, err := u.queries.CreateUser(ctx, models.CreateUserParams{
		Email:          user.Email,
		Name:           user.Name,
		Password:       hashPwd,
		CompanyID:      user.CompanyID,
		SubscriptionID: user.SubscriptionID,
		ID:             uuid.New().String(),
	})
	if err != nil {
		logger.DBError(ctx, "user.Create", err.Error())
		respond.SendWithError(w, &errors.Error{
			Code:       http.StatusInternalServerError,
			Type:       errors.DATABASE_ERROR,
			Message:    err.Error(),
			Err:        err,
			StatusCode: http.StatusInternalServerError,
		})
		return
	}
	err = u.ctxC.UpdateCompanyOp(ctx, types.UpdateCompanyDTO{
		CreatedBy:      data.ID,
		CompanyID:      data.CompanyID,
		SubscriptionID: data.SubscriptionID,
		IsActive:       true,
	})

	if err != nil {
		logger.DBError(ctx, "while updating company createdBy", err.Error())
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
		Message: "User created successfully",
		Data:    data,
	})
}
