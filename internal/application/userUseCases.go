package application

import (
	"fmt"

	"github.com/joaovds/auth-jwt/internal/domain"
	"github.com/joaovds/auth-jwt/internal/infra/cryptography"
	"github.com/joaovds/auth-jwt/internal/infra/mongo/repositories"
)

type UserUseCases struct {
	UserRepository domain.UserRepository
  Hasher domain.Cryptography
}

func NewUserUseCases() *UserUseCases {
	ur := repositories.NewUserRepository()
  hasher := cryptography.NewBcrypt(14)

	return &UserUseCases{
		UserRepository: ur,
    Hasher: hasher,
	}
}

func (u *UserUseCases) GetAll() ([]*domain.User, error) {
	users, err := u.UserRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserUseCases) GetByID(id string) (*domain.User, error) {
	user, err := u.UserRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserUseCases) Create(user *domain.User) error {
  passwordHash, err := u.Hasher.Hasher(user.Password)
  if err != nil {
    return err
  }
  user.Password = passwordHash

  err = u.UserRepository.Create(user)
  if err != nil {
    return err
  }

  return nil
}

func (u *UserUseCases) Login(email, password string) (string ,error) {
  user, err := u.UserRepository.GetByEmail(email)
  if err != nil {
    return "", err
  }

  if !u.Hasher.HashComparer(password, user.Password) {
    return "", fmt.Errorf("invalid password")
  }

  return "token", nil
}
