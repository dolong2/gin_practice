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
	return createToken(userEmail, accessSecret, accessType, 15)
}

func CreateRefreshToken(userEmail string) (string, error) {
	refreshSecret := os.Getenv("REFRESH_SECRET")
	return createToken(userEmail, refreshSecret, accessType, 60*24)
}

func createToken(userEmail string, secret string, tokenType string, ttl int) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_email"] = userEmail
	atClaims["type"] = tokenType
	atClaims["exp"] = time.Now().Add(time.Duration(ttl) * time.Minute).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}
