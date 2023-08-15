package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"strconv"
	"time"
)

func CreateAccessToken(userEmail string) (string, error) {
	accessSecret := os.Getenv("ACCESS_SECRET")
	accessExp, _ := strconv.Atoi(os.Getenv("ACCESS_EXP"))
	return createToken(userEmail, accessSecret, AccessType, accessExp)
}

func CreateRefreshToken(userEmail string) (string, error) {
	refreshSecret := os.Getenv("REFRESH_SECRET")
	refreshExp, _ := strconv.Atoi(os.Getenv("REFRESH_EXP"))
	return createToken(userEmail, refreshSecret, RefreshType, refreshExp)
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
