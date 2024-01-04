package cryptography

import "time"

type Cryptography struct {
  jwtAdapter JWTAdapter
  bcryptAdapter BcryptAdapter
}

func NewCryptography() *Cryptography {
  return &Cryptography{
    jwtAdapter: *NewJWTAdapter("secret-abcde-key"),
    bcryptAdapter: *NewBcrypt(14),
  }
}

func (c *Cryptography) Hasher(plaintext string) (string, error) {
  return c.bcryptAdapter.Hasher(plaintext)
}

func (c *Cryptography) HashComparer(plaintext, hash string) bool {
  return c.bcryptAdapter.HashComparer(plaintext, hash)
}

func (c *Cryptography) Encrypt(plaintext string, expirationTime time.Time) (string, time.Time, error) {
  return c.jwtAdapter.Encrypt(plaintext, expirationTime)
}

func (c *Cryptography) Decrypt(ciphertext string) (error) {
  return c.jwtAdapter.Decrypt(ciphertext)
}
