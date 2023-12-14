package repositories

import (
	"context"

	"github.com/joaovds/auth-jwt/internal/domain"
	database "github.com/joaovds/auth-jwt/internal/infra/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) GetAll() ([]*domain.User, error) {
	usersCollection := database.Instance.DB.Collection("users")

	cursor, err := usersCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

  users := make([]*domain.User, 0)
	err = cursor.All(context.TODO(), &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
