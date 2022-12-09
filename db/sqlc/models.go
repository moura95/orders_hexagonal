// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package sqlc

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type AvailableCatalogsByCompany struct {
	ID        int32
	CompanyID sql.NullInt32
	UserID    sql.NullInt32
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AvailableFactoriesByCompany struct {
	ID        int32
	CompanyID sql.NullInt32
	UserID    sql.NullInt32
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Catalog struct {
	ID        int32
	CompanyID sql.NullInt32
	UserID    sql.NullInt32
	FactoryID sql.NullInt32
	Data      sql.NullTime
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Company struct {
	ID        int32
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Contact struct {
	ID                int32
	CompanyID         sql.NullInt32
	UserID            sql.NullInt32
	Email             sql.NullString
	Website           sql.NullString
	Address           sql.NullString
	InscricaoEstadual sql.NullString
	Cnpj              sql.NullString
	Name              sql.NullString
	Cellphone         sql.NullString
	LogoUrl           sql.NullString
	FantasyName       sql.NullString
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type Costumer struct {
	ID        int32
	CompanyID sql.NullInt32
	UserID    sql.NullInt32
	ContactID sql.NullInt32
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Factory struct {
	ID        int32
	CompanyID sql.NullInt32
	ContactID sql.NullInt32
	UserID    sql.NullInt32
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Item struct {
	ID          int32
	CompanyID   sql.NullInt32
	CatalogID   sql.NullInt32
	FactoryID   sql.NullInt32
	UserID      sql.NullInt32
	Code        sql.NullString
	Reference   sql.NullString
	Description sql.NullString
	ImageUrl    sql.NullString
	Price       sql.NullFloat64
	Ipi         sql.NullFloat64
	Discount    sql.NullInt16
	Quantity    sql.NullInt16
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ItemPrice struct {
	ID        int32
	CompanyID sql.NullInt32
	ItemID    sql.NullInt32
	UserID    sql.NullInt32
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Order struct {
	ID         int32
	CompanyID  sql.NullInt32
	CostumerID sql.NullInt32
	ContactID  sql.NullInt32
	UserID     sql.NullInt32
	FactoryID  sql.NullInt32
	SubTotal   sql.NullFloat64
	Total      sql.NullFloat64
	Ipi        sql.NullFloat64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type OrderItem struct {
	ID        int32
	CompanyID sql.NullInt32
	UserID    sql.NullInt32
	OrderID   sql.NullInt32
	Ipi       sql.NullFloat64
	Discount  sql.NullInt16
	Quantity  sql.NullInt16
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Permission struct {
	ID        int32
	CompanyID sql.NullInt32
	UserID    sql.NullInt32
	Name      sql.NullString
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PermissionRole struct {
	ID        int32
	CompanyID sql.NullInt32
	UserID    sql.NullInt32
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Role struct {
	ID        int32
	CompanyID sql.NullInt32
	UserID    sql.NullInt32
	Name      sql.NullString
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Session struct {
	ID           uuid.UUID
	Email        string
	RefreshToken string
	UserAgent    string
	ClientIp     string
	IsBlocked    bool
	ExpiresAt    time.Time
	CreatedAt    time.Time
}

type User struct {
	ID        int32
	CompanyID sql.NullInt32
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserCompany struct {
	ID        int32
	CompanyID sql.NullInt32
	UserID    sql.NullInt32
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRole struct {
	ID        int32
	CompanyID sql.NullInt32
	UserID    sql.NullInt32
	CreatedAt time.Time
	UpdatedAt time.Time
}
