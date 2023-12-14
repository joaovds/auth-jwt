package repositories

import (
	"context"
	"errors"

	"github.com/joaovds/auth-jwt/internal/domain"
	database "github.com/joaovds/auth-jwt/internal/infra/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func (ur *UserRepository) GetByID(id string) (*domain.User, error) {
  usersCollection := database.Instance.DB.Collection("users")

  idHex, err := primitive.ObjectIDFromHex(id)
  if err != nil {
    return nil, err
  }

  user := &domain.User{}
  err = usersCollection.FindOne(context.TODO(), bson.M{
    "_id": idHex,
  }).Decode(user)
  if err != nil {
    if err == mongo.ErrNoDocuments {
      return nil, errors.New("user not found")
    }
    return nil, err
  }

  return user, nil
}
