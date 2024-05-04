package auth_usecase

import (
	"context"
	"emvn/consts"
	"emvn/internal/model"
	user_repository "emvn/internal/repository/user"
	"fmt"
	"log/slog"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IAuthUsecase interface {
	SignUp(ctx context.Context, user model.User) error
	SignIn(ctx context.Context, username, password string) (SignInOutput, error)
}

type authUsecase struct {
	userRepo user_repository.IUserRepository
}

var localAuthUsecase IAuthUsecase

func InitAuthUsecase(userRepo user_repository.IUserRepository) {
	localAuthUsecase = &authUsecase{
		userRepo: userRepo,
	}
}

func AuthUsecase() IAuthUsecase {
	return localAuthUsecase
}

func (u *authUsecase) SignUp(ctx context.Context, user model.User) error {
	// check if user already exists
	dbUser, err := u.userRepo.GetUserByUsername(ctx, user.Username)
	if err != nil {
		if err.Error() != mongo.ErrNoDocuments.Error() {
			return consts.CodeInternalError
		}
	}

	if dbUser.Username == user.Username {
		return consts.CodeUserAlreadyExists
	}
	// Add hash password
	user.Password, err = hashPassword(user.Password)
	if err != nil {
		slog.Error(err.Error())
		return consts.CodeInternalError
	}

	user.ID = primitive.NewObjectID().Hex()
	_, err = u.userRepo.CreateUser(ctx, user)
	if err != nil {
		return consts.CodeInternalError
	}

	return nil
}

func (u *authUsecase) SignIn(ctx context.Context, username, password string) (SignInOutput, error) {
	var output SignInOutput
	dbUser, err := u.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		if err.Error() == mongo.ErrNoDocuments.Error() {
			return output, consts.CodeUserNotFound
		}
		return output, err
	}
	if !checkPasswordHash(password, dbUser.Password) {
		return output, consts.CodeWrongPassword
	}

	// Gen Access Token
	acToken := AccessToken{
		Sub: fmt.Sprintf("%v", dbUser.ID),
		Iss: fmt.Sprintf("%v", dbUser.ID),
	}
	acTokenString, err := acToken.Gen()
	if err != nil {
		return output, err
	}

	output.Token = acTokenString
	output.Exp = acToken.Exp
	return output, nil
}
