package repository

import (
	"context"
	"errors"

	"auth-back/internal/models"
	"auth-back/pkg/db/mongodb"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

type UsersRepo struct {
	db *mongo.Collection
}

func NewUsersRepo(db *mongo.Database) *UsersRepo {
	return &UsersRepo{
		db: db.Collection(usersCollection),
	}
}

func (r *UsersRepo) Create(ctx context.Context, user models.User) error {
	_, err := r.db.InsertOne(ctx, user)
	if mongodb.IsDuplicate(err) {
		return models.ErrUserAlreadyExists
	}

	return err
}

func (r *UsersRepo) GetByCredentials(ctx context.Context, email, password string) (models.User, error) {
	var user models.User
	if err := r.db.FindOne(ctx, bson.M{"email": email, "password": password}).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.User{}, models.ErrUserNotFound
		}

		return models.User{}, err
	}

	return user, nil
}
