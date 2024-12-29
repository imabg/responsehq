package services

import (
	"github.com/imabg/responehq/models"
	"github.com/imabg/responehq/pkg/logger"
	"github.com/imabg/responehq/pkg/respond"
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
	sub := &models.Subscription{}
	err := respond.GetBody(ctx, r, sub)
	if err != nil {
		respond.StatusInternalServerError(ctx, w, err)
	}
	data, err := s.queries.CreateSubscription(ctx, sub.Plan)
	if err != nil {
		logger.Error(ctx, "Failed to create subscription", err)
		return
	}
	respond.StatusOk(ctx, w, data)
	return
}
