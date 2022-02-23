package users

import (
	"context"
	"time"
)

type Domain struct {
	Id          uint
	Username    string
	Email       string
	Password    string
	Token       string
	Birthdate   time.Time
	Gender      string
	PhoneNumber string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

// Data Access Layer ke Domain
type DomainRepository interface {
	Create(ctx context.Context, domain Domain) (Domain, error)
	// Login(ctx context.Context, email string, password string) (Domain, error)
	// GetById(ctx context.Context, id uint) (Domain, error)
	// Update(ctx context.Context, domain Domain) (Domain, error)
}

type DomainService interface {
	Register(ctx context.Context, domain Domain) (Domain, error)
	// Login(ctx context.Context, email string, password string) (Domain, error)
	// GetById(ctx context.Context, id uint) (Domain, error)
	// Update(ctx context.Context, domain Domain) (Domain, error)
}
