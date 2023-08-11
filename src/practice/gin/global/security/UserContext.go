package security

import (
	"gin_practice/src/practice/gin/domain/user/entity"
	"gin_practice/src/practice/gin/global/security/exception"
)

var currentUser *entity.User

func GetUser() (*entity.User, error) {
	if currentUser == nil {
		return nil, exception.ContextEmptyException()
	}
	return currentUser, nil
}

func SetUser(user entity.User) {
	currentUser = &user
}
