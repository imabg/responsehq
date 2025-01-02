package types

type AddCompanyDTO struct {
	Name           string `validate:"required" json:"name"`
	SubscriptionID int32  `validate:"required" json:"subscriptionID"`
}

type UpdateCompanyDTO struct {
	Name           string `validate:"required" json:"name"`
	CompanyID      string `validate:"required" json:"companyID"`
	SubscriptionID int32  `validate:"required" json:"subscriptionID"`
	CreatedBy      string `validate:"required" json:"createdBy"`
	IsActive       bool   `validate:"required" json:"isActive"`
}
