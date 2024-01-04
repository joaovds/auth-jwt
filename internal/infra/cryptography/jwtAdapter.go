package cryptography

type JWTAdapter struct {
  secret string
}

func NewJWTAdapter(secret string) *JWTAdapter {
  return &JWTAdapter{
    secret: secret,
  }
}

func (j *JWTAdapter) Encrypt(plaintext string) (string, error) {
  return "", nil
}
