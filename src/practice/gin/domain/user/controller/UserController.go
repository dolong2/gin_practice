package controller

import (
	"gin_practice/src/practice/gin/domain/user/controller/dto/request"
	"gin_practice/src/practice/gin/domain/user/service"
	globalController "gin_practice/src/practice/gin/global/util/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUpUserController(r *gin.Engine) *gin.Engine {
	signUp := func(request request.SignUpRequest) error {
		return service.CreateUser(request)
	}
	globalController.POST(r, "/users", request.SignUpRequest{}, http.StatusCreated, signUp)

	return r
}
