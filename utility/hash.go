package utility

import (
	"emvn/config"
	"emvn/consts"
	"errors"
	"fmt"
	"log/slog"
	"math/rand"

	"github.com/golang-jwt/jwt/v5"
)

var ConfigJWTHashFunction = jwt.SigningMethodHS256

// Gen JWT
// Args: payload jwt map claims
func GenJWT(payload jwt.MapClaims) (token string, err error) {
	tokenJWT := jwt.NewWithClaims(ConfigJWTHashFunction, payload)
	token, err = tokenJWT.SignedString([]byte(config.GetConfig().Auth.SecretKey))
	return
}

// Validate JWT token
func ParseJWT(tokenString string) (jwt.MapClaims, error) {
	tokenParse, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.GetConfig().Auth.SecretKey), nil
	})

	var claims jwt.MapClaims

	if tokenParse != nil {
		var ok bool
		claims, ok = tokenParse.Claims.(jwt.MapClaims)
		if !ok {
			slog.Error(err.Error())
			return nil, consts.CodeInvalidToken
		}
	}
	switch {
	case err == nil && tokenParse.Valid:
		return claims, nil
	case err != nil && (errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet)):
		return claims, jwt.ErrTokenExpired
	default:
		return nil, consts.CodeInvalidToken
	}
}

func GenerateRandomPassword() string {
	// Define character sets
	uppercaseLetters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercaseLetters := "abcdefghijklmnopqrstuvwxyz"
	digits := "0123456789"
	specialCharacters := "@"

	// Combine character sets
	allCharacters := uppercaseLetters + lowercaseLetters + digits + specialCharacters

	// Generate a password
	var passwordBuilder []rune

	// Add special char
	specialCharIndex := rand.Intn(len(specialCharacters))
	passwordBuilder = append(passwordBuilder, []rune(specialCharacters)[specialCharIndex])

	for i := 0; i < 8; i++ {
		// Randomly select a character from allCharacters
		charIndex := rand.Intn(len(allCharacters))
		passwordBuilder = append(passwordBuilder, []rune(allCharacters)[charIndex])
	}

	// Add number char
	numberCharIndex := rand.Intn(len(digits))
	passwordBuilder = append(passwordBuilder, []rune(digits)[numberCharIndex])

	return string(passwordBuilder)
}
