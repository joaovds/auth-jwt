package domain

type User struct {
	ID       string `json:"id" bson:"_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRepository interface {
	GetAll() ([]*User, error)
	GetByID(id string) (*User, error)
	Create(user *User) error
}

type UserUseCases interface {
	GetAll() ([]*User, error)
	GetByID(id string) (*User, error)
	Create(user *User) error
}
