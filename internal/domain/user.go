package domain

import "time"

type User struct {
	ID       string `json:"id" bson:"_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

type UserRepository interface {
	GetAll() ([]*User, error)
	GetByID(id string) (*User, error)
	GetByEmail(email string) (*User, error)
	Create(user *User) error
}

type Cryptography interface {
  Hasher(plaintext string) (string, error)
  HashComparer(plaintext, hash string) bool
  Encrypt(plaintext string, expirationTime time.Time) (string, time.Time, error)
}

type UserUseCases interface {
	GetAll() ([]*User, error)
	GetByID(id string) (*User, error)
	Create(user *User) error
  Login(email, password string) (string, time.Time, error)
}
