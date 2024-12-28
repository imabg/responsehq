package user

import (
	"github.com/google/uuid"
	"github.com/imabg/responehq/models"
	"github.com/imabg/responehq/pkg/logger"
	"github.com/imabg/responehq/pkg/respond"
	"net/http"
)

type IUser interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type User struct {
	queries *models.Queries
}

func NewUser(queries *models.Queries) IUser {
	return &User{
		queries: queries,
	}
}

func (u *User) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var user models.User
	respond.GetBody(ctx, w, user)
	data, err := u.queries.CreateUser(ctx, models.CreateUserParams{
		Email:     user.Email,
		Name:      user.Name,
		Password:  "123",
		CompanyID: uuid.New(),
		ID:        uuid.New(),
	})
	if err != nil {
		logger.Error(ctx, "user.Create", err.Error())
		return
	}
	respond.StatusOk(ctx, w, data)
	return
}
