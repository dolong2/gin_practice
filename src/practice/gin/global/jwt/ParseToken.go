package jwt

import (
	"gin_practice/src/practice/gin/domain/user/entity"
	"gin_practice/src/practice/gin/domain/user/repository"
	"gin_practice/src/practice/gin/global/jwt/exception"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strings"
)

func GetCurrentUser(r *http.Request) (*entity.User, error) {
	email, err := ExtractTokenMetadata(r, "user_email")
	if err != nil {
		return nil, err
	}
	user, _ := repository.GetByEmail(*email)
	return &user, nil
}

func ParseToken(r *http.Request) (string, error) {
	bearToken := r.Header.Get("Authorization")
	if !strings.Contains(bearToken, "Bearer ") {
		return "", exception.InvalidTokenException()
	}
	return strings.Replace(bearToken, "Bearer ", "", 1), nil
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString, parseException := ParseToken(r)
	if parseException != nil {
		return nil, parseException
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, exception.InvalidTokenException()
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractTokenMetadata(r *http.Request, key string) (*string, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		data, ok := claims[key].(string)
		if !ok {
			return nil, exception.InvalidTokenException()
		}
		return &data, nil
	}
	return nil, err
}
