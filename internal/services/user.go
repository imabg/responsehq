package services

import (
	"github.com/google/uuid"
	"github.com/imabg/responehq/config"
	"github.com/imabg/responehq/models"
	"github.com/imabg/responehq/pkg/errors"
	"github.com/imabg/responehq/pkg/logger"
	"github.com/imabg/responehq/pkg/respond"
	"github.com/imabg/responehq/pkg/token"
	"github.com/imabg/responehq/pkg/types"
	"github.com/mdobak/go-xerrors"
	"net/http"
	"time"
)

type IUser interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type User struct {
	queries *models.Queries
	ctxC    ICompany
	config  *config.Config
}

func NewUser(queries *models.Queries, companyCtx ICompany) IUser {
	return &User{
		queries: queries,
		ctxC:    companyCtx,
		config:  config.NewConfig(),
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
	hashPwd, err := encryptPassword(user.Password, &Params{
		memory:      u.config.PwdMemory,
		iterations:  u.config.PwdIterations,
		parallelism: u.config.PwdParallelism,
		saltLength:  u.config.PwdSaltLength,
		keyLength:   u.config.PwsKeyLength,
	})
	if err != nil {
		logger.Error(ctx, "while encrypting password", err)
		respond.SendWithError(w, &errors.Error{
			Code:       http.StatusInternalServerError,
			Message:    err.Error(),
			Err:        err,
			StatusCode: http.StatusInternalServerError,
			Type:       errors.INTERNAL_SERVER_ERROR,
		})
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

func (u *User) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var login types.UserLoginDTO
	err := respond.GetBody(r, &login)
	if err != nil {
		logger.Error(ctx, "while reading request body: Login user", err)
		respond.SendWithError(w, &errors.Error{
			Code:       http.StatusBadRequest,
			Message:    err.Error(),
			Err:        err,
			StatusCode: http.StatusBadRequest,
			Type:       errors.VALIDATION_ERROR,
		})
	}
	user, err := u.queries.GetUserByEmail(ctx, login.Email)
	if err != nil {
		logger.DBError(ctx, "while getting user by email", err.Error())
		respond.SendWithError(w, &errors.Error{
			Code:    http.StatusNotFound,
			Type:    errors.DATABASE_ERROR,
			Message: xerrors.New("user not found").Error(),
			Err:     err,
		})
		return
	}
	isMatch, err := VerifyPassword(login.Password, user.Password)
	if err != nil {
		logger.DBError(ctx, "while verifying password", err.Error())
		respond.SendWithError(w, &errors.Error{
			Code:       http.StatusUnauthorized,
			Err:        err,
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
			Type:       errors.VALIDATION_ERROR,
		})
		return
	}
	if !isMatch {
		respond.SendWithError(w, &errors.Error{
			Code:       http.StatusUnauthorized,
			Err:        xerrors.New("password is wrong"),
			Type:       errors.VALIDATION_ERROR,
			Message:    "password is wrong",
			StatusCode: http.StatusUnauthorized,
		})
		return
	}
	// Generate JWT token
	tk := token.New(u.config.JwtSecret)
	accessToken, err := tk.Generate(token.CustomClaimData{
		UserId:         user.ID,
		Email:          user.Email,
		CompanyId:      user.CompanyID,
		SubscriptionId: user.SubscriptionID,
	}, 24*time.Hour)
	if err != nil {
		logger.Error(ctx, "while generating access token", err)
		respond.SendWithError(w, &errors.Error{
			Code:    http.StatusInternalServerError,
			Type:    errors.INTERNAL_SERVER_ERROR,
			Message: err.Error(),
			Err:     err,
		})
		return
	}
	respond.Send(ctx, w, respond.Response{
		Code:    http.StatusOK,
		Message: "User logged successfully",
		Data:    accessToken,
	})
}
