package cryptography

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTAdapter struct {
  secret string
}

type Claims struct {
  ID string `json:"id"`
  jwt.RegisteredClaims
}

func NewJWTAdapter(secret string) *JWTAdapter {
  return &JWTAdapter{
    secret: secret,
  }
}

func (j *JWTAdapter) Encrypt(plaintext string, expirationTime time.Time) (string, time.Time, error) {
  claims := &Claims{
    ID: plaintext,
    RegisteredClaims: jwt.RegisteredClaims{
      ExpiresAt: jwt.NewNumericDate(expirationTime),
    },
  }

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  tokenString, err := token.SignedString([]byte(j.secret))
  if err != nil {
    return "", time.Now(), err
  }

  return tokenString, expirationTime, nil
}

func (j *JWTAdapter) Decrypt(ciphertext string) (error) {
  token, err := jwt.ParseWithClaims(ciphertext, &Claims{}, func(token *jwt.Token) (interface{}, error) {
    return []byte(j.secret), nil
  })
  if err != nil {
    return err
  }

  _, ok := token.Claims.(*Claims)
  if !ok {
    return err
  }

  return nil
}
