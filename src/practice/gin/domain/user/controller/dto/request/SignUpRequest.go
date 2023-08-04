package request

import "gin_practice/src/practice/gin/domain/user/entity"

type converter interface {
	toUser() entity.User
}

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (request SignUpRequest) ToUser() entity.User {
	return entity.User{
		Email:    request.Email,
		Password: request.Password,
		Name:     request.Name,
	}
}
