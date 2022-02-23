package users

import (
	"context"
	"project-ecommerce/bussinesses/users"
	"project-ecommerce/helpers"

	"gorm.io/gorm"
)

type UserRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) users.DomainRepository {
	return &UserRepository{Conn: conn}
}

// //////////////////////Create(ctx context.Context, domain Domain) (Domain, error)
func (u *UserRepository) Create(ctx context.Context, user users.Domain) (users.Domain, error) {

	ctx = context.TODO()

	password, err := helpers.HashPassword(user.Password)
	if err != nil {
		return users.Domain{}, err
	}

	createdUser := Users{
		Username: user.Username,
		Email:    user.Email,
		Password: password,
	}

	insertErr := u.Conn.Create(&createdUser).Error
	if insertErr != nil {
		return users.Domain{}, insertErr
	}

	return createdUser.ToDomain(), nil
}
