package service

import (
	"gin_practice/src/practice/gin/domain/user/controller/dto/request"
	"gin_practice/src/practice/gin/domain/user/exception"
	"gin_practice/src/practice/gin/domain/user/repository"
)

func CreateUser(request request.SignUpRequest) error {
	if repository.ExistsByEmail(request.Email) {
		return exception.UserAlreadyExistsException()
	}
	repository.Save(request.ToUser())
	return nil
}
