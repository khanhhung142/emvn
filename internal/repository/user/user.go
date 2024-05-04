package user_repository

import (
	"context"
	"emvn/consts"
	"emvn/database/nosql"
	"emvn/internal/model"
	"fmt"
	"log/slog"

	"go.mongodb.org/mongo-driver/bson"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user model.User) (model.User, error)
	GetUserByUsername(ctx context.Context, username string) (model.User, error)
}

type userRepository struct {
	noSqlDB nosql.NoSQLInterface
}

// Singleton pattern
var localUserRepository IUserRepository

func InitUserRepository(noSqlDB nosql.NoSQLInterface) {
	localUserRepository = &userRepository{
		noSqlDB: noSqlDB,
	}
}

func UserRepository() IUserRepository {
	return localUserRepository
}

func (r *userRepository) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	result, err := r.noSqlDB.CreateIfNotExists(ctx, consts.MongoDBCollectionUsers, bson.M{"username": user.Username}, user)
	if err != nil {
		slog.Error("CreateUser", "error", err)
		return model.User{}, err
	}
	err = result.Decode(&user)
	if err != nil {
		slog.Error("CreateUser", "error", err)
		return model.User{}, err
	}
	return user, nil
}

func (r *userRepository) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	result, err := r.noSqlDB.FindOne(ctx, consts.MongoDBCollectionUsers, bson.M{"username": username})
	if err != nil {
		fmt.Println("GetUserByUsername", "error", err)
		return model.User{}, err
	}

	var user model.User
	err = result.Decode(&user)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
