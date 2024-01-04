package cryptography

import (
	"golang.org/x/crypto/bcrypt"
)

type BcryptAdapter struct {
  salt int
}

func NewBcrypt(salt int) *BcryptAdapter {
  return &BcryptAdapter{
    salt: salt,
  }
}

func (b *BcryptAdapter) Hasher(plaintext string) (string, error) {
  bytes, err := bcrypt.GenerateFromPassword([]byte(plaintext), b.salt)
  return string(bytes), err
}

func (b *BcryptAdapter) HashComparer(plaintext, hash string) bool {
  err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plaintext))
  return err == nil
}

