package service

import (
	"gin_practice/src/practice/gin/domain/user/controller/dto/request"
	"gin_practice/src/practice/gin/domain/user/exception"
	"gin_practice/src/practice/gin/domain/user/repository"
	"gin_practice/src/practice/gin/global/passwordEncoder"
)

func CreateUser(request request.SignUpRequest) error {
	if repository.ExistsByEmail(request.Email) {
		return exception.UserAlreadyExistsException()
	}
	password, err := passwordEncoder.EncodePassword(request.Password)
	if err != nil {
		return err
	}
	repository.Save(request.ToUser(password))
	return nil
}
