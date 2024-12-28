// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package models

import (
	"database/sql/driver"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type History string

const (
	History7   History = "7"
	History14  History = "14"
	History30  History = "30"
	History90  History = "90"
	History180 History = "180"
	History365 History = "365"
)

func (e *History) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = History(s)
	case string:
		*e = History(s)
	default:
		return fmt.Errorf("unsupported scan type for History: %T", src)
	}
	return nil
}

type NullHistory struct {
	History History `json:"history"`
	Valid   bool    `json:"valid"` // Valid is true if History is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullHistory) Scan(value interface{}) error {
	if value == nil {
		ns.History, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.History.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullHistory) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.History), nil
}

type Plans string

const (
	PlansFree       Plans = "free"
	PlansBasic      Plans = "basic"
	PlansEnterprise Plans = "enterprise"
)

func (e *Plans) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Plans(s)
	case string:
		*e = Plans(s)
	default:
		return fmt.Errorf("unsupported scan type for Plans: %T", src)
	}
	return nil
}

type NullPlans struct {
	Plans Plans `json:"plans"`
	Valid bool  `json:"valid"` // Valid is true if Plans is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullPlans) Scan(value interface{}) error {
	if value == nil {
		ns.Plans, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Plans.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullPlans) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Plans), nil
}

type Company struct {
	ID             uuid.UUID        `json:"id"`
	Name           string           `json:"name"`
	CreatedBy      string           `json:"createdBy"`
	IsActive       pgtype.Bool      `json:"isActive"`
	SubscriptionID int32            `json:"subscriptionId"`
	CreatedAt      pgtype.Timestamp `json:"createdAt"`
	UpdatedAt      pgtype.Timestamp `json:"updatedAt"`
}

type StatusPage struct {
	ID           uuid.UUID        `json:"id"`
	Url          string           `json:"url"`
	IsActive     pgtype.Bool      `json:"isActive"`
	SupportUrl   pgtype.Text      `json:"supportUrl"`
	LogoUrl      pgtype.Text      `json:"logoUrl"`
	Timezone     pgtype.Text      `json:"timezone"`
	HistoryShows History          `json:"historyShows"`
	CompanyID    uuid.UUID        `json:"companyId"`
	CreatedAt    pgtype.Timestamp `json:"createdAt"`
	UpdatedAt    pgtype.Timestamp `json:"updatedAt"`
}

type Subscription struct {
	ID        int64            `json:"id"`
	IsActive  pgtype.Bool      `json:"isActive"`
	Plan      Plans            `json:"plan"`
	CreatedAt pgtype.Timestamp `json:"createdAt"`
	UpdatedAt pgtype.Timestamp `json:"updatedAt"`
}

type User struct {
	ID        uuid.UUID        `json:"id"`
	Email     string           `json:"email"`
	CompanyID uuid.UUID        `json:"companyId"`
	Name      string           `json:"name"`
	Password  string           `json:"password"`
	CreatedAt pgtype.Timestamp `json:"createdAt"`
	UpdatedAt pgtype.Timestamp `json:"updatedAt"`
}