package domain

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRepository interface {
  GetAll() ([]*User, error)
}

type UserUseCases interface {
  GetAll() ([]*User, error)
}
