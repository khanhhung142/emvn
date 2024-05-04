package auth_usecase

import (
	"context"
	"emvn/config"
	"emvn/utility"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Token interface {
	Gen() (token string, err error)
	Verify(token string) (err error)
}

type AccessToken struct {
	Iss string `json:"iss"`
	Sub string `json:"sub"`
	Exp int64  `json:"exp"`
}

func (ac *AccessToken) Gen() (token string, err error) {
	expTime := time.Now().Add(time.Duration(config.GetConfig().Auth.AccessTokenExpireTime) * time.Minute)
	ac.Exp = expTime.Unix()

	mapClaims := jwt.MapClaims{}
	dataByte, _ := json.Marshal(ac)
	json.Unmarshal(dataByte, &mapClaims)

	token, err = utility.GenJWT(mapClaims)
	return
}

func (ac *AccessToken) Verify(ctx context.Context, token string) (err error) {
	jwtMap, err := utility.ParseJWT(token)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	dataByte, _ := json.Marshal(jwtMap)
	err = json.Unmarshal(dataByte, &ac)
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
