package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var accessType = "ACCESS"
var refreshType = "REFRESH"

func CreateAccessToken(userEmail string) (string, error) {
	accessSecret := os.Getenv("ACCESS_SECRET")
	return createToken(userEmail, accessSecret, accessType)
}

func CreateRefreshToken(userEmail string) (string, error) {
	refreshSecret := os.Getenv("REFRESH_SECRET")
	return createToken(userEmail, refreshSecret, accessType)
}

func createToken(userEmail string, secret string, tokenType string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_email"] = userEmail
	atClaims["type"] = tokenType
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}
