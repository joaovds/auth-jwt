package cryptography

import (
	"golang.org/x/crypto/bcrypt"
)

type Bcrypt struct{
  salt int
}

func NewBcrypt(salt int) *Bcrypt {
  return &Bcrypt{
    salt: salt,
  }
}

func (b *Bcrypt) Hasher(plaintext string) (string, error) {
  bytes, err := bcrypt.GenerateFromPassword([]byte(plaintext), b.salt)
  return string(bytes), err
}

func (b *Bcrypt) HashComparer(plaintext, hash string) bool {
  err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plaintext))
  return err == nil
}
