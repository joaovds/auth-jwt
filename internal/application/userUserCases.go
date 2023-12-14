package application

import "github.com/joaovds/auth-jwt/internal/domain"

type UserUseCases struct{}

func NewUserUseCases() *UserUseCases {
	return &UserUseCases{}
}

func (u *UserUseCases) GetAll() ([]*domain.User, error) {
	users := []*domain.User{
		{},
		{},
	}

	return users, nil
}
