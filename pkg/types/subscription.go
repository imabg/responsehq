package types

import "github.com/imabg/responehq/models"

type AddSubscriptionDTO struct {
	Plan models.Plans `validate:"required" json:"plan"`
}
