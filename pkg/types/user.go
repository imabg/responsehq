package types

type AddUserDTO struct {
	Name           string `validate:"required" json:"name"`
	Email          string `validate:"required,email" json:"email"`
	SubscriptionID int32  `validate:"required,number" json:"subscriptionID"`
	CompanyID      string `validate:"required,uuid4" json:"companyId"`
	Password       string `validate:"required,min=8" json:"password"`
}
