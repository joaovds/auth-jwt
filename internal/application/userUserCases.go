package application

import (
	"github.com/joaovds/auth-jwt/internal/domain"
	"github.com/joaovds/auth-jwt/internal/infra/mongo/repositories"
)

type UserUseCases struct {
	UserRepository domain.UserRepository
}

func NewUserUseCases() *UserUseCases {
	ur := repositories.NewUserRepository()

	return &UserUseCases{
		UserRepository: ur,
	}
}

func (u *UserUseCases) GetAll() ([]*domain.User, error) {
	users, err := u.UserRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}
