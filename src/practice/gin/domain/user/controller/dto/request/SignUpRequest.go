package request

import "gin_practice/src/practice/gin/domain/user/entity"

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (request SignUpRequest) ToUser(encodedPassword string) entity.User {
	return entity.User{
		Email:    request.Email,
		Password: encodedPassword,
		Name:     request.Name,
	}
}
