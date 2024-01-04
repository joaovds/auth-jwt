package application

import (
	"fmt"
	"time"

	"github.com/joaovds/auth-jwt/internal/domain"
	"github.com/joaovds/auth-jwt/internal/infra/cryptography"
	"github.com/joaovds/auth-jwt/internal/infra/mongo/repositories"
)

type UserUseCases struct {
	UserRepository domain.UserRepository
  Crypt domain.Cryptography
}

func NewUserUseCases() *UserUseCases {
	ur := repositories.NewUserRepository()
  crypt := cryptography.NewCryptography()

	return &UserUseCases{
		UserRepository: ur,
    Crypt: crypt,
	}
}

func (u *UserUseCases) GetAll() ([]*domain.User, error) {
	users, err := u.UserRepository.GetAll()
	if err != nil {
		return nil, err
	}

  for _, user := range users {
    user.Password = ""
  }

	return users, nil
}

func (u *UserUseCases) GetByID(id string) (*domain.User, error) {
	user, err := u.UserRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

  user.Password = ""

	return user, nil
}

func (u *UserUseCases) Create(user *domain.User) error {
  passwordHash, err := u.Crypt.Hasher(user.Password)
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

func (u *UserUseCases) Login(email, password string) (string, time.Time, error) {
  user, err := u.UserRepository.GetByEmail(email)
  if err != nil {
    return "", time.Now(), err
  }

  if !u.Crypt.HashComparer(password, user.Password) {
    return "", time.Now(), fmt.Errorf("invalid password")
  }

  expirationTime := time.Now().Add(5 * time.Minute) // 5 minutes

  token, expirationTime, err := u.Crypt.Encrypt(user.ID, expirationTime)
  if err != nil {
    return "", time.Now(), err
  }

  return token, expirationTime, nil
}
