package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/pimentafm/auction-go/configuration/logger"
	"github.com/pimentafm/auction-go/internal/entity/user_entity"
	"github.com/pimentafm/auction-go/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserEntityMongo struct {
	Id   string `bson:"_id"`
	Name string `bson:"name"`
}

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(database *mongo.Database) *UserRepository {
	return &UserRepository{
		Collection: database.Collection("users"),
	}
}

func (ur *UserRepository) FindUserById(ctx context.Context, id string) (*user_entity.User, *internal_error.InternalError) {
	filter := bson.M{"_id": id}

	var userEntityMongo UserEntityMongo

	err := ur.Collection.FindOne(ctx, filter).Decode(&userEntityMongo)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error(fmt.Sprintf("User not found with this Id = %d", id), err)
			return nil, internal_error.NewNotFoundError(fmt.Sprintf("User not found with this Id = %d", id))
		}

		logger.Error(fmt.Sprintf("Error trying to find user by Id"), err)
		return nil, internal_error.NewInternalServerError("Error trying to find user by Id")
	}

	userEntity := &user_entity.User{
		Id:   userEntityMongo.Id,
		Name: userEntityMongo.Name,
	}

	return userEntity, nil
}
