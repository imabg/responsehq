package services

import (
	"github.com/imabg/responehq/models"
	"github.com/imabg/responehq/pkg/errors"
	"github.com/imabg/responehq/pkg/logger"
	"github.com/imabg/responehq/pkg/respond"
	"github.com/imabg/responehq/pkg/types"
	"net/http"
)

type Subscription struct {
	queries *models.Queries
}

type ISubscription interface {
	CreateSub(w http.ResponseWriter, r *http.Request)
}

func NewSubscription(queries *models.Queries) ISubscription {
	return &Subscription{
		queries: queries,
	}
}

func (s *Subscription) CreateSub(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var sub types.AddSubscriptionDTO
	err := respond.GetBody(r, &sub)
	if err != nil {
		logger.Error(ctx, "While reading request body", err)
		respond.SendWithError(w, &errors.Error{
			Code:    http.StatusInternalServerError,
			Type:    errors.VALIDATION_ERROR,
			Message: err.Error(),
			Err:     err,
		})
		return
	}
	data, err := s.queries.CreateSubscription(ctx, sub.Plan)
	if err != nil {
		logger.DBError(ctx, "While creating new subscription", err.Error())
		respond.SendWithError(w, &errors.Error{
			Code:    http.StatusBadGateway,
			Type:    errors.DATABASE_ERROR,
			Err:     err,
			Message: err.Error(),
		})
		return
	}
	respond.Send(ctx, w, respond.Response{
		Code: http.StatusOK,
		Data: data,
	})
}
