package goauthultimate

import (
	"crypto/rsa"
	"encoding/base64"
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

var secret = []byte(base64.StdEncoding.EncodeToString(RandomBytes(32)))

func CreateToken(Username string, t time.Duration) (string, error) {

	// Verify username and password

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	// Set claims
	claims["username"] = Username
	claims["exp"] = time.Now().Add(t).Unix()

	// Create token
	tokenString, err := token.SignedString(secret)

	return tokenString, err
}
func ValidateToken(tokenString string) (bool, error) {

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return false, err
	}

	// Validate claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		_ = claims["username"].(string)
		log.Println(token.SignedString(secret))
		return true, nil

	} else {
		return false, errors.New("invalid token")
	}

}
func UsernameFromToken(T *string) (string, error) {
	token, err := jwt.Parse(*T, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if username := claims["username"].(string); username != "" {

			return username, nil
		} else {
			return "", err
		}
	} else {
		return "", err
	}
}
func CreateRefreshToken(username string) (string, error) {
	// Lookup refresh token in DB and make sure valid

	// Create new JWT token
	refreshToken, err := CreateToken(username, time.Duration(168*time.Hour))
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}

var PublicKey *rsa.PublicKey
var BaseToken string

func ValidateTokenSso(tokenString string) (bool, string, error) {

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return PublicKey, nil
	})
	if err != nil {
		return false, "", err
	}

	// Validate claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		usernamee := claims["username"].(string)
		// log.Println(token.SignedString(secret))
		return true, usernamee, nil

	} else {
		return false, "", errors.New("invalid token")
	}
}
